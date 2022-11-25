package v1beta1

type Autoscale struct {
	Enabled                           bool `json:"enabled,omitempty"`
	MaxReplicas                       int  `json:"max_replicas,omitempty"`
	TargetCPUUtilizationPercentage    int  `json:"target_cpu_utilization_percentage,omitempty"`
	TargetMemoryUtilizationPercentage int  `json:"target_memory_utilization_percentage,omitempty"`
}

type CapacitySpec struct {
	RolloutMaxUnavailable  string    `json:"rolloutmaxunavailable,omitempty"`
	EvictionMaxUnavailable string    `json:"evictionmaxunavailable,omitempty"`
	Replicas               int       `json:"replicas,omitempty"`
	Autoscale              Autoscale `json:"autoscale,omitempty"`
}
