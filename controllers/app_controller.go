/*
Copyright 2022.
*/

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1beta1"

	platformv1beta1 "migueleliasweb.github.io/api/v1beta1"
)

type SecondaryReconcilers interface {
	ReconcileWithApp(
		ctx context.Context,
		req ctrl.Request,
		app platformv1beta1.App,
	) (ctrl.Result, error)
}

// AppReconciler reconciles a App object
type AppReconciler struct {
	Client               client.Client
	Scheme               *runtime.Scheme
	SecondaryReconcilers []reconcile.Reconciler
}

//+kubebuilder:rbac:groups=platform.migueleliasweb.github.io,resources=apps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=platform.migueleliasweb.github.io,resources=apps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=platform.migueleliasweb.github.io,resources=apps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// fetch current definition
	// set status to "reconciling"?
	// call secondary reconcilers
	// done

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&platformv1beta1.App{}).
		Owns(&appsv1.Deployment{}).
		Owns(&appsv1.ReplicaSet{}).
		Owns(&corev1.Pod{}).
		Owns(&policyv1.PodDisruptionBudget{}).
		Owns(&corev1.ServiceAccount{}).
		// Owns(&corev1.Service{}).
		Complete(r)
}
