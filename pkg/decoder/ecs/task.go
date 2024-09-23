package ecs

import (
	"ecs-task-def-action/pkg/decoder"
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type EcsTaskDecoder struct {
}

func NewEcsTaskDecoder() decoder.EcsTaskDecoder {
	return &EcsTaskDecoder{}
}

// TODO: エラーハンドリング
func (d *EcsTaskDecoder) Decode(definition ecs.TaskDefinition, format decoder.Format) []byte {
	switch format {
	case decoder.Json:
		v, err := d.doJson(definition)
		if err != nil {
			return nil
		}
		return v
	case decoder.Yaml:
		v, err := d.doYaml(definition)
		if err != nil {
			return nil
		}
		return v
	}
	return nil
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
