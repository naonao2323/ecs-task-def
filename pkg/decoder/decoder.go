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

type (
	DecoderImpl[P ecs.EcsTarget] struct{}
	Decoder[P ecs.EcsTarget]     interface {
		Decode(definition P, format Format) ([]byte, error)
	}
)

func NewDecoderImpl[P ecs.EcsTarget]() Decoder[P] {
	return DecoderImpl[P]{}
}

func (d DecoderImpl[P]) Decode(definition P, format Format) ([]byte, error) {
	switch format {
	case Json:
		v, err := d.decodeJson(definition)
		if err != nil {
			return nil, err
		}
		return v, nil
	case Yaml:
		v, err := d.decodeYaml(definition)
		if err != nil {
			return nil, err
		}
		return v, nil
	default:
		return nil, errors.New("unknown file extension")
	}
}

func (d DecoderImpl[P]) decodeJson(definition P) ([]byte, error) {
	v, err := json.MarshalIndent(definition, "", "  ")
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (d DecoderImpl[P]) decodeYaml(definition P) ([]byte, error) {
	v, err := yaml.Marshal(definition)
	if err != nil {
		return nil, err
	}
	return v, nil
}
