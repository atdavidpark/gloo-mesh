// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller



import (
	"context"

	certificates_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the IssuedCertificate Resource across clusters.
// implemented by the user
type MulticlusterIssuedCertificateReconciler interface {
	ReconcileIssuedCertificate(clusterName string, obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) (reconcile.Result, error)
}

// Reconcile deletion events for the IssuedCertificate Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterIssuedCertificateDeletionReconciler interface {
	ReconcileIssuedCertificateDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterIssuedCertificateReconcilerFuncs struct {
	OnReconcileIssuedCertificate         func(clusterName string, obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) (reconcile.Result, error)
	OnReconcileIssuedCertificateDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterIssuedCertificateReconcilerFuncs) ReconcileIssuedCertificate(clusterName string, obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) (reconcile.Result, error) {
	if f.OnReconcileIssuedCertificate == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileIssuedCertificate(clusterName, obj)
}

func (f *MulticlusterIssuedCertificateReconcilerFuncs) ReconcileIssuedCertificateDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileIssuedCertificateDeletion == nil {
		return nil
	}
	return f.OnReconcileIssuedCertificateDeletion(clusterName, req)
}

type MulticlusterIssuedCertificateReconcileLoop interface {
	// AddMulticlusterIssuedCertificateReconciler adds a MulticlusterIssuedCertificateReconciler to the MulticlusterIssuedCertificateReconcileLoop.
	AddMulticlusterIssuedCertificateReconciler(ctx context.Context, rec MulticlusterIssuedCertificateReconciler, predicates ...predicate.Predicate)
}

type multiclusterIssuedCertificateReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterIssuedCertificateReconcileLoop) AddMulticlusterIssuedCertificateReconciler(ctx context.Context, rec MulticlusterIssuedCertificateReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericIssuedCertificateMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterIssuedCertificateReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterIssuedCertificateReconcileLoop {
	return &multiclusterIssuedCertificateReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &certificates_mesh_gloo_solo_io_v1.IssuedCertificate{}, options)}
}

type genericIssuedCertificateMulticlusterReconciler struct {
	reconciler MulticlusterIssuedCertificateReconciler
}

func (g genericIssuedCertificateMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterIssuedCertificateDeletionReconciler); ok {
		return deletionReconciler.ReconcileIssuedCertificateDeletion(cluster, req)
	}
	return nil
}

func (g genericIssuedCertificateMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.IssuedCertificate)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: IssuedCertificate handler received event for %T", object)
	}
	return g.reconciler.ReconcileIssuedCertificate(cluster, obj)
}

// Reconcile Upsert events for the CertificateRequest Resource across clusters.
// implemented by the user
type MulticlusterCertificateRequestReconciler interface {
	ReconcileCertificateRequest(clusterName string, obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) (reconcile.Result, error)
}

// Reconcile deletion events for the CertificateRequest Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterCertificateRequestDeletionReconciler interface {
	ReconcileCertificateRequestDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterCertificateRequestReconcilerFuncs struct {
	OnReconcileCertificateRequest         func(clusterName string, obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) (reconcile.Result, error)
	OnReconcileCertificateRequestDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterCertificateRequestReconcilerFuncs) ReconcileCertificateRequest(clusterName string, obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) (reconcile.Result, error) {
	if f.OnReconcileCertificateRequest == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCertificateRequest(clusterName, obj)
}

func (f *MulticlusterCertificateRequestReconcilerFuncs) ReconcileCertificateRequestDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileCertificateRequestDeletion == nil {
		return nil
	}
	return f.OnReconcileCertificateRequestDeletion(clusterName, req)
}

type MulticlusterCertificateRequestReconcileLoop interface {
	// AddMulticlusterCertificateRequestReconciler adds a MulticlusterCertificateRequestReconciler to the MulticlusterCertificateRequestReconcileLoop.
	AddMulticlusterCertificateRequestReconciler(ctx context.Context, rec MulticlusterCertificateRequestReconciler, predicates ...predicate.Predicate)
}

type multiclusterCertificateRequestReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterCertificateRequestReconcileLoop) AddMulticlusterCertificateRequestReconciler(ctx context.Context, rec MulticlusterCertificateRequestReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericCertificateRequestMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterCertificateRequestReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterCertificateRequestReconcileLoop {
	return &multiclusterCertificateRequestReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &certificates_mesh_gloo_solo_io_v1.CertificateRequest{}, options)}
}

type genericCertificateRequestMulticlusterReconciler struct {
	reconciler MulticlusterCertificateRequestReconciler
}

func (g genericCertificateRequestMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterCertificateRequestDeletionReconciler); ok {
		return deletionReconciler.ReconcileCertificateRequestDeletion(cluster, req)
	}
	return nil
}

func (g genericCertificateRequestMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.CertificateRequest)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CertificateRequest handler received event for %T", object)
	}
	return g.reconciler.ReconcileCertificateRequest(cluster, obj)
}

// Reconcile Upsert events for the PodBounceDirective Resource across clusters.
// implemented by the user
type MulticlusterPodBounceDirectiveReconciler interface {
	ReconcilePodBounceDirective(clusterName string, obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) (reconcile.Result, error)
}

// Reconcile deletion events for the PodBounceDirective Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterPodBounceDirectiveDeletionReconciler interface {
	ReconcilePodBounceDirectiveDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterPodBounceDirectiveReconcilerFuncs struct {
	OnReconcilePodBounceDirective         func(clusterName string, obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) (reconcile.Result, error)
	OnReconcilePodBounceDirectiveDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterPodBounceDirectiveReconcilerFuncs) ReconcilePodBounceDirective(clusterName string, obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) (reconcile.Result, error) {
	if f.OnReconcilePodBounceDirective == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcilePodBounceDirective(clusterName, obj)
}

func (f *MulticlusterPodBounceDirectiveReconcilerFuncs) ReconcilePodBounceDirectiveDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcilePodBounceDirectiveDeletion == nil {
		return nil
	}
	return f.OnReconcilePodBounceDirectiveDeletion(clusterName, req)
}

type MulticlusterPodBounceDirectiveReconcileLoop interface {
	// AddMulticlusterPodBounceDirectiveReconciler adds a MulticlusterPodBounceDirectiveReconciler to the MulticlusterPodBounceDirectiveReconcileLoop.
	AddMulticlusterPodBounceDirectiveReconciler(ctx context.Context, rec MulticlusterPodBounceDirectiveReconciler, predicates ...predicate.Predicate)
}

type multiclusterPodBounceDirectiveReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterPodBounceDirectiveReconcileLoop) AddMulticlusterPodBounceDirectiveReconciler(ctx context.Context, rec MulticlusterPodBounceDirectiveReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericPodBounceDirectiveMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterPodBounceDirectiveReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterPodBounceDirectiveReconcileLoop {
	return &multiclusterPodBounceDirectiveReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &certificates_mesh_gloo_solo_io_v1.PodBounceDirective{}, options)}
}

type genericPodBounceDirectiveMulticlusterReconciler struct {
	reconciler MulticlusterPodBounceDirectiveReconciler
}

func (g genericPodBounceDirectiveMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterPodBounceDirectiveDeletionReconciler); ok {
		return deletionReconciler.ReconcilePodBounceDirectiveDeletion(cluster, req)
	}
	return nil
}

func (g genericPodBounceDirectiveMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.PodBounceDirective)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: PodBounceDirective handler received event for %T", object)
	}
	return g.reconciler.ReconcilePodBounceDirective(cluster, obj)
}
