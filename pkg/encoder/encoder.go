package encoder

import "ecs-task-def-action/pkg/plovider/ecs"

type Format int

const (
	Json Format = iota
	Yaml
	Unknow
)

func GetFormat(ext string) Format {
	switch ext {
	case ".json":
		return Json
	case ".yaml":
		return Yaml
	case ".yml":
		return Yaml
	default:
		return Unknow
	}
}

type EcsTaskEncoder interface {
	Encode(in []byte, format Format) (*ecs.TaskDefinition, error)
}

type EcsContainerEncoder interface {
	Encode(in []byte, format Format) (*[]ecs.ContainerDefinition, error)
}
