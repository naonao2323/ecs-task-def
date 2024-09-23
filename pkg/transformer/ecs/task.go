package ecs

import (
	"ecs-task-def-action/pkg/plovider/ecs"
	"ecs-task-def-action/pkg/transformer"
	"fmt"
	"strings"
)

type EcsTaskTransformer struct {
}

func NewTaskTransformer() transformer.EcsTaskTransformer {
	return &EcsTaskTransformer{}
}

func (t *EcsTaskTransformer) Transform(tag string, appName string, definition ecs.TaskDefinition) ecs.TaskDefinition {
	for i, v := range definition.ContainerDefinitions {
		if appName == v.Name {
			splited := strings.Split(definition.ContainerDefinitions[i].Image, ":")
			definition.ContainerDefinitions[i].Image = fmt.Sprintf("%s:%s", splited[0], tag)
		}
	}
	return definition
}
