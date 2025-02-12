// Code generated by skv2. DO NOT EDIT.

package input

import (
	"context"
	"time"

	"github.com/solo-io/skv2/contrib/pkg/input"
	sk_core_v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/reconcile"

	networking_istio_io_v1alpha3_controllers "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/controller"
	security_istio_io_v1beta1_controllers "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1/controller"
	v1_controllers "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/controller"
	certificates_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"
	certificates_mesh_gloo_solo_io_v1_controllers "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1/controller"
	discovery_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	discovery_mesh_gloo_solo_io_v1_controllers "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/controller"
	networking_enterprise_mesh_gloo_solo_io_v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1"
	networking_enterprise_mesh_gloo_solo_io_v1beta1_controllers "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1/controller"
	networking_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	networking_mesh_gloo_solo_io_v1_controllers "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1/controller"
	observability_enterprise_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/observability.enterprise.mesh.gloo.solo.io/v1"
	observability_enterprise_mesh_gloo_solo_io_v1_controllers "github.com/solo-io/gloo-mesh/pkg/api/observability.enterprise.mesh.gloo.solo.io/v1/controller"
	settings_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1"
	settings_mesh_gloo_solo_io_v1_controllers "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1/controller"
	xds_agent_enterprise_mesh_gloo_solo_io_v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1beta1"
	xds_agent_enterprise_mesh_gloo_solo_io_v1beta1_controllers "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1beta1/controller"
	multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	multicluster_solo_io_v1alpha1_controllers "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/controller"
	networking_istio_io_v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	security_istio_io_v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// The Input Reconciler calls a simple func(id) error whenever a
// storage event is received for any of:
// * IssuedCertificates
// * PodBounceDirectives
// * XdsConfigs
// * DestinationRules
// * EnvoyFilters
// * Gateways
// * ServiceEntries
// * VirtualServices
// * AuthorizationPolicies
// from a remote cluster.
// * Settings
// * Destinations
// * Workloads
// * Meshes
// * TrafficPolicies
// * AccessPolicies
// * VirtualMeshes
// * WasmDeployments
// * VirtualDestinations
// * AccessLogRecords
// * Secrets
// * KubernetesClusters
// from the local cluster.

type ReconcileOptions struct {
	Remote RemoteReconcileOptions
	Local  LocalReconcileOptions

	// the ReconcileInterval, if greater than 0, will limit the number of reconciles
	// to one per interval.
	ReconcileInterval time.Duration
}

// register the given multi cluster reconcile func with the cluster watcher
// register the given single cluster reconcile func with the local manager
func RegisterInputReconciler(
	ctx context.Context,
	clusters multicluster.ClusterWatcher,
	multiClusterReconcileFunc input.MultiClusterReconcileFunc,
	mgr manager.Manager,
	singleClusterReconcileFunc input.SingleClusterReconcileFunc,
	options ReconcileOptions,
) (input.InputReconciler, error) {
	// [certificates.mesh.gloo.solo.io/v1 xds.agent.enterprise.mesh.gloo.solo.io/v1beta1 networking.istio.io/v1alpha3 security.istio.io/v1beta1] false 4
	// [settings.mesh.gloo.solo.io/v1 discovery.mesh.gloo.solo.io/v1 networking.mesh.gloo.solo.io/v1 networking.enterprise.mesh.gloo.solo.io/v1beta1 observability.enterprise.mesh.gloo.solo.io/v1 v1 multicluster.solo.io/v1alpha1]

	base := input.NewInputReconciler(
		ctx,
		multiClusterReconcileFunc,
		singleClusterReconcileFunc,
		options.ReconcileInterval,
	)

	// initialize reconcile loops

	// initialize IssuedCertificates reconcile loop for remote clusters
	certificates_mesh_gloo_solo_io_v1_controllers.NewMulticlusterIssuedCertificateReconcileLoop("IssuedCertificate", clusters, options.Remote.IssuedCertificates).AddMulticlusterIssuedCertificateReconciler(ctx, &remoteInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize PodBounceDirectives reconcile loop for remote clusters
	certificates_mesh_gloo_solo_io_v1_controllers.NewMulticlusterPodBounceDirectiveReconcileLoop("PodBounceDirective", clusters, options.Remote.PodBounceDirectives).AddMulticlusterPodBounceDirectiveReconciler(ctx, &remoteInputReconciler{base: base}, options.Remote.Predicates...)

	// initialize XdsConfigs reconcile loop for remote clusters
	xds_agent_enterprise_mesh_gloo_solo_io_v1beta1_controllers.NewMulticlusterXdsConfigReconcileLoop("XdsConfig", clusters, options.Remote.XdsConfigs).AddMulticlusterXdsConfigReconciler(ctx, &remoteInputReconciler{base: base}, options.Remote.Predicates...)

	// initialize DestinationRules reconcile loop for remote clusters
	networking_istio_io_v1alpha3_controllers.NewMulticlusterDestinationRuleReconcileLoop("DestinationRule", clusters, options.Remote.DestinationRules).AddMulticlusterDestinationRuleReconciler(ctx, &remoteInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize EnvoyFilters reconcile loop for remote clusters
	networking_istio_io_v1alpha3_controllers.NewMulticlusterEnvoyFilterReconcileLoop("EnvoyFilter", clusters, options.Remote.EnvoyFilters).AddMulticlusterEnvoyFilterReconciler(ctx, &remoteInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize Gateways reconcile loop for remote clusters
	networking_istio_io_v1alpha3_controllers.NewMulticlusterGatewayReconcileLoop("Gateway", clusters, options.Remote.Gateways).AddMulticlusterGatewayReconciler(ctx, &remoteInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize ServiceEntries reconcile loop for remote clusters
	networking_istio_io_v1alpha3_controllers.NewMulticlusterServiceEntryReconcileLoop("ServiceEntry", clusters, options.Remote.ServiceEntries).AddMulticlusterServiceEntryReconciler(ctx, &remoteInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize VirtualServices reconcile loop for remote clusters
	networking_istio_io_v1alpha3_controllers.NewMulticlusterVirtualServiceReconcileLoop("VirtualService", clusters, options.Remote.VirtualServices).AddMulticlusterVirtualServiceReconciler(ctx, &remoteInputReconciler{base: base}, options.Remote.Predicates...)

	// initialize AuthorizationPolicies reconcile loop for remote clusters
	security_istio_io_v1beta1_controllers.NewMulticlusterAuthorizationPolicyReconcileLoop("AuthorizationPolicy", clusters, options.Remote.AuthorizationPolicies).AddMulticlusterAuthorizationPolicyReconciler(ctx, &remoteInputReconciler{base: base}, options.Remote.Predicates...)

	// initialize Settings reconcile loop for local cluster
	if err := settings_mesh_gloo_solo_io_v1_controllers.NewSettingsReconcileLoop("Settings", mgr, options.Local.Settings).RunSettingsReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	// initialize Destinations reconcile loop for local cluster
	if err := discovery_mesh_gloo_solo_io_v1_controllers.NewDestinationReconcileLoop("Destination", mgr, options.Local.Destinations).RunDestinationReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}
	// initialize Workloads reconcile loop for local cluster
	if err := discovery_mesh_gloo_solo_io_v1_controllers.NewWorkloadReconcileLoop("Workload", mgr, options.Local.Workloads).RunWorkloadReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}
	// initialize Meshes reconcile loop for local cluster
	if err := discovery_mesh_gloo_solo_io_v1_controllers.NewMeshReconcileLoop("Mesh", mgr, options.Local.Meshes).RunMeshReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	// initialize TrafficPolicies reconcile loop for local cluster
	if err := networking_mesh_gloo_solo_io_v1_controllers.NewTrafficPolicyReconcileLoop("TrafficPolicy", mgr, options.Local.TrafficPolicies).RunTrafficPolicyReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}
	// initialize AccessPolicies reconcile loop for local cluster
	if err := networking_mesh_gloo_solo_io_v1_controllers.NewAccessPolicyReconcileLoop("AccessPolicy", mgr, options.Local.AccessPolicies).RunAccessPolicyReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}
	// initialize VirtualMeshes reconcile loop for local cluster
	if err := networking_mesh_gloo_solo_io_v1_controllers.NewVirtualMeshReconcileLoop("VirtualMesh", mgr, options.Local.VirtualMeshes).RunVirtualMeshReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	// initialize WasmDeployments reconcile loop for local cluster
	if err := networking_enterprise_mesh_gloo_solo_io_v1beta1_controllers.NewWasmDeploymentReconcileLoop("WasmDeployment", mgr, options.Local.WasmDeployments).RunWasmDeploymentReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}
	// initialize VirtualDestinations reconcile loop for local cluster
	if err := networking_enterprise_mesh_gloo_solo_io_v1beta1_controllers.NewVirtualDestinationReconcileLoop("VirtualDestination", mgr, options.Local.VirtualDestinations).RunVirtualDestinationReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	// initialize AccessLogRecords reconcile loop for local cluster
	if err := observability_enterprise_mesh_gloo_solo_io_v1_controllers.NewAccessLogRecordReconcileLoop("AccessLogRecord", mgr, options.Local.AccessLogRecords).RunAccessLogRecordReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	// initialize Secrets reconcile loop for local cluster
	if err := v1_controllers.NewSecretReconcileLoop("Secret", mgr, options.Local.Secrets).RunSecretReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	// initialize KubernetesClusters reconcile loop for local cluster
	if err := multicluster_solo_io_v1alpha1_controllers.NewKubernetesClusterReconcileLoop("KubernetesCluster", mgr, options.Local.KubernetesClusters).RunKubernetesClusterReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	return base, nil
}

// Options for reconciling a snapshot in remote clusters
type RemoteReconcileOptions struct {

	// Options for reconciling IssuedCertificates
	IssuedCertificates reconcile.Options
	// Options for reconciling PodBounceDirectives
	PodBounceDirectives reconcile.Options

	// Options for reconciling XdsConfigs
	XdsConfigs reconcile.Options

	// Options for reconciling DestinationRules
	DestinationRules reconcile.Options
	// Options for reconciling EnvoyFilters
	EnvoyFilters reconcile.Options
	// Options for reconciling Gateways
	Gateways reconcile.Options
	// Options for reconciling ServiceEntries
	ServiceEntries reconcile.Options
	// Options for reconciling VirtualServices
	VirtualServices reconcile.Options

	// Options for reconciling AuthorizationPolicies
	AuthorizationPolicies reconcile.Options

	// optional predicates for filtering remote events
	Predicates []predicate.Predicate
}

type remoteInputReconciler struct {
	base input.InputReconciler
}

func (r *remoteInputReconciler) ReconcileIssuedCertificate(clusterName string, obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteInputReconciler) ReconcileIssuedCertificateDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteInputReconciler) ReconcilePodBounceDirective(clusterName string, obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteInputReconciler) ReconcilePodBounceDirectiveDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteInputReconciler) ReconcileXdsConfig(clusterName string, obj *xds_agent_enterprise_mesh_gloo_solo_io_v1beta1.XdsConfig) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteInputReconciler) ReconcileXdsConfigDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteInputReconciler) ReconcileDestinationRule(clusterName string, obj *networking_istio_io_v1alpha3.DestinationRule) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteInputReconciler) ReconcileDestinationRuleDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteInputReconciler) ReconcileEnvoyFilter(clusterName string, obj *networking_istio_io_v1alpha3.EnvoyFilter) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteInputReconciler) ReconcileEnvoyFilterDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteInputReconciler) ReconcileGateway(clusterName string, obj *networking_istio_io_v1alpha3.Gateway) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteInputReconciler) ReconcileGatewayDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteInputReconciler) ReconcileServiceEntry(clusterName string, obj *networking_istio_io_v1alpha3.ServiceEntry) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteInputReconciler) ReconcileServiceEntryDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteInputReconciler) ReconcileVirtualService(clusterName string, obj *networking_istio_io_v1alpha3.VirtualService) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteInputReconciler) ReconcileVirtualServiceDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteInputReconciler) ReconcileAuthorizationPolicy(clusterName string, obj *security_istio_io_v1beta1.AuthorizationPolicy) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteInputReconciler) ReconcileAuthorizationPolicyDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

// Options for reconciling a snapshot in remote clusters
type LocalReconcileOptions struct {

	// Options for reconciling Settings
	Settings reconcile.Options

	// Options for reconciling Destinations
	Destinations reconcile.Options
	// Options for reconciling Workloads
	Workloads reconcile.Options
	// Options for reconciling Meshes
	Meshes reconcile.Options

	// Options for reconciling TrafficPolicies
	TrafficPolicies reconcile.Options
	// Options for reconciling AccessPolicies
	AccessPolicies reconcile.Options
	// Options for reconciling VirtualMeshes
	VirtualMeshes reconcile.Options

	// Options for reconciling WasmDeployments
	WasmDeployments reconcile.Options
	// Options for reconciling VirtualDestinations
	VirtualDestinations reconcile.Options

	// Options for reconciling AccessLogRecords
	AccessLogRecords reconcile.Options

	// Options for reconciling Secrets
	Secrets reconcile.Options

	// Options for reconciling KubernetesClusters
	KubernetesClusters reconcile.Options

	// optional predicates for filtering local events
	Predicates []predicate.Predicate
}

type localInputReconciler struct {
	base input.InputReconciler
}

func (r *localInputReconciler) ReconcileSettings(obj *settings_mesh_gloo_solo_io_v1.Settings) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileSettingsDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localInputReconciler) ReconcileDestination(obj *discovery_mesh_gloo_solo_io_v1.Destination) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileDestinationDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localInputReconciler) ReconcileWorkload(obj *discovery_mesh_gloo_solo_io_v1.Workload) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileWorkloadDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localInputReconciler) ReconcileMesh(obj *discovery_mesh_gloo_solo_io_v1.Mesh) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileMeshDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localInputReconciler) ReconcileTrafficPolicy(obj *networking_mesh_gloo_solo_io_v1.TrafficPolicy) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileTrafficPolicyDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localInputReconciler) ReconcileAccessPolicy(obj *networking_mesh_gloo_solo_io_v1.AccessPolicy) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileAccessPolicyDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localInputReconciler) ReconcileVirtualMesh(obj *networking_mesh_gloo_solo_io_v1.VirtualMesh) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileVirtualMeshDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localInputReconciler) ReconcileWasmDeployment(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileWasmDeploymentDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localInputReconciler) ReconcileVirtualDestination(obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileVirtualDestinationDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localInputReconciler) ReconcileAccessLogRecord(obj *observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileAccessLogRecordDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localInputReconciler) ReconcileSecret(obj *v1.Secret) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileSecretDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localInputReconciler) ReconcileKubernetesCluster(obj *multicluster_solo_io_v1alpha1.KubernetesCluster) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileKubernetesClusterDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}
