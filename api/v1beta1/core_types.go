package v1beta1

type CoreSpec struct {
	Name        string            `json:"name"`
	Image       string            `json:"image"`
	EnvVars     map[string]string `json:"envvars,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Command     []string          `json:"command,omitempty"`
	Args        []string          `json:"args,omitempty"`
}
