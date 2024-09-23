package ecs

import (
	"ecs-task-def-action/pkg/decoder"
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type EcsContainerDecoder struct{}

func NewEcsContainerDecoder() decoder.EcsContainerDecoder {
	return &EcsContainerDecoder{}
}

func (d *EcsContainerDecoder) Decode(definition []ecs.ContainerDefinition, format decoder.Format) []byte {
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
