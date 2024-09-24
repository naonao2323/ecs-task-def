package ecs

import (
	"ecs-task-def-action/pkg/decoder"
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type EcsContainerDecoder struct {
	logger *zap.Logger
}

func NewEcsContainerDecoder(logger *zap.Logger) decoder.EcsContainerDecoder {
	return &EcsContainerDecoder{
		logger: logger,
	}
}

func (d *EcsContainerDecoder) Decode(definition []ecs.ContainerDefinition, format decoder.Format) ([]byte, error) {
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
		return nil, errors.New("unknown extension file")
	}
}

func (d *EcsContainerDecoder) doJson(definition []ecs.ContainerDefinition) ([]byte, error) {
	v, err := json.MarshalIndent(definition, "", "  ")
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (d *EcsContainerDecoder) doYaml(definition []ecs.ContainerDefinition) ([]byte, error) {
	v, err := yaml.Marshal(definition)
	if err != nil {
		return nil, err
	}
	return v, nil
}
