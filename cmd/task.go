package main

import (
	"ecs-task-def/pkg/cli"
	"log"
)

func main() {
	cli := cli.NewCommand()
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
