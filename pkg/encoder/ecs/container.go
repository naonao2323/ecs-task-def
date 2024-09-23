package ecs

import (
	"ecs-task-def-action/pkg/encoder"
	"ecs-task-def-action/pkg/plovider/ecs"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

// エラーハンドリング
// getterの抽象化する
type EcsContainer struct {
}

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

// func (t Container) GetImageTag(appName string) string {
// 	var tag string
// 	for _, v := range t.definition {
// 		if appName == v.Name {
// 			tag = strings.Split(v.Image, ":")[1]
// 		}
// 	}
// 	return tag
// }

func (c EcsContainer) doJson(in []byte) (*[]ecs.ContainerDefinition, error) {
	var def []ecs.ContainerDefinition
	err := json.Unmarshal(in, &def)
	if err != nil {
		return nil, err
	}
	fmt.Println(def)
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
