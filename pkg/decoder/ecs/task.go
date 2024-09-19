package ecs

import (
	"ecs-task-def-action/pkg/decoder"
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type TaskDecoder struct {
	definition ecs.TaskDefinition
}

func NewTaskDecoder(definition ecs.TaskDefinition) decoder.Decorder {
	return TaskDecoder{definition: definition}
}

// TODO: エラーハンドリング
func (d TaskDecoder) Decode(format decoder.Format) []byte {
	switch format {
	case decoder.Json:
		v, err := d.doJson()
		if err != nil {
			return nil
		}
		return v
	case decoder.Yaml:
		v, err := d.doYaml()
		if err != nil {
			return nil
		}
		return v
	}
	return nil
}

func (d TaskDecoder) doJson() ([]byte, error) {
	v, err := json.Marshal(d.definition)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (d TaskDecoder) doYaml() ([]byte, error) {
	v, err := yaml.Marshal(d.definition)
	if err != nil {
		return nil, err
	}
	return v, nil
}
