package secondary_reconcilers

import (
	"context"

	platformv1beta1 "migueleliasweb.github.io/api/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
)

type ReconcilerWithApp interface {
	ReconcileWithApp(
		ctx context.Context,
		req ctrl.Request,
		app *platformv1beta1.App,
	) error

	DeleteWithApp(
		ctx context.Context,
		req ctrl.Request,
		app *platformv1beta1.App,
	) error

	ValidateWithApp(
		ctx context.Context,
		req ctrl.Request,
		app *platformv1beta1.App,
	) error
}
