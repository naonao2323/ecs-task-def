package ecs

import (
	"ecs-task-def-action/pkg/decoder"
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type EcsTaskDecoder struct {
	logger *zap.Logger
}

func NewEcsTaskDecoder(logger *zap.Logger) decoder.EcsTaskDecoder {
	return &EcsTaskDecoder{
		logger: logger,
	}
}

func (d *EcsTaskDecoder) Decode(definition ecs.TaskDefinition, format decoder.Format) ([]byte, error) {
	switch format {
	case decoder.Json:
		v, err := d.doJson(definition)
		if err != nil {
			d.logger.Error("fail to decode json file", zap.Error(err))
			return nil, err
		}
		return v, nil
	case decoder.Yaml:
		v, err := d.doYaml(definition)
		if err != nil {
			d.logger.Error("fail to decode yaml file", zap.Error(err))
			return nil, err
		}
		return v, nil
	default:
		return nil, errors.New("unknown file extension")
	}
}

func (d *EcsTaskDecoder) doJson(definition ecs.TaskDefinition) ([]byte, error) {
	v, err := json.MarshalIndent(definition, "", "  ")
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (d *EcsTaskDecoder) doYaml(definition ecs.TaskDefinition) ([]byte, error) {
	v, err := yaml.Marshal(definition)
	if err != nil {
		return nil, err
	}
	return v, nil
}
