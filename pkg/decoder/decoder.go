package decoder

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
)

type DecodeTarget interface {
	ecs.TaskDefinition | []ecs.ContainerDefinition
}

type Decoder[P DecodeTarget] interface {
	Decode(definition P, format Format) ([]byte, error)
}

func Decode[P DecodeTarget](definition P, format Format) ([]byte, error) {
	switch format {
	case Json:
		v, err := decodeJson(definition)
		if err != nil {
			return nil, err
		}
		return v, nil
	case Yaml:
		v, err := decodeYaml(definition)
		if err != nil {
			return nil, err
		}
		return v, nil
	default:
		return nil, errors.New("unknown file extension")
	}
}

func decodeJson[P DecodeTarget](definition P) ([]byte, error) {
	v, err := json.MarshalIndent(definition, "", "  ")
	if err != nil {
		return nil, err
	}
	return v, nil
}

func decodeYaml[P DecodeTarget](definition P) ([]byte, error) {
	v, err := yaml.Marshal(definition)
	if err != nil {
		return nil, err
	}
	return v, nil
}
