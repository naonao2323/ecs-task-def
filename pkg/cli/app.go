package cli

import (
	"context"
	"ecs-task-def-action/pkg/decoder"
	"ecs-task-def-action/pkg/encoder"
	"ecs-task-def-action/pkg/git"
	"ecs-task-def-action/pkg/github"
	"ecs-task-def-action/pkg/logger"
	"ecs-task-def-action/pkg/transformer"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	ecsDecoder "ecs-task-def-action/pkg/decoder/ecs"

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
	a := app{
		logger: logger,
	}
	cmd := cobra.Command{
		Use:   "ecs-task-def-action",
		Short: "start ecs-task-def-action",
		RunE:  a.run,
	}
	cmd.Flags().StringVar(&a.tag, "target-tag", a.tag, "target tag")
	cmd.Flags().StringVar(&a.containerName, "container-name", a.containerName, "container name")
	cmd.Flags().StringVarP(&a.taskPath, "task-path", a.taskPath, "", "the path to the task definition")
	cmd.Flags().StringVarP(&a.containerPath, "container-path", a.containerPath, "", "the path to the container definition")
	cmd.Flags().StringVar(&a.githubOwner, "github-owner", a.githubOwner, "github owner")
	cmd.Flags().StringVar(&a.githubToken, "github-token", a.githubToken, "github token")
	cmd.Flags().StringVar(&a.githubRepository, "github-repository", a.githubRepository, "github repositoy")
	cmd.Flags().StringVar(&a.gitEmail, "github-email", a.gitEmail, "git email")
	cmd.Flags().StringVar(&a.gitUsername, "github-username", a.gitUsername, "git username")
	cmd.Flags().StringVar(&a.githubUrl, "github-url", a.githubUrl, "github url")
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
	defer func() {
		err := a.logger.Sync()
		if err != nil && errors.Is(err, syscall.EINVAL) {
			// Sync is not supported on os.Stderr / os.Stdout on arm64 alpine:3.20.3
		} else if err != nil {
			log.Fatal(err)
		}
	}()
	ctx := context.Background()
	target := func(tag string, path string) string {
		return fmt.Sprintf("/%s/%s", tag, path)
	}
	outputer := func(in []byte, tag, path string) error {
		if err := os.WriteFile(target(tag, path), in, 0o644); err != nil {
			a.logger.Error("fail to write file", zap.Error(errors.ErrUnsupported))
			return err
		}
		return nil
	}
	gitClient := git.NewGitClient(a.logger, a.gitUsername, a.gitEmail, a.tag, a.githubToken)
	githubClient := github.NewGithubClient(ctx, a.logger, a.githubToken, a.githubOwner, a.githubRepository)
	if err := gitClient.Clone(a.githubUrl); err != nil {
		return err
	}

	var s strategy
	if a.containerPath != "" && a.taskPath != "" {
		s = CONTAINER_DEFINITION
	} else if a.containerPath == "" {
		s = TASK_DEFINITION
	} else if a.taskPath == "" {
		s = CONTAINER_DEFINITION
	} else {
		return errors.New("empty definition file")
	}

	switch s {
	case TASK_DEFINITION:
		ext := filepath.Ext(target(a.tag, a.taskPath))
		format := encoder.GetFormat(ext)
		if format == encoder.Unknow {
			err := errors.New("unknow extension")
			a.logger.Error("unknown extension", zap.Error(err))
			return err
		}
		in, err := os.ReadFile(target(a.tag, a.taskPath))
		if err != nil {
			a.logger.Error("fail to open target file", zap.Error(err))
			return err
		}
		encoder := ecsEncoder.NewEcsTask(a.logger)
		transformer := ecsTransformer.NewTaskTransformer()
		decoder := ecsDecoder.NewEcsTaskDecoder(a.logger)
		err = executeTaskDefinition(ctx, in, a.containerName, a.tag, a.taskPath, a.githubUrl, format, encoder, transformer, decoder, outputer, gitClient, githubClient)
		if err != nil {
			return err
		}

	case CONTAINER_DEFINITION:
		ext := filepath.Ext(target(a.tag, a.taskPath))
		format := encoder.GetFormat(ext)
		if format == encoder.Unknow {
			err := errors.New("unknow extension")
			a.logger.Error("", zap.Error(err))
			return err
		}
		in, err := os.ReadFile(target(a.tag, a.taskPath))
		if err != nil {
			a.logger.Error("fail to open target file", zap.Error(err))
			return err
		}
		encoder := ecsEncoder.NewEcsContainer(a.logger)
		transformer := ecsTransformer.NewEcsContainerTransformer()
		decoder := ecsDecoder.NewEcsContainerDecoder(a.logger)
		err = executeContainerDefinition(ctx, in, a.containerName, a.tag, a.containerPath, a.githubUrl, format, encoder, transformer, decoder, outputer, gitClient, githubClient)
		if err != nil {
			return err
		}
	}

	return nil
}

func executeContainerDefinition(
	ctx context.Context,
	in []byte,
	app string,
	tag string,
	path string,
	githubUrl string,
	format encoder.Format,
	encoder encoder.EcsContainerEncoder,
	transformer transformer.EcsContainerTransformer,
	decoder decoder.EcsContainerDecoder,
	outputer func(in []byte, tag, path string) error,
	gitClient git.Git,
	githubClient github.Github,
) error {
	def, err := encoder.Encode(in, format)
	if def == nil {
		return errors.New("empty definition")
	}
	if err != nil {
		return err
	}
	transformed := transformer.Transform(tag, app, *def)
	decoded, err := decoder.Decode(transformed, convertFormat(format))
	if err != nil {
		return err
	}
	if err := outputer(decoded, tag, path); err != nil {
		return err
	}
	if err := gitClient.Add(path); err != nil {
		return err
	}
	if err := gitClient.Commit(tag); err != nil {
		return err
	}
	if err := gitClient.CheckOut(tag); err != nil {
		return err
	}
	if err := gitClient.Push(tag); err != nil {
		return err
	}
	if err := githubClient.CreatePullRequest(ctx, tag, tag); err != nil {
		return err
	}
	return nil
}

func executeTaskDefinition(
	ctx context.Context,
	in []byte,
	app string,
	tag string,
	path string,
	githubUrl string,
	format encoder.Format,
	encoder encoder.EcsTaskEncoder,
	transformer transformer.EcsTaskTransformer,
	decoder decoder.EcsTaskDecoder,
	outputer func(in []byte, tag, path string) error,
	gitClient git.Git,
	githubClient github.Github,
) error {
	def, err := encoder.Encode(in, format)
	if def == nil {
		return errors.New("empty definition")
	}
	if err != nil {
		return err
	}
	transformed := transformer.Transform(tag, app, *def)
	decoded, err := decoder.Decode(transformed, convertFormat(format))
	if err != nil {
		return err
	}
	if err := outputer(decoded, tag, path); err != nil {
		return err
	}
	if err := gitClient.Add(path); err != nil {
		return err
	}
	if err := gitClient.Commit(tag); err != nil {
		return err
	}
	if err := gitClient.CheckOut(tag); err != nil {
		return err
	}
	if err := gitClient.Push(tag); err != nil {
		return err
	}
	if err := githubClient.CreatePullRequest(ctx, tag, tag); err != nil {
		return err
	}
	return nil
}

func convertFormat(format encoder.Format) decoder.Format {
	return decoder.Format(format)
}