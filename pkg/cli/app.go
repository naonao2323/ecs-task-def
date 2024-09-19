package cli

import (
	"os"

	"github.com/spf13/cobra"
)

type app struct {
	containerName string
	taskPath      string
	containerPath string
	output        string
}

func NewCommand() cobra.Command {
	app := app{}
	cmd := cobra.Command{
		Use:   "ecs-task-def-action",
		Short: "start ecs-task-def-action",
		RunE:  app.run,
	}
	cmd.Flags().StringVar(&app.containerName, "container-name", app.containerName, "container name")
	cmd.Flags().StringVarP(&app.taskPath, "task-path", app.taskPath, "", "the path to the task definition")
	cmd.Flags().StringVarP(&app.containerPath, "container-path", app.containerPath, "", "the path to the container definition")
	cmd.Flags().StringVar(&app.output, "output", app.output, "container name")
	return cmd
}

func Execute(cmd cobra.Command) {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func (a *app) run(cmd *cobra.Command, args []string) error {
	println("ddd")
	println(a.containerName)
	println(a.output)
	println(a.taskPath)
	return nil
}
