package v1beta1

type BaseSpec struct {
	Name    string            `json:"name,omitempty"`
	Image   string            `json:"image,omitempty"`
	EnvVars map[string]string `json:"envvars,omitempty"`
	Labels  map[string]string `json:"labels,omitempty"`
}
