package secondary_reconcilers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	platformv1beta1 "migueleliasweb.github.io/api/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Deployment struct {
	Client client.Client
	Scheme *runtime.Scheme
}

func (Deployment *Deployment) ReconcileWithApp(
	ctx context.Context,
	req ctrl.Request,
	app *platformv1beta1.App,
) error {
	return nil
}

func (Deployment *Deployment) DeleteWithApp(
	ctx context.Context,
	req ctrl.Request,
	app *platformv1beta1.App,
) error {
	return nil
}

func (Deployment *Deployment) ValidateWithApp(
	ctx context.Context,
	req ctrl.Request,
	app *platformv1beta1.App,
) error {
	return nil
}
