package encoder

import (
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"
	"errors"

	"gopkg.in/yaml.v2"
)

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

type (
	Encoder[P ecs.EcsTarget] interface {
		Encode(in []byte, format Format) (*P, error)
	}
	EncoderImpl[P ecs.EcsTarget] struct{}
)

func NewEncoder[P ecs.EcsTarget]() Encoder[P] {
	return EncoderImpl[P]{}
}

func (e EncoderImpl[P]) Encode(in []byte, format Format) (*P, error) {
	switch format {
	case Json:
		def, err := EncodeJson[P](in)
		if err != nil {
			return nil, errors.New("fail to encode json file")
		}
		return def, nil
	case Yaml:
		def, err := EncodeYaml[P](in)
		if err != nil {
			return nil, errors.New("fail to encode yaml file")
		}
		return def, nil
	default:
		return nil, errors.New("unknown extension")
	}
}

func EncodeJson[P ecs.EcsTarget](in []byte) (*P, error) {
	var def P
	err := json.Unmarshal(in, &def)
	if err != nil {
		return nil, err
	}
	return &def, nil
}

func EncodeYaml[P ecs.EcsTarget](in []byte) (*P, error) {
	var def P
	err := yaml.Unmarshal(in, &def)
	if err != nil {
		return nil, err
	}
	return &def, nil
}
