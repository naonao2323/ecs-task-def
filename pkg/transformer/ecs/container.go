package ecs

import (
	"ecs-task-def-action/pkg/plovider/ecs"
	"ecs-task-def-action/pkg/transformer"
	"fmt"
	"strings"
)

type EcsContainerTransformer struct{}

func NewEcsContainerTransformer() transformer.EcsContainerTransformer {
	return &EcsContainerTransformer{}
}

func (c *EcsContainerTransformer) Transform(tag string, appName string, definition []ecs.ContainerDefinition) []ecs.ContainerDefinition {
	for i, v := range definition {
		if appName == v.Name {
			splited := strings.Split(definition[i].Image, ":")
			definition[i].Image = fmt.Sprintf("%s:%s", splited[0], tag)
		}
	}
	return definition
}
