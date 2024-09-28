package decoder

import (
	"ecs-task-def/pkg/plovider/ecs"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Format int

const (
	Json Format = iota
	Yaml
)

type (
	DecoderImpl[P ecs.EcsTarget] struct {
		logger *zap.Logger
	}
	Decoder[P ecs.EcsTarget] interface {
		Decode(definition P, format Format) ([]byte, error)
	}
)

func NewDecoder[P ecs.EcsTarget](logger *zap.Logger) Decoder[P] {
	return DecoderImpl[P]{logger}
}

func (d DecoderImpl[P]) Decode(definition P, format Format) ([]byte, error) {
	switch format {
	case Json:
		v, err := d.decodeJson(definition)
		if err != nil {
			d.logger.Error("fail to decode json file", zap.Error(err))
			return nil, err
		}
		return v, nil
	case Yaml:
		v, err := d.decodeYaml(definition)
		if err != nil {
			d.logger.Error("fail to decode yaml file", zap.Error(err))
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
