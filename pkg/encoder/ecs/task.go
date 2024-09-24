package ecs

import (
	"ecs-task-def-action/pkg/encoder"
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type EcsTask struct {
	logger *zap.Logger
}

func NewEcsTask(logger *zap.Logger) encoder.EcsTaskEncoder {
	return EcsTask{logger: logger}
}

func (t EcsTask) Encode(in []byte, format encoder.Format) (*ecs.TaskDefinition, error) {
	switch format {
	case encoder.Json:
		def, err := t.doJson(in)
		if err != nil {
			t.logger.Error("fail to encode json file", zap.Error(err))
			return nil, err
		}
		return def, nil
	case encoder.Yaml:
		def, err := t.doYaml(in)
		if err != nil {
			t.logger.Error("fail to encode yaml file", zap.Error(err))
			return nil, errors.New("fail to encode yaml file")
		}
		return def, nil
	default:
		t.logger.Warn("unknown extension")
		return nil, errors.New("unknown extension")
	}
}

func (t EcsTask) doJson(in []byte) (*ecs.TaskDefinition, error) {
	var definition ecs.TaskDefinition
	err := json.Unmarshal(in, &definition)
	if err != nil {
		return nil, err
	}
	return &definition, nil
}

func (t EcsTask) doYaml(in []byte) (*ecs.TaskDefinition, error) {
	var definition ecs.TaskDefinition
	err := yaml.Unmarshal(in, &definition)
	if err != nil {
		return nil, err
	}
	return &definition, nil
}
