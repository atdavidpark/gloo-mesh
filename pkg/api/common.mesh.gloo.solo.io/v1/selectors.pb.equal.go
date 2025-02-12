// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/common/v1/selectors.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *DestinationSelector) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DestinationSelector)
	if !ok {
		that2, ok := that.(DestinationSelector)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetKubeServiceMatcher()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKubeServiceMatcher()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKubeServiceMatcher(), target.GetKubeServiceMatcher()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetKubeServiceRefs()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKubeServiceRefs()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKubeServiceRefs(), target.GetKubeServiceRefs()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *WorkloadSelector) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkloadSelector)
	if !ok {
		that2, ok := that.(WorkloadSelector)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetKubeWorkloadMatcher()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKubeWorkloadMatcher()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKubeWorkloadMatcher(), target.GetKubeWorkloadMatcher()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *IdentitySelector) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IdentitySelector)
	if !ok {
		that2, ok := that.(IdentitySelector)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetKubeIdentityMatcher()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKubeIdentityMatcher()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKubeIdentityMatcher(), target.GetKubeIdentityMatcher()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetKubeServiceAccountRefs()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKubeServiceAccountRefs()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKubeServiceAccountRefs(), target.GetKubeServiceAccountRefs()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *DestinationSelector_KubeServiceMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DestinationSelector_KubeServiceMatcher)
	if !ok {
		that2, ok := that.(DestinationSelector_KubeServiceMatcher)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetLabels()) != len(target.GetLabels()) {
		return false
	}
	for k, v := range m.GetLabels() {

		if strings.Compare(v, target.GetLabels()[k]) != 0 {
			return false
		}

	}

	if len(m.GetNamespaces()) != len(target.GetNamespaces()) {
		return false
	}
	for idx, v := range m.GetNamespaces() {

		if strings.Compare(v, target.GetNamespaces()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetClusters()) != len(target.GetClusters()) {
		return false
	}
	for idx, v := range m.GetClusters() {

		if strings.Compare(v, target.GetClusters()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *DestinationSelector_KubeServiceRefs) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DestinationSelector_KubeServiceRefs)
	if !ok {
		that2, ok := that.(DestinationSelector_KubeServiceRefs)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetServices()) != len(target.GetServices()) {
		return false
	}
	for idx, v := range m.GetServices() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetServices()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetServices()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WorkloadSelector_KubeWorkloadMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WorkloadSelector_KubeWorkloadMatcher)
	if !ok {
		that2, ok := that.(WorkloadSelector_KubeWorkloadMatcher)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetLabels()) != len(target.GetLabels()) {
		return false
	}
	for k, v := range m.GetLabels() {

		if strings.Compare(v, target.GetLabels()[k]) != 0 {
			return false
		}

	}

	if len(m.GetNamespaces()) != len(target.GetNamespaces()) {
		return false
	}
	for idx, v := range m.GetNamespaces() {

		if strings.Compare(v, target.GetNamespaces()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetClusters()) != len(target.GetClusters()) {
		return false
	}
	for idx, v := range m.GetClusters() {

		if strings.Compare(v, target.GetClusters()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *IdentitySelector_KubeIdentityMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IdentitySelector_KubeIdentityMatcher)
	if !ok {
		that2, ok := that.(IdentitySelector_KubeIdentityMatcher)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetNamespaces()) != len(target.GetNamespaces()) {
		return false
	}
	for idx, v := range m.GetNamespaces() {

		if strings.Compare(v, target.GetNamespaces()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetClusters()) != len(target.GetClusters()) {
		return false
	}
	for idx, v := range m.GetClusters() {

		if strings.Compare(v, target.GetClusters()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *IdentitySelector_KubeServiceAccountRefs) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IdentitySelector_KubeServiceAccountRefs)
	if !ok {
		that2, ok := that.(IdentitySelector_KubeServiceAccountRefs)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetServiceAccounts()) != len(target.GetServiceAccounts()) {
		return false
	}
	for idx, v := range m.GetServiceAccounts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetServiceAccounts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetServiceAccounts()[idx]) {
				return false
			}
		}

	}

	return true
}
