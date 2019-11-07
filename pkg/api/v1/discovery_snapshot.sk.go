// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type DiscoverySnapshot struct {
	Pods        github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList
	Upstreams   gloo_solo_io.UpstreamList
	Deployments github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.DeploymentList
	Tlssecrets  TlsSecretList
}

func (s DiscoverySnapshot) Clone() DiscoverySnapshot {
	return DiscoverySnapshot{
		Pods:        s.Pods.Clone(),
		Upstreams:   s.Upstreams.Clone(),
		Deployments: s.Deployments.Clone(),
		Tlssecrets:  s.Tlssecrets.Clone(),
	}
}

func (s DiscoverySnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashPods(),
		s.hashUpstreams(),
		s.hashDeployments(),
		s.hashTlssecrets(),
	)
}

func (s DiscoverySnapshot) hashPods() uint64 {
	return hashutils.HashAll(s.Pods.AsInterfaces()...)
}

func (s DiscoverySnapshot) hashUpstreams() uint64 {
	return hashutils.HashAll(s.Upstreams.AsInterfaces()...)
}

func (s DiscoverySnapshot) hashDeployments() uint64 {
	return hashutils.HashAll(s.Deployments.AsInterfaces()...)
}

func (s DiscoverySnapshot) hashTlssecrets() uint64 {
	return hashutils.HashAll(s.Tlssecrets.AsInterfaces()...)
}

func (s DiscoverySnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("pods", s.hashPods()))
	fields = append(fields, zap.Uint64("upstreams", s.hashUpstreams()))
	fields = append(fields, zap.Uint64("deployments", s.hashDeployments()))
	fields = append(fields, zap.Uint64("tlssecrets", s.hashTlssecrets()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type DiscoverySnapshotStringer struct {
	Version     uint64
	Pods        []string
	Upstreams   []string
	Deployments []string
	Tlssecrets  []string
}

func (ss DiscoverySnapshotStringer) String() string {
	s := fmt.Sprintf("DiscoverySnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Pods %v\n", len(ss.Pods))
	for _, name := range ss.Pods {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Deployments %v\n", len(ss.Deployments))
	for _, name := range ss.Deployments {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Tlssecrets %v\n", len(ss.Tlssecrets))
	for _, name := range ss.Tlssecrets {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s DiscoverySnapshot) Stringer() DiscoverySnapshotStringer {
	return DiscoverySnapshotStringer{
		Version:     s.Hash(),
		Pods:        s.Pods.NamespacesDotNames(),
		Upstreams:   s.Upstreams.NamespacesDotNames(),
		Deployments: s.Deployments.NamespacesDotNames(),
		Tlssecrets:  s.Tlssecrets.NamespacesDotNames(),
	}
}
