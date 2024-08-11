package controller

import (
	"context"

	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/jwcesign/kloud/pkg/apis/cluster/v1alpha1"
)

type Controller struct {
	kubeClient client.Client
}

func SetupController(mgr manager.Manager) error {
	c, err := controller.New("analyzer-controller", mgr,
		controller.Options{Reconciler: &Controller{kubeClient: mgr.GetClient()}})
	if err != nil {
		return err
	}
	err = c.Watch(source.Kind(mgr.GetCache(),
		&v1alpha1.ClusterMigration{}, &handler.TypedEnqueueRequestForObject[*v1alpha1.ClusterMigration]{}))
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	klog.Info("Start to reconcile ClusterMigration:%s...", req)
	return reconcile.Result{}, nil
}
