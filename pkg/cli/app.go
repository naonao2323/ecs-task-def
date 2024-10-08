package cli

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"github.com/naonao2323/ecs-task-def/pkg/decoder"
	"github.com/naonao2323/ecs-task-def/pkg/encoder"
	"github.com/naonao2323/ecs-task-def/pkg/git"
	"github.com/naonao2323/ecs-task-def/pkg/github"
	"github.com/naonao2323/ecs-task-def/pkg/logger"
	"github.com/naonao2323/ecs-task-def/pkg/plovider/ecs"
	"github.com/naonao2323/ecs-task-def/pkg/transformer"

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
		Use:   "ecs-task-def",
		Short: "start ecs-task-def",
		RunE:  a.run,
	}
	cmd.Flags().StringVar(&a.tag, "target-tag", a.tag, "target tag")
	cmd.Flags().StringVar(&a.containerName, "container-name", a.containerName, "container name")
	cmd.Flags().StringVarP(&a.taskPath, "task-path", a.taskPath, "", "the path to the task definition")
	cmd.Flags().StringVarP(&a.containerPath, "container-path", a.containerPath, "", "the path to the container definition")
	cmd.Flags().StringVar(&a.githubOwner, "github-owner", a.githubOwner, "github owner")
	cmd.Flags().StringVar(&a.githubToken, "github-token", a.githubToken, "github token")
	cmd.Flags().StringVar(&a.githubRepository, "github-repository", a.githubRepository, "github repositoy (etc ecs-task-def)")
	cmd.Flags().StringVar(&a.gitEmail, "github-email", a.gitEmail, "git email")
	cmd.Flags().StringVar(&a.gitUsername, "github-username", a.gitUsername, "git username")
	cmd.Flags().StringVar(&a.githubUrl, "github-url", a.githubUrl, "github repository https url (etc https://github.com/naonao2323/ecs-task-def.git)")
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
	UNKNOW_DEFINITION
)

func (a *app) run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	target := func(destination string, path string) string {
		return fmt.Sprintf("%s/%s", destination, path)
	}
	outputer := func(in []byte, destination, path string) error {
		if err := os.WriteFile(target(destination, path), in, 0o644); err != nil {
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
	defer func() {
		err := gitClient.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = a.logger.Sync()
		if err != nil && errors.Is(err, syscall.EINVAL) {
			// Sync is not supported on os.Stderr / os.Stdout on arm64 alpine:3.20.3
		} else if err != nil {
			log.Fatal(err)
		}
	}()

	s := selectStrategy(a.containerPath, a.taskPath)

	switch s {
	case TASK_DEFINITION:
		ext := filepath.Ext(target(gitClient.GetDestination(), a.taskPath))
		format := encoder.GetFormat(ext)
		if format == encoder.Unknow {
			err := errors.New("unknow extension")
			a.logger.Error("unknown extension", zap.Error(err))
			return err
		}
		in, err := os.ReadFile(target(gitClient.GetDestination(), a.taskPath))
		if err != nil {
			a.logger.Error("fail to open target file", zap.Error(err))
			return err
		}
		transformer := transformer.NewTransformer[ecs.TaskDefinition]()
		encoder := encoder.NewEncoder[ecs.TaskDefinition](a.logger)
		decoder := decoder.NewDecoder[ecs.TaskDefinition](a.logger)
		err = execute(
			ctx,
			a.logger,
			in,
			a.containerName,
			a.tag,
			a.taskPath,
			a.githubUrl,
			format,
			transformer,
			encoder,
			decoder,
			outputer,
			gitClient,
			githubClient,
		)
		if err != nil {
			a.logger.Error("execute task definition", zap.Error(err))
			return err
		}

	case CONTAINER_DEFINITION:
		ext := filepath.Ext(target(gitClient.GetDestination(), a.containerPath))
		format := encoder.GetFormat(ext)
		if format == encoder.Unknow {
			err := errors.New("unknow extension")
			a.logger.Error("", zap.Error(err))
			return err
		}
		in, err := os.ReadFile(target(gitClient.GetDestination(), a.containerPath))
		if err != nil {
			a.logger.Error("fail to open target file", zap.Error(err))
			return err
		}
		transformer := transformer.NewTransformer[[]ecs.ContainerDefinition]()
		encoder := encoder.NewEncoder[[]ecs.ContainerDefinition](a.logger)
		decoder := decoder.NewDecoder[[]ecs.ContainerDefinition](a.logger)
		err = execute(
			ctx,
			a.logger,
			in,
			a.containerName,
			a.tag,
			a.containerPath,
			a.githubUrl,
			format,
			transformer,
			encoder,
			decoder,
			outputer,
			gitClient,
			githubClient,
		)
		if err != nil {
			a.logger.Info("execute container definition", zap.Error(err))
			return err
		}
	case UNKNOW_DEFINITION:
		return errors.New("unknow definition")
	}

	return nil
}

func selectStrategy(containerPath, taskPath string) strategy {
	if containerPath != "" && taskPath != "" {
		return CONTAINER_DEFINITION
	} else if containerPath == "" && taskPath != "" {
		return TASK_DEFINITION
	} else if containerPath != "" && taskPath == "" {
		return CONTAINER_DEFINITION
	} else {
		return UNKNOW_DEFINITION
	}
}

func convertFormat(format encoder.Format) decoder.Format {
	return decoder.Format(format)
}

func execute[P ecs.EcsTarget](
	ctx context.Context,
	logger *zap.Logger,
	in []byte,
	app string,
	tag string,
	path string,
	githubUrl string,
	format encoder.Format,
	transformer transformer.Transformer[P],
	encoder encoder.Encoder[P],
	decoder decoder.Decoder[P],
	outputer func(in []byte, destination, path string) error,
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
	if err := outputer(decoded, gitClient.GetDestination(), path); err != nil {
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
