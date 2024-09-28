package main

import (
	"log"

	"github.com/naonao2323/ecs-task-def/pkg/cli"
)

func main() {
	cli := cli.NewCommand()
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
