package ecs

type ContainerDefinition struct {
	Name              string `json:"name" yaml:"name"`
	Image             string `json:"image" yaml:"image"`
	Cpu               int    `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Gpu               int    `json:"gpu,omitempty" yaml:"gpu,omitempty"`
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
	Essential   bool `json:"essential" yaml:"essential"`
	Environment *[]struct {
		Name  string `json:"name" yaml:"name"`
		Value string `json:"value" yaml:"value"`
	} `json:"environment,omitempty" yaml:"environment,omitempty"`
	MountPoints *[]struct {
		SourceVolume  string `json:"sourceVolume" yaml:"sourceVolume"`
		ContainerPath string `json:"containerPath" yaml:"containerPath"`
		ReadOnly      bool   `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	} `json:"mountPoints,omitempty" yaml:"mountPoints,omitempty"`
	VolumesFrom *[]struct {
		SourceContainer string `json:"sourceContainer" yaml:"sourceContainer"`
		ReadOnly        bool   `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	} `json:"volumesFrom,omitempty" yaml:"volumesFrom,omitempty"`
	LogConfiguration *struct {
		LogDriver     string             `json:"logDriver" yaml:"logDriver"`
		Options       *map[string]string `json:"options,omitempty" yaml:"options,omitempty"`
		SecretOptions *struct {
			Name      string `json:"name" yaml:"name"`
			ValueFrom string `json:"valueFrom" yaml:"valueFrom"`
		} `json:"secretOptions,omitempty" yaml:"secretOptions,omitempty"`
	} `json:"logConfiguration,omitempty" yaml:"logConfiguration,omitempty"`
	RepositoryCredentials *struct {
		CredentialsParameter *string `json:"credentialsParameter,omitempty" yaml:"credentialsParameter,omitempty"`
	} `json:"repositoryCredentials,omitempty" yaml:"repositoryCredentials,omitempty"`
	RestartPolicy struct {
		Enabled              bool  `json:"enabled" yaml:"enabled"`
		IgnoredExitCodes     []int `json:"ignoredExitCodes,omitempty" yaml:"ignoredExitCodes,omitempty"`
		RestartAttemptPeriod int   `json:"restartAttemptPeriod,omitempty" yaml:"restartAttemptPeriod,omitempty"`
	} `json:"restartPolicy,omitempty" yaml:"restartPolicy,omitempty"`
	HealthCheck *struct {
		Command     []string `json:"command,omitempty" yaml:"command,omitempty"`
		Interval    int      `json:"interval,omitempty" yaml:"interval,omitempty"`
		Timeout     int      `json:"timeout,omitempty" yaml:"timeout,omitempty"`
		Retries     int      `json:"retries,omitempty" yaml:"retries,omitempty"`
		StartPeriod int      `json:"startPeriod,omitempty" yaml:"startPeriod,omitempty"`
	} `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`
	EntryPoint       []string `json:"entryPoint,omitempty" yaml:"entryPoint,omitempty"`
	Command          []string `json:"command,omitempty" yaml:"command,omitempty"`
	WorkingDirectory string   `json:"workingDirectory,omitempty" yaml:"workingDirectory,omitempty"`
	EnvironmentFiles *[]struct {
		Value string `json:"value" yaml:"value"`
		Type  string `json:"type" yaml:"type"`
	} `json:"environmentFiles,omitempty" yaml:"environmentFiles,omitempty"`
	Secrets *[]struct {
		Name      string `json:"name" yaml:"name"`
		ValueFrom string `json:"valueFrom" yaml:"valueFrom"`
	} `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	DisableNetworking bool     `json:"disableNetworking,omitempty" yaml:"disableNetworking,omitempty"`
	Links             []string `json:"links,omitempty" yaml:"links,omitempty"`
	Hostname          string   `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	DnsServers        []string `json:"dnsServers,omitempty" yaml:"dnsServers,omitempty"`
	DnsSearchDomains  []string `json:"dnsSearchDomains,omitempty" yaml:"dnsSearchDomains,omitempty"`
	ExtraHosts        *[]struct {
		IpAddress string `json:"ipAddress" yaml:"ipAddress"`
		Hostname  string `json:"hostname" yaml:"hostname"`
	} `json:"extraHosts,omitempty" yaml:"extraHosts,omitempty"`
	ReadonlyRootFilesystem bool `json:"readonlyRootFilesystem,omitempty" yaml:"readonlyRootFilesystem,omitempty"`
	FirelensConfiguration  *[]struct {
		Type   string             `json:"type" yaml:"type"`
		Option *map[string]string `json:"option,omitempty" yaml:"option,omitempty"`
	} `json:"firelensConfiguration,omitempty" yaml:"firelensConfiguration,omitempty"`
	CredentialSpecs       []string `json:"credentialSpecs,omitempty" yaml:"credentialSpecs,omitempty"`
	Privileged            bool     `json:"privileged,omitempty" yaml:"privileged,omitempty"`
	User                  string   `json:"user,omitempty" yaml:"user,omitempty"`
	DockerSecurityOptions []string `json:"dockerSecurityOptions,omitempty" yaml:"dockerSecurityOptions,omitempty"`
	Ulimits               *[]struct {
		Name      string `json:"name" yaml:"name"`
		HardLimit int    `json:"hardLimit" yaml:"hardLimit"`
		SoftLimit int    `json:"softLimit" yaml:"softLimit"`
	} `json:"ulimits,omitempty" yaml:"ulimits,omitempty"`
	DockerLabels    *map[string]string `json:"dockerLabels,omitempty" yaml:"dockerLabels,omitempty"`
	LinuxParameters *[]struct {
		Capabilities *struct {
			Add  []string `json:"add,omitempty" yaml:"add,omitempty"`
			Drop []string `json:"drop,omitempty" yaml:"drop,omitempty"`
		} `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
		Devices *[]struct {
			HostPath      string   `json:"hostPath" yaml:"hostPath"`
			ContainerPath string   `json:"containerPath,omitempty" yaml:"containerPath,omitempty"`
			Permissions   []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`
		} `json:"devices,omitempty" yaml:"devices,omitempty"`
		InitProcessEnabled bool `json:"initProcessEnabled,omitempty" yaml:"initProcessEnabled,omitempty"`
		MaxSwap            int  `json:"maxSwap,omitempty" yaml:"maxSwap,omitempty"`
		SharedMemorySize   int  `json:"sharedMemorySize,omitempty" yaml:"sharedMemorySize,omitempty"`
		Swappiness         int  `json:"swappiness,omitempty" yaml:"swappiness,omitempty"`
		Tmpfs              *[]struct {
			ContainerPath string   `json:"containerPath" yaml:"containerPath"`
			MountOptions  []string `json:"mountOptions,omitempty" yaml:"mountOptions,omitempty"`
			Size          int      `json:"size,omitempty" yaml:"size,omitempty"`
		} `json:"tmpfs,omitempty" yaml:"tmpfs,omitempty"`
	} `json:"linuxParameters,omitempty" yaml:"linuxParameters,omitempty"`
	DependsOn *[]struct {
		ContainerName string `json:"containerName" yaml:"containerName"`
		Condition     string `json:"condition" yaml:"condition"`
	} `json:"dependsOn,omitempty" yaml:"dependsOn,omitempty"`
	StartTimeout   int `json:"startTimeout,omitempty" yaml:"startTimeout,omitempty"`
	StopTimeout    int `json:"stopTimeout,omitempty" yaml:"stopTimeout,omitempty"`
	SystemControls *[]struct {
		Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
		Value     string `json:"value,omitempty" yaml:"value,omitempty"`
	} `json:"systemControls,omitempty" yaml:"systemControls,omitempty"`
	Interactive    bool `json:"interactive,omitempty" yaml:"interactive,omitempty"`
	PseudoTerminal bool `json:"pseudoTerminal,omitempty" yaml:"pseudoTerminal,omitempty"`
}

type TaskDefinition struct {
	TaskDefinitionArn    string                `json:"taskDefinitionArn,omitempty" yaml:"taskDefinitionArn,omitempty"`
	ContainerDefinitions []ContainerDefinition `json:"containerDefinitions" yaml:"containerDefinitions"`
	Family               string                `json:"family" yaml:"family"`
	ExecutionRoleArn     string                `json:"executionRoleArn" yaml:"executionRoleArn"`
	NetWorkMode          string                `json:"networkMode,omitempty" yaml:"networkMode,omitempty"`
	Revision             int                   `json:"revision" yaml:"revision"`
	Volumes              *[]struct {
		Name string `json:"name,omitempty" yaml:"name,omitempty"`
		Host *struct {
			SourcePath string `json:"sourcePath,omitempty" yaml:"sourcePath,omitempty"`
		} `json:"host,omitempty" yaml:"host,omitempty"`
		ConfiguredAtLaunch        bool `json:"configuredAtLaunch,omitempty" yaml:"configuredAtLaunch,omitempty"`
		DockerVolumeConfiguration struct {
			Scope         string `json:"scope,omitempty" yaml:"scope,omitempty"`
			Autoprovision bool   `json:"autoprovision,omitempty" yaml:"autoprovision,omitempty"`
			Driver        string `json:"driver,omitempty" yaml:"driver,omitempty"`
			DriverOpts    struct {
				KeyName string `json:"keyName,omitempty" yaml:"keyName,omitempty"`
			} `json:"driverOpts,omitempty" yaml:"driverOpts,omitempty"`
			Labels struct {
				KeyName string `json:"keyName,omitempty" yaml:"keyName,omitempty"`
			} `json:"labels,omitempty" yaml:"labels,omitempty"`
		} `json:"dockerVolumeConfiguration,omitempty" yaml:"dockerVolumeConfiguration,omitempty"`
		EfsVolumeConfiguration struct {
			FileSystemId          string `json:"fileSystemId" yaml:"fileSystemId"`
			RootDirectory         string `json:"rootDirectory,omitempty" yaml:"rootDirectory,omitempty"`
			TransitEncryption     string `json:"transitEncryption,omitempty" yaml:"transitEncryption,omitempty"`
			TransitEncryptionPort string `json:"transitEncryptionPort,omitempty" yaml:"transitEncryptionPort,omitempty"`
			AuthorizationConfig   struct {
				AccessPointId string `json:"accessPointId,omitempty" yaml:"accessPointId,omitempty"`
				Iam           string `json:"iam,omitempty" yaml:"iam,omitempty"`
			} `json:"authorizationConfig,omitempty" yaml:"authorizationConfig,omitempty"`
		} `json:"efsVolumeConfiguration,omitempty" yaml:"efsVolumeConfiguration,omitempty"`
		FsxWindowsFileServerVolumeConfiguration struct {
			FileSystemId        string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`
			RootDirectory       string `json:"rootDirectory,omitempty" yaml:"rootDirectory,omitempty"`
			AuthorizationConfig struct {
				CredentialsParameter string `json:"credentialsParameter,omitempty" yaml:"credentialsParameter,omitempty"`
				Domain               string `json:"domain,omitempty" yaml:"domain,omitempty"`
			} `json:"authorizationConfig,omitempty" yaml:"authorizationConfig,omitempty"`
		} `json:"fsxWindowsFileServerVolumeConfiguration" yaml:"fsxWindowsFileServerVolumeConfiguration"`
	} `json:"volumes,omitempty" yaml:"volumes,omitempty"`
	Status             string `json:"status" yaml:"status"`
	RequiresAttributes []struct {
		Name string `json:"name" yaml:"name"`
	} `json:"requiresAttributes" yaml:"requiresAttributes"`
	PlacementConstraints *[]struct {
		Type       string `json:"type,omitempty" yaml:"type,omitempty"`
		Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`
	} `json:"placementConstraints,omitempty" yaml:"placementConstraints,omitempty"`
	Compatibilities []string `json:"compatibilities" yaml:"compatibilities"`
	Cpu             string   `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory          string   `json:"memory,omitempty" yaml:"memory,omitempty"`
	RegisteredAt    string   `json:"registeredAt" yaml:"registeredAt"`
	RegisteredBy    string   `json:"registeredBy" yaml:"registeredBy"`
	Tags            *[]struct {
		Key   string `json:"key,omitempty" yaml:"key,omitempty"`
		Value string `json:"value,omitempty" yaml:"value,omitempty"`
	} `json:"tags,omitempty" yaml:"tags,omitempty"`
	RequiresCompatibilities []string `json:"requiresCompatibilities,omitempty" yaml:"requiresCompatibilities,omitempty"`
	OperatingSystemFamily   string   `json:"operatingSystemFamily,omitempty" yaml:"operatingSystemFamily,omitempty"`
	CpuArchitecture         string   `json:"cpuArchitecture,omitempty" yaml:"cpuArchitecture,omitempty"`
	EphemeralStorage        *struct {
		SizeInGiB int `json:"sizeInGiB,omitempty" yaml:"sizeInGiB,omitempty"`
	} `json:"ephemeralStorage,omitempty" yaml:"ephemeralStorage,omitempty"`
	PidMode               string `json:"pidMode,omitempty" yaml:"pidMode,omitempty"`
	IpcMode               string `json:"ipcMode,omitempty" yaml:"ipcMode,omitempty"`
	InferenceAccelerators *[]struct {
		DeviceName string `json:"deviceName" yaml:"deviceName"`
		DeviceType string `json:"deviceType" yaml:"deviceType"`
	} `json:"inferenceAccelerators,omitempty" yaml:"inferenceAccelerators,omitempty"`
	ProxyConfiguration *struct {
		Type          string `json:"type,omitempty" yaml:"type,omitempty"`
		ContainerName string `json:"containerName" yaml:"containerName"`
		Properties    *[]struct {
			Name  string `json:"name,omitempty" yaml:"name,omitempty"`
			Value string `json:"value,omitempty" yaml:"value,omitempty"`
		} `json:"properties,omitempty" yaml:"properties,omitempty"`
	} `json:"proxyConfiguration,omitempty" yaml:"proxyConfiguration,omitempty"`
	RuntimePlatform *struct {
		CpuArchitecture       string `json:"cpuArchitecture,omitempty" yaml:"cpuArchitecture,omitempty"`
		OperatingSystemFamily string `json:"operatingSystemFamily,omitempty" yaml:"operatingSystemFamily,omitempty"`
	} `json:"runtimePlatform,omitempty" yaml:"runtimePlatform,omitempty"`
}

type EcsTarget interface {
	TaskDefinition | []ContainerDefinition
}
