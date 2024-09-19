package ecs

import (
	"encoding/json"
	"fmt"
)

type TaskEncoder interface {
	Encoder(in []byte)
}

type Task struct {
	definition struct {
		TaskDefinitionArn    string
		ContainerDefinitions []struct {
			Name         string
			Image        string
			Cpu          int
			Memory       int
			PortMappings []struct {
				ContainerPort int
				Protocol      string
			}
			Essential        bool
			Environment      []struct{}
			MountPoints      []struct{}
			VolumesFrom      []struct{}
			LogConfiguration struct {
				LogDriver string
				Options   struct {
					AwslogsGroup        string
					AwslogsRegion       string
					AwslogsStreamPrefix string
				}
			}
		}
		Family             string
		ExecutionRoleArn   string
		NetWorkMode        string
		Revision           int
		Volumes            []struct{}
		Status             string
		RequiresAttributes []struct {
			Name string
		}
		PlacementConstraints []struct{}
		Compatibilities      []string
		RegisteredAt         string
		RegisteredBy         string
		Tags                 []string
	}
}

func NewTask() TaskEncoder {
	return Task{}
}

func (t Task) Encoder(in []byte) {
	err := json.Unmarshal(in, &t.definition)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.definition)
}
