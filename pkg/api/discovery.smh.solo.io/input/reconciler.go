// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconciler.go -destination mocks/reconciler.go

// The Input Reconciler calls a simple func() error whenever a
// storage event is received for any of:
// * ConfigMaps
// * Services
// * Pods
// * Nodes
// * Deployments
// * ReplicaSets
// * DaemonSets
// * StatefulSets
// for a given cluster or set of clusters.
//
// Input Reconcilers can be be constructed from either a single Manager (watch events in a single cluster)
// or a ClusterWatcher (watch events in multiple clusters).
package input

import (
	"context"

	"github.com/solo-io/skv2/contrib/pkg/input"
	sk_core_v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/reconcile"

	"sigs.k8s.io/controller-runtime/pkg/manager"

	v1_controllers "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/controller"
	v1 "k8s.io/api/core/v1"

	apps_v1_controllers "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/controller"
	apps_v1 "k8s.io/api/apps/v1"
)

// the multiClusterReconciler reconciles events for input resources across clusters
type multiClusterReconciler interface {
	v1_controllers.MulticlusterConfigMapReconciler
	v1_controllers.MulticlusterServiceReconciler
	v1_controllers.MulticlusterPodReconciler
	v1_controllers.MulticlusterNodeReconciler

	apps_v1_controllers.MulticlusterDeploymentReconciler
	apps_v1_controllers.MulticlusterReplicaSetReconciler
	apps_v1_controllers.MulticlusterDaemonSetReconciler
	apps_v1_controllers.MulticlusterStatefulSetReconciler
}

var _ multiClusterReconciler = &multiClusterReconcilerImpl{}

type multiClusterReconcilerImpl struct {
	base input.MultiClusterReconciler
}

// register the reconcile func with the cluster watcher
func RegisterMultiClusterReconciler(
	ctx context.Context,
	clusters multicluster.ClusterWatcher,
	reconcileFunc input.MultiClusterReconcileFunc,
) {

	base := input.NewMultiClusterReconcilerImpl(
		ctx,
		reconcileFunc,
	)

	r := &multiClusterReconcilerImpl{
		base: base,
	}

	// initialize reconcile loops

	v1_controllers.NewMulticlusterConfigMapReconcileLoop("ConfigMap", clusters).AddMulticlusterConfigMapReconciler(ctx, r)
	v1_controllers.NewMulticlusterServiceReconcileLoop("Service", clusters).AddMulticlusterServiceReconciler(ctx, r)
	v1_controllers.NewMulticlusterPodReconcileLoop("Pod", clusters).AddMulticlusterPodReconciler(ctx, r)
	v1_controllers.NewMulticlusterNodeReconcileLoop("Node", clusters).AddMulticlusterNodeReconciler(ctx, r)

	apps_v1_controllers.NewMulticlusterDeploymentReconcileLoop("Deployment", clusters).AddMulticlusterDeploymentReconciler(ctx, r)
	apps_v1_controllers.NewMulticlusterReplicaSetReconcileLoop("ReplicaSet", clusters).AddMulticlusterReplicaSetReconciler(ctx, r)
	apps_v1_controllers.NewMulticlusterDaemonSetReconcileLoop("DaemonSet", clusters).AddMulticlusterDaemonSetReconciler(ctx, r)
	apps_v1_controllers.NewMulticlusterStatefulSetReconcileLoop("StatefulSet", clusters).AddMulticlusterStatefulSetReconciler(ctx, r)

}

func (r *multiClusterReconcilerImpl) ReconcileConfigMap(clusterName string, obj *v1.ConfigMap) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileConfigMapDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileService(clusterName string, obj *v1.Service) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileServiceDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcilePod(clusterName string, obj *v1.Pod) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcilePodDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileNode(clusterName string, obj *v1.Node) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileNodeDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileDeployment(clusterName string, obj *apps_v1.Deployment) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileDeploymentDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileReplicaSet(clusterName string, obj *apps_v1.ReplicaSet) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileReplicaSetDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileDaemonSet(clusterName string, obj *apps_v1.DaemonSet) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileDaemonSetDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileStatefulSet(clusterName string, obj *apps_v1.StatefulSet) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileStatefulSetDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

// the singleClusterReconciler reconciles events for input resources across clusters
type singleClusterReconciler interface {
	v1_controllers.ConfigMapReconciler
	v1_controllers.ServiceReconciler
	v1_controllers.PodReconciler
	v1_controllers.NodeReconciler

	apps_v1_controllers.DeploymentReconciler
	apps_v1_controllers.ReplicaSetReconciler
	apps_v1_controllers.DaemonSetReconciler
	apps_v1_controllers.StatefulSetReconciler
}

var _ singleClusterReconciler = &singleClusterReconcilerImpl{}

type singleClusterReconcilerImpl struct {
	base input.SingleClusterReconciler
}

// register the reconcile func with the manager
func RegisterSingleClusterReconciler(
	ctx context.Context,
	mgr manager.Manager,
	reconcileFunc input.SingleClusterReconcileFunc,
) error {

	base := input.NewSingleClusterReconciler(
		ctx,
		reconcileFunc,
	)

	r := &singleClusterReconcilerImpl{
		base: base,
	}

	// initialize reconcile loops

	if err := v1_controllers.NewConfigMapReconcileLoop("ConfigMap", mgr, reconcile.Options{}).RunConfigMapReconciler(ctx, r); err != nil {
		return err
	}
	if err := v1_controllers.NewServiceReconcileLoop("Service", mgr, reconcile.Options{}).RunServiceReconciler(ctx, r); err != nil {
		return err
	}
	if err := v1_controllers.NewPodReconcileLoop("Pod", mgr, reconcile.Options{}).RunPodReconciler(ctx, r); err != nil {
		return err
	}
	if err := v1_controllers.NewNodeReconcileLoop("Node", mgr, reconcile.Options{}).RunNodeReconciler(ctx, r); err != nil {
		return err
	}

	if err := apps_v1_controllers.NewDeploymentReconcileLoop("Deployment", mgr, reconcile.Options{}).RunDeploymentReconciler(ctx, r); err != nil {
		return err
	}
	if err := apps_v1_controllers.NewReplicaSetReconcileLoop("ReplicaSet", mgr, reconcile.Options{}).RunReplicaSetReconciler(ctx, r); err != nil {
		return err
	}
	if err := apps_v1_controllers.NewDaemonSetReconcileLoop("DaemonSet", mgr, reconcile.Options{}).RunDaemonSetReconciler(ctx, r); err != nil {
		return err
	}
	if err := apps_v1_controllers.NewStatefulSetReconcileLoop("StatefulSet", mgr, reconcile.Options{}).RunStatefulSetReconciler(ctx, r); err != nil {
		return err
	}

	return nil
}

func (r *singleClusterReconcilerImpl) ReconcileConfigMap(obj *v1.ConfigMap) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileConfigMapDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileService(obj *v1.Service) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileServiceDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcilePod(obj *v1.Pod) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcilePodDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileNode(obj *v1.Node) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileNodeDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileDeployment(obj *apps_v1.Deployment) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileDeploymentDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileReplicaSet(obj *apps_v1.ReplicaSet) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileReplicaSetDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileDaemonSet(obj *apps_v1.DaemonSet) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileDaemonSetDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileStatefulSet(obj *apps_v1.StatefulSet) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileStatefulSetDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}
