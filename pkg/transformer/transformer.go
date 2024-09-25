package transformer

import (
	"ecs-task-def-action/pkg/plovider/ecs"
	"fmt"
	"strings"
)

type (
	TransformerImpl[P ecs.EcsTarget] struct{}
	Transformer[P ecs.EcsTarget]     interface {
		Transform(tag, appName string, definition P) P
	}
)

func NewTransformer[P ecs.EcsTarget]() Transformer[P] {
	return TransformerImpl[P]{}
}

func (t TransformerImpl[P]) Transform(tag, appName string, definition P) P {
	switch d := any(definition).(type) {
	case ecs.TaskDefinition:
		transformed := t.transFormTaskDefinition(tag, appName, d)
		return any(transformed).(P)
	case []ecs.ContainerDefinition:
		transformed := t.transFormContainerDefinition(tag, appName, d)
		return any(transformed).(P)
	default:
		return any(d).(P)
	}
}

func (t TransformerImpl[P]) transFormTaskDefinition(tag, appName string, definition ecs.TaskDefinition) ecs.TaskDefinition {
	for i, v := range definition.ContainerDefinitions {
		if appName == v.Name {
			splited := strings.Split(definition.ContainerDefinitions[i].Image, ":")
			definition.ContainerDefinitions[i].Image = fmt.Sprintf("%s:%s", splited[0], tag)
		}
	}
	return definition
}

func (t TransformerImpl[P]) transFormContainerDefinition(tag, appName string, definition []ecs.ContainerDefinition) []ecs.ContainerDefinition {
	for i, v := range definition {
		if appName == v.Name {
			splited := strings.Split(definition[i].Image, ":")
			definition[i].Image = fmt.Sprintf("%s:%s", splited[0], tag)
		}
	}
	return definition
}
