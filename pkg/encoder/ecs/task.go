package ecs

import (
	"ecs-task-def-action/pkg/encoder"
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

type TaskGetter interface {
	GetImageTag(appName string) string
}

type Task struct {
	definition ecs.TaskDefinition
}

func NewTask() encoder.TaskEncoder {
	return Task{}
}

func (t Task) Encode(in []byte, format encoder.Format) {
	var err error
	switch format {
	case encoder.Json:
		err = t.doJson(in)
	case encoder.Yaml:
		err = t.doYaml(in)
	}
	if err != nil {
		panic(err)
	}
}

func (t Task) GetImageTag(appName string) string {
	var tag string
	for _, v := range t.definition.ContainerDefinitions {
		if appName == v.Name {
			tag = strings.Split(v.Image, ":")[1]
		}
	}
	return tag
}

func (t Task) doJson(in []byte) error {
	err := json.Unmarshal(in, &t.definition)
	if err != nil {
		return err
	}
	fmt.Println(t.definition.ContainerDefinitions[0].Image)
	return nil
}

func (t Task) doYaml(in []byte) error {
	err := yaml.Unmarshal(in, &t.definition)
	if err != nil {
		return err
	}
	fmt.Println(t.definition.ContainerDefinitions[0].Image)
	return nil
}
