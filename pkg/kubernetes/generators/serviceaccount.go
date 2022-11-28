package generators

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"migueleliasweb.github.io/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ServiceAccount struct {
	Client client.Client
	Scheme *runtime.Scheme
}

func (SA *ServiceAccount) GenerateFromApp(app v1beta1.App) (*corev1.ServiceAccount, error) {
	return &corev1.ServiceAccount{}, nil
}
