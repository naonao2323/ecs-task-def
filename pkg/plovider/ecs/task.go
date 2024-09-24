package ecs

type ContainerDefinition struct {
	Name              string `json:"name" yaml:"name"`
	Image             string `json:"image" yaml:"image"`
	Cpu               int    `json:"cpu" yaml:"cpu"`
	Memory            int    `json:"memory,omitempty" yaml:"memory,omitempty"`
	MemoryReservation int    `json:"memoryReservation,omitempty" yaml:"memoryReservation,omitempty"`
	PortMappings      []struct {
		ContainerPort      int    `json:"containerPort" yaml:"containerPort"`
		Protocol           string `json:"protocol" yaml:"protocol"`
		AppProtocol        string `json:"appProtocol,omitempty" yaml:"appProtocol,omitempty"`
		ContainerPortRange string `json:"containerPortRange,omitempty" yaml:"containerPortRange,omitempty"`
		HostPortRange      string `json:"hostPortRange,omitempty" yaml:"hostPortRange,omitempty"`
		HostPort           int    `json:"hostPort,omitempty" yaml:"hostPort,omitempty"`
		Name               string `json:"name,omitempty" yaml:"name,omitempty"`
	} `json:"portMappings,omitempty" yaml:"portMappings,omitempty"`
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
	RepositoryCredentials *struct {
		CredentialsParameter *string `json:"credentialsParameter,omitempty" yaml:"credentialsParameter,omitempty"`
	} `json:"repositoryCredentials,omitempty" yaml:"repositoryCredentials,omitempty"`
	RestartPolicy struct {
		Enabled              bool  `json:"enabled" yaml:"enabled"`
		IgnoredExitCodes     []int `json:"ignoredExitCodes,omitempty" yaml:"ignoredExitCodes,omitempty"`
		RestartAttemptPeriod int   `json:"restartAttemptPeriod,omitempty" yaml:"restartAttemptPeriod,omitempty"`
	} `json:"restartPolicy,omitempty" yaml:"restartPolicy,omitempty"`
	HealthCheck *struct{} `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"` // TODO
}

type TaskDefinition struct {
	TaskDefinitionArn    string                `json:"taskDefinitionArn,omitempty" yaml:"taskDefinitionArn,omitempty"`
	ContainerDefinitions []ContainerDefinition `json:"containerDefinitions" yaml:"containerDefinitions"`
	Family               string                `json:"family" yaml:"family"`
	ExecutionRoleArn     string                `json:"executionRoleArn" yaml:"executionRoleArn"`
	NetWorkMode          string                `json:"networkMode,omitempty" yaml:"networkMode,omitempty"`
	Revision             int                   `json:"revision" yaml:"revision"`
	Volumes              []struct{}            `json:"volumes" yaml:"volumes"`
	Status               string                `json:"status" yaml:"status"`
	RequiresAttributes   []struct {
		Name string `json:"name" yaml:"name"`
	} `json:"requiresAttributes" yaml:"requiresAttributes"`
	PlacementConstraints    []struct{} `json:"placementConstraints" yaml:"placementConstraints"`
	Compatibilities         []string   `json:"compatibilities" yaml:"compatibilities"`
	Cpu                     string     `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory                  string     `json:"memory,omitempty" yaml:"memory,omitempty"`
	RegisteredAt            string     `json:"registeredAt" yaml:"registeredAt"`
	RegisteredBy            string     `json:"registeredBy" yaml:"registeredBy"`
	Tags                    []string   `json:"tags" yaml:"tags"`
	RequiresCompatibilities []string   `json:"requiresCompatibilities,omitempty" yaml:"requiresCompatibilities,omitempty"`
	OperatingSystemFamily   string     `json:"operatingSystemFamily,omitempty" yaml:"operatingSystemFamily,omitempty"`
	CpuArchitecture         string     `json:"cpuArchitecture,omitempty" yaml:"cpuArchitecture,omitempty"`
}
