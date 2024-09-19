package ecs

import (
	"ecs-task-def-action/pkg/plovider/ecs"
	"ecs-task-def-action/pkg/transformer"
)

type TaskTransformer struct {
	definition ecs.TaskDefinition
}

func NewTransformer(definition ecs.TaskDefinition) transformer.Transformer {
	return TaskTransformer{definition: definition}
}

func (t TaskTransformer) Transform(tag string, appName string) {
	t.setImageTag(tag, appName)
}

func (t TaskTransformer) setImageTag(tag string, appName string) {
	for i, v := range t.definition.ContainerDefinitions {
		if appName == v.Name {
			t.definition.ContainerDefinitions[i].Image = tag
		}
	}
}
