package ecs

import (
	"ecs-task-def-action/pkg/encoder"
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type EcsContainer struct{}

func NewEcsContainer() encoder.EcsContainerEncoder {
	return EcsContainer{}
}

func (t EcsContainer) Encode(in []byte, format encoder.Format) *[]ecs.ContainerDefinition {
	var err error
	var def *[]ecs.ContainerDefinition
	switch format {
	case encoder.Json:
		def, err = t.doJson(in)
	case encoder.Yaml:
		def, err = t.doYaml(in)
	}
	if err != nil {
		panic(err)
	}
	return def
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
