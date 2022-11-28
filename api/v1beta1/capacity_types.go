package v1beta1

type AutoscalingSpec struct {
	Enabled                           bool `json:"enabled,omitempty"`
	MaxReplicas                       int  `json:"max_replicas,omitempty"`
	TargetCPUUtilizationPercentage    int  `json:"target_cpu_utilization_percentage,omitempty"`
	TargetMemoryUtilizationPercentage int  `json:"target_memory_utilization_percentage,omitempty"`
}

type ResourcesSpec struct {
	LimitCPU    string `json:"limit_cpu,omitempty"`
	LimitMEM    string `json:"limit_mem,omitempty"`
	RequestsCPU string `json:"requests_cpu,omitempty"`
	RequestsMEM string `json:"requests_mem,omitempty"`
}

type CapacitySpec struct {
	RolloutMaxUnavailable  string          `json:"rollout_max_unavailable,omitempty"`
	EvictionMaxUnavailable string          `json:"eviction_max_unavailable,omitempty"`
	Replicas               int             `json:"replicas,omitempty"`
	Autoscaling            AutoscalingSpec `json:"autoscaling,omitempty"`
	Resources              ResourcesSpec   `json:"resourcesspec,omitempty"`
}
