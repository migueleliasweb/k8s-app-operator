/*
Copyright 2022.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1beta1"

	platformv1beta1 "migueleliasweb.github.io/api/v1beta1"
)

const AppFinalizerName = "platform.migueleliasweb.github.io/finalizer"

type SecondaryReconcilers interface {
	ReconcileWithApp(
		ctx context.Context,
		req ctrl.Request,
		app platformv1beta1.App,
	) (ctrl.Result, error)
}

// AppReconciler reconciles a App object
type AppReconciler struct {
	client.Client
	Scheme               *runtime.Scheme
	Log                  *logr.Logger
	SecondaryReconcilers []reconcile.Reconciler
}

// See: https://book.kubebuilder.io/reference/using-finalizers.html
func (r *AppReconciler) handleFinalizer(
	ctx context.Context,
	req ctrl.Request,
	app *platformv1beta1.App,
) error {
	// examine D	eletionTimestamp to determine if object is under deletion
	if app.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !controllerutil.ContainsFinalizer(app, AppFinalizerName) {
			controllerutil.AddFinalizer(app, AppFinalizerName)
			if err := r.Update(ctx, app); err != nil {
				return err
			}
		}
	} else {
		// The object is being deleted
		if controllerutil.ContainsFinalizer(app, AppFinalizerName) {
			// our finalizer is present, so lets handle any external dependency
			if err := r.deleteExternalResources(app); err != nil {
				// if fail to delete the external dependency here, return with error
				// so that it can be retried
				return err
			}

			// remove our finalizer from the list and update it.
			controllerutil.RemoveFinalizer(app, AppFinalizerName)
			if err := r.Update(ctx, app); err != nil {
				return err
			}
		}

		// Stop reconciliation as the item is being deleted
		return nil
	}

	return nil
}

func (r *AppReconciler) deleteExternalResources(app *platformv1beta1.App) error {
	// call all downstream reconcilers
	// exit
	return nil
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
	log := r.Log.WithValues("app", req.NamespacedName)

	var app *platformv1beta1.App
	if err := r.Get(ctx, req.NamespacedName, app); err != nil {
		log.Error(err, "unable to fetch CronJob")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if err := r.handleFinalizer(
		ctx,
		req,
		app,
	); err != nil {
		return ctrl.Result{}, err
	}

	// Your reconcile logic

	// fetch current definitions
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
