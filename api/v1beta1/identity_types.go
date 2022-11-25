package v1beta1

type AWSIdentity struct {
	IAMRole string `json:"iam_role,omitempty"`
}

type IdentitySpec struct {
	AWS AWSIdentity `json:"aws,omitempty"`
}
