package ecs

import (
	"ecs-task-def-action/pkg/encoder"
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type EcsContainer struct {
	logger *zap.Logger
}

func NewEcsContainer(logger *zap.Logger) encoder.EcsContainerEncoder {
	return EcsContainer{
		logger: logger,
	}
}

func (t EcsContainer) Encode(in []byte, format encoder.Format) (*[]ecs.ContainerDefinition, error) {
	switch format {
	case encoder.Json:
		def, err := t.doJson(in)
		if err != nil {
			t.logger.Error("fail to encode json file", zap.Error(err))
			return nil, errors.New("fail to encode json file")
		}
		return def, nil
	case encoder.Yaml:
		def, err := t.doYaml(in)
		t.logger.Error("fail to encode yaml file", zap.Error(err))
		if err != nil {
			return nil, errors.New("fail to encode yaml file")
		}
		return def, nil
	default:
		t.logger.Warn("unknow extension")
		return nil, errors.New("unknown extension")
	}
}

func (c EcsContainer) doJson(in []byte) (*[]ecs.ContainerDefinition, error) {
	var def []ecs.ContainerDefinition
	err := json.Unmarshal(in, &def)
	if err != nil {
		return nil, err
	}
	return &def, nil
}

func (c EcsContainer) doYaml(in []byte) (*[]ecs.ContainerDefinition, error) {
	var def []ecs.ContainerDefinition
	err := yaml.Unmarshal(in, &def)
	if err != nil {
		return nil, err
	}
	return &def, nil
}
