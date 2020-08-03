// Code generated by skv2. DO NOT EDIT.

/*
	Utility for manually building input snapshots. Used primarily in tests.
*/
package input

import (
	certificates_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2"
	certificates_smh_solo_io_v1alpha2_sets "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2/sets"
)

type InputSnapshotManualBuilder struct {
	name string

	issuedCertificates certificates_smh_solo_io_v1alpha2_sets.IssuedCertificateSet
}

func NewInputSnapshotManualBuilder(name string) *InputSnapshotManualBuilder {
	return &InputSnapshotManualBuilder{
		name: name,

		issuedCertificates: certificates_smh_solo_io_v1alpha2_sets.NewIssuedCertificateSet(),
	}
}

func (i *InputSnapshotManualBuilder) Build() Snapshot {
	return NewSnapshot(
		i.name,

		i.issuedCertificates,
	)
}
func (i *InputSnapshotManualBuilder) AddIssuedCertificates(issuedCertificates []*certificates_smh_solo_io_v1alpha2.IssuedCertificate) *InputSnapshotManualBuilder {
	i.issuedCertificates.Insert(issuedCertificates...)
	return i
}
