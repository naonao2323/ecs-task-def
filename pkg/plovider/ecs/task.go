package ecs

type ContainerDefinition struct {
	Name         string `json:"name" yaml:"name"`
	Image        string `json:"image" yaml:"image"`
	Cpu          int    `json:"cpu" yaml:"cpu"`
	Memory       int    `json:"memory" yaml:"memory"`
	PortMappings []struct {
		ContainerPort int    `json:"containerPort" yaml:"containerPort"`
		Protocol      string `json:"protocol" yaml:"protocol"`
	} `json:"portMappings" yaml:"portMappings"`
	Essential        bool       `json:"essential" yaml:"essential"`
	Environment      []struct{} `json:"environment" yaml:"environment"`
	MountPoints      []struct{} `json:"mountPoints" yaml:"mountPoints"`
	VolumesFrom      []struct{} `json:"volumesFrom" yaml:"volumesFrom"`
	LogConfiguration struct {
		LogDriver string `json:"logDriver" yaml:"logDriver"`
		Options   struct {
			AwslogsGroup        string `json:"awslogs-group" yaml:"awslogs-group"`
			AwslogsRegion       string `json:"awslogs-region" yaml:"awslogs-region"`
			AwslogsStreamPrefix string `json:"awslogs-stream-prefix" yaml:"awslogs-stream-prefix"`
		} `json:"options" yaml:"options"`
	} `json:"logConfiguration" yaml:"logConfiguration"`
}
type TaskDefinition struct {
	TaskDefinitionArn    string                `json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	ContainerDefinitions []ContainerDefinition `json:"containerDefinitions" yaml:"containerDefinitions"`
	Family               string                `json:"family" yaml:"family"`
	ExecutionRoleArn     string                `json:"executionRoleArn" yaml:"executionRoleArn"`
	NetWorkMode          string                `json:"networkMode" yaml:"networkMode"`
	Revision             int                   `json:"revision" yaml:"revision"`
	Volumes              []struct{}            `json:"volumes" yaml:"volumes"`
	Status               string                `json:"status" yaml:"status"`
	RequiresAttributes   []struct {
		Name string `json:"name" yaml:"name"`
	} `json:"requiresAttributes" yaml:"requiresAttributes"`
	PlacementConstraints []struct{} `json:"placementConstraints" yaml:"placementConstraints"`
	Compatibilities      []string   `json:"compatibilities" yaml:"compatibilities"`
	Cpu                  string     `json: "cpu" yaml: "cpu"`
	Memory               string     `json: "memory" yaml: "memory"`
	RegisteredAt         string     `json:"registeredAt" yaml:"registeredAt"`
	RegisteredBy         string     `json:"registeredBy" yaml:"registeredBy"`
	Tags                 []string   `json:"tags" yaml:"tags"`
}
