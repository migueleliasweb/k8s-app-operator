package v1beta1

type AWSIdentity struct {
	IAMRole string `json:"iam_role,omitempty"`
}

// TODO: How to nicely expose this without exposing the whole underlying Kind?
type K8sRBACIdentity struct {
	// Roles
}

type K8sIdentity struct {
	ServiceAccount string          `json:"service_account,omitempty"`
	RBAC           K8sRBACIdentity `json:"rbac,omitempty"`
}

type IdentitySpec struct {
	AWS AWSIdentity `json:"aws,omitempty"`
	K8s K8sIdentity `json:"k8s,omitempty"`
}
