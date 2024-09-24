package decoder

import "ecs-task-def-action/pkg/plovider/ecs"

type Format int

const (
	Json Format = iota
	Yaml
)

type EcsTaskDecoder interface {
	Decode(definition ecs.TaskDefinition, format Format) ([]byte, error)
}

type EcsContainerDecoder interface {
	Decode(definition []ecs.ContainerDefinition, format Format) ([]byte, error)
}
