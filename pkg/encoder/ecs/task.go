package ecs

import (
	"ecs-task-def-action/pkg/encoder"
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type EcsTask struct{}

func NewEcsTask() encoder.EcsTaskEncoder {
	return EcsTask{}
}

func (t EcsTask) Encode(in []byte, format encoder.Format) *ecs.TaskDefinition {
	var err error
	var def *ecs.TaskDefinition
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
	fmt.Println(definition)
	return &definition, nil
}
