package types

type ChronosJob struct {
	Name                 string                   `json:"name"`
	Command              string                   `json:"command"`
	Schedule             string                   `json:"schedule"`
	Shell                bool                     `json:"shell"`
	Epsilon              string                   `json:"epsilon"`
	Executor             string                   `json:"executor"`
	ExecutorFlags        string                   `json:"executorFlags"`
	Retries              uint64                   `json:"retries"`
	Owner                string                   `json:"owner"`
	OwnerName            string                   `json:"ownerName"`
	Description          string                   `json:"description"`
	Async                bool                     `json:"async"`
	SuccessCount         uint64                   `json:"successCount"`
	ErrorCount           uint64                   `json:"errorCount"`
	LastSuccess          string                   `json:"lastSuccess"`
	LastError            string                   `json:"lastError"`
	CPUs                 float64                  `json:"cpus"`
	Disk                 float64                  `json:"disk"`
	Mem                  float64                  `json:"mem"`
	Disabled             bool                     `json:"disabled"`
	Parents              []string                 `json:"parents"`
	EnvironmentVariables *[]EnvironmentVariable   `json:"environmentVariables,omitempty"`
	Constraints          []string                 `json:"constraints"`
	Arguments            []string                 `json:"arguments"`
	RunAsUser            string                   `json:"runAsUser"`
	Container            *ChronosContainerOptions `json:"container,omitempty"`
}

type EnvironmentVariable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NewChronosJob struct {
	Name                 string                   `json:"name"`
	Command              string                   `json:"command,omitempty"`
	Schedule             string                   `json:"schedule,omitempty"`
	Parents              *[]string                `json:"parents,omitempty"`
	Epsilon              string                   `json:"epsilon,omitempty"`
	Async                bool                     `json:"async,omitempty"`
	Owner                string                   `json:"owner,omitempty"`
	OwnerName            string                   `json:"ownerName,omitempty"`
	Description          string                   `json:"description,omitempty"`
	CPUs                 float64                  `json:"cpus,omitempty"`
	Disk                 float64                  `json:"disk,omitempty"`
	Mem                  float64                  `json:"mem,omitempty"`
	EnvironmentVariables *[]EnvironmentVariable   `json:"environmentVariables,omitempty"`
	Constraints          []string                 `json:"constraints,omitempty"`
	Container            *ChronosContainerOptions `json:"container,omitempty"`
}

type ChronosJobStatus struct {
	Name        string
	LastOutcome string
	Status      string
}

type ChronosContainerOptions struct {
	Type           string `json:"type"`
	Image          string `json:"image"`
	Network        string `json:"network,omitempty"`
	ForcePullImage bool   `json:"forcePullImage,omitempty"`
}
