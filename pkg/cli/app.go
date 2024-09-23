package cli

import (
	"context"
	"ecs-task-def-action/pkg/decoder"
	"ecs-task-def-action/pkg/git"
	"ecs-task-def-action/pkg/github"
	"ecs-task-def-action/pkg/logger"
	"ecs-task-def-action/pkg/transformer"
	"fmt"
	"log"
	"os"
	"path/filepath"

	ecsDecoder "ecs-task-def-action/pkg/decoder/ecs"

	"ecs-task-def-action/pkg/encoder"
	ecsEncoder "ecs-task-def-action/pkg/encoder/ecs"
	ecsTransformer "ecs-task-def-action/pkg/transformer/ecs"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type app struct {
	logger *zap.Logger

	containerName    string
	taskPath         string
	containerPath    string
	githubToken      string
	githubOwner      string
	githubRepository string
	githubUrl        string
	gitEmail         string
	gitUsername      string
	tag              string
}

func NewCommand() cobra.Command {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
	app := app{
		logger: logger,
	}
	cmd := cobra.Command{
		Use:   "ecs-task-def-action",
		Short: "start ecs-task-def-action",
		RunE:  app.run,
	}
	cmd.Flags().StringVar(&app.tag, "target-tag", app.tag, "target tag")
	cmd.Flags().StringVar(&app.containerName, "container-name", app.containerName, "container name")
	cmd.Flags().StringVarP(&app.taskPath, "task-path", app.taskPath, "", "the path to the task definition")
	cmd.Flags().StringVarP(&app.containerPath, "container-path", app.containerPath, "", "the path to the container definition")
	cmd.Flags().StringVar(&app.githubOwner, "github-owner", app.githubOwner, "github owner")
	cmd.Flags().StringVar(&app.githubToken, "github-token", app.githubToken, "github token")
	cmd.Flags().StringVar(&app.githubRepository, "github-repository", app.githubRepository, "github repositoy")
	cmd.Flags().StringVar(&app.gitEmail, "github-email", app.gitEmail, "git email")
	cmd.Flags().StringVar(&app.gitUsername, "github-username", app.gitUsername, "git username")
	cmd.Flags().StringVar(&app.githubUrl, "github-url", app.githubUrl, "github url")
	return cmd
}

func Execute(cmd cobra.Command) {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

type strategy int

const (
	TASK_DEFINITION strategy = iota
	CONTAINER_DEFINITION
)

func (a *app) run(cmd *cobra.Command, args []string) error {
	var strategy strategy
	if a.containerPath != "" && a.taskPath != "" {
		strategy = CONTAINER_DEFINITION
	} else if a.containerPath == "" {
		strategy = TASK_DEFINITION
	} else if a.taskPath == "" {
		strategy = CONTAINER_DEFINITION
	} else {
		return fmt.Errorf("empty definition file")
	}

	outputer := func(in []byte, path string) error {
		return os.WriteFile(path, in, 0644)
	}

	gitClient := git.NewGitClient(a.gitUsername, a.gitEmail, a.githubToken)
	githubClient := github.NewGithubClient(a.githubToken, a.githubOwner, a.githubRepository)
	gitClient.Clone(a.githubUrl, a.tag)

	switch strategy {
	case TASK_DEFINITION:
		ext := filepath.Ext("/" + a.tag + "/" + a.taskPath)
		format := encoder.GetFormat(ext)
		if format == encoder.Unknow {
			return fmt.Errorf("unknow extension")
		}
		in, err := os.ReadFile("/" + a.tag + "/" + a.taskPath)
		if err != nil {
			return err
		}
		encoder := ecsEncoder.NewEcsTask()
		transformer := ecsTransformer.NewTaskTransformer()
		decoder := ecsDecoder.NewEcsTaskDecoder()
		err = executeTaskDefinition(in, a.containerName, a.tag, a.taskPath, a.githubUrl, format, encoder, transformer, decoder, outputer, gitClient, githubClient)
		if err != nil {
			return err
		}

	case CONTAINER_DEFINITION:
		ext := filepath.Ext("/" + a.tag + "/" + a.containerPath)
		format := encoder.GetFormat(ext)
		if format == encoder.Unknow {
			return fmt.Errorf("unknow extension")
		}
		in, err := os.ReadFile("/" + a.tag + "/" + a.containerPath)
		if err != nil {
			return err
		}
		encoder := ecsEncoder.NewEcsContainer()
		transformer := ecsTransformer.NewEcsContainerTransformer()
		decoder := ecsDecoder.NewEcsContainerDecoder()
		err = executeContainerDefinition(in, a.containerName, a.tag, a.containerPath, a.githubUrl, format, encoder, transformer, decoder, outputer, gitClient, githubClient)
		if err != nil {
			return err
		}
	}

	return nil
}

func executeContainerDefinition(
	in []byte,
	app string,
	tag string,
	path string,
	githubUrl string,
	format encoder.Format,
	encoder encoder.EcsContainerEncoder,
	transformer transformer.EcsContainerTransformer,
	decoder decoder.EcsContainerDecoder,
	outputer func(in []byte, path string) error,
	gitClient git.Git,
	githubClient github.Github,
) error {
	ctx := context.Background()
	def := encoder.Encode(in, format)
	if def == nil {
		return fmt.Errorf("empty definition")
	}
	transformed := transformer.Transform(tag, app, *def)
	decoded := decoder.Decode(transformed, convertFormat(format))
	if err := outputer(decoded, "/"+tag+"/"+path); err != nil {
		return err
	}
	if err := gitClient.Add(path, tag); err != nil {
		return err
	}
	if err := gitClient.Commit(tag, tag); err != nil {
		return err
	}
	if err := gitClient.CheckOut(tag, tag); err != nil {
		return err
	}
	if err := gitClient.Push(tag, tag); err != nil {
		return err
	}
	if err := githubClient.CreatePullRequest(ctx, tag); err != nil {
		return err
	}
	return nil
}

func executeTaskDefinition(
	in []byte,
	app string,
	tag string,
	path string,
	githubUrl string,
	format encoder.Format,
	encoder encoder.EcsTaskEncoder,
	transformer transformer.EcsTaskTransformer,
	decoder decoder.EcsTaskDecoder,
	outputer func(in []byte, path string) error,
	gitClient git.Git,
	githubClient github.Github,
) error {
	ctx := context.Background()
	def := encoder.Encode(in, format)
	if def == nil {
		return fmt.Errorf("empty definition")
	}
	transformed := transformer.Transform(tag, app, *def)
	decoded := decoder.Decode(transformed, convertFormat(format))
	if err := outputer(decoded, "/"+tag+"/"+path); err != nil {
		return err
	}
	if err := gitClient.Add(path, tag); err != nil {
		return err
	}
	if err := gitClient.Commit(tag, tag); err != nil {
		return err
	}
	if err := gitClient.CheckOut(tag, tag); err != nil {
		return err
	}
	if err := gitClient.Push(tag, tag); err != nil {
		return err
	}
	if err := githubClient.CreatePullRequest(ctx, tag); err != nil {
		return err
	}
	return nil
}

func convertFormat(format encoder.Format) decoder.Format {
	return decoder.Format(format)
}
