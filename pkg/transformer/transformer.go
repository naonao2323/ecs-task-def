package transformer

import "ecs-task-def-action/pkg/plovider/ecs"

type EcsTaskTransformer interface {
	Transform(tag string, appName string, definition ecs.TaskDefinition) ecs.TaskDefinition
}

type EcsContainerTransformer interface {
	Transform(tag string, appName string, definition []ecs.ContainerDefinition) []ecs.ContainerDefinition
}
