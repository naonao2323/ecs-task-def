package main

import (
	"ecs-task-def-action/pkg/cli"
	"log"
)

func main() {
	cli := cli.NewCommand()
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
