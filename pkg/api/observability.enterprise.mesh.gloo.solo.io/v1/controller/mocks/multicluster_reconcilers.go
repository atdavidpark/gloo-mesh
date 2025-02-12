// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/observability.enterprise.mesh.gloo.solo.io/v1"
	controller "github.com/solo-io/gloo-mesh/pkg/api/observability.enterprise.mesh.gloo.solo.io/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterAccessLogRecordReconciler is a mock of MulticlusterAccessLogRecordReconciler interface
type MockMulticlusterAccessLogRecordReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAccessLogRecordReconcilerMockRecorder
}

// MockMulticlusterAccessLogRecordReconcilerMockRecorder is the mock recorder for MockMulticlusterAccessLogRecordReconciler
type MockMulticlusterAccessLogRecordReconcilerMockRecorder struct {
	mock *MockMulticlusterAccessLogRecordReconciler
}

// NewMockMulticlusterAccessLogRecordReconciler creates a new mock instance
func NewMockMulticlusterAccessLogRecordReconciler(ctrl *gomock.Controller) *MockMulticlusterAccessLogRecordReconciler {
	mock := &MockMulticlusterAccessLogRecordReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAccessLogRecordReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterAccessLogRecordReconciler) EXPECT() *MockMulticlusterAccessLogRecordReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessLogRecord mocks base method
func (m *MockMulticlusterAccessLogRecordReconciler) ReconcileAccessLogRecord(clusterName string, obj *v1.AccessLogRecord) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessLogRecord", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAccessLogRecord indicates an expected call of ReconcileAccessLogRecord
func (mr *MockMulticlusterAccessLogRecordReconcilerMockRecorder) ReconcileAccessLogRecord(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessLogRecord", reflect.TypeOf((*MockMulticlusterAccessLogRecordReconciler)(nil).ReconcileAccessLogRecord), clusterName, obj)
}

// MockMulticlusterAccessLogRecordDeletionReconciler is a mock of MulticlusterAccessLogRecordDeletionReconciler interface
type MockMulticlusterAccessLogRecordDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAccessLogRecordDeletionReconcilerMockRecorder
}

// MockMulticlusterAccessLogRecordDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterAccessLogRecordDeletionReconciler
type MockMulticlusterAccessLogRecordDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterAccessLogRecordDeletionReconciler
}

// NewMockMulticlusterAccessLogRecordDeletionReconciler creates a new mock instance
func NewMockMulticlusterAccessLogRecordDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterAccessLogRecordDeletionReconciler {
	mock := &MockMulticlusterAccessLogRecordDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAccessLogRecordDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterAccessLogRecordDeletionReconciler) EXPECT() *MockMulticlusterAccessLogRecordDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessLogRecordDeletion mocks base method
func (m *MockMulticlusterAccessLogRecordDeletionReconciler) ReconcileAccessLogRecordDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessLogRecordDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileAccessLogRecordDeletion indicates an expected call of ReconcileAccessLogRecordDeletion
func (mr *MockMulticlusterAccessLogRecordDeletionReconcilerMockRecorder) ReconcileAccessLogRecordDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessLogRecordDeletion", reflect.TypeOf((*MockMulticlusterAccessLogRecordDeletionReconciler)(nil).ReconcileAccessLogRecordDeletion), clusterName, req)
}

// MockMulticlusterAccessLogRecordReconcileLoop is a mock of MulticlusterAccessLogRecordReconcileLoop interface
type MockMulticlusterAccessLogRecordReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAccessLogRecordReconcileLoopMockRecorder
}

// MockMulticlusterAccessLogRecordReconcileLoopMockRecorder is the mock recorder for MockMulticlusterAccessLogRecordReconcileLoop
type MockMulticlusterAccessLogRecordReconcileLoopMockRecorder struct {
	mock *MockMulticlusterAccessLogRecordReconcileLoop
}

// NewMockMulticlusterAccessLogRecordReconcileLoop creates a new mock instance
func NewMockMulticlusterAccessLogRecordReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterAccessLogRecordReconcileLoop {
	mock := &MockMulticlusterAccessLogRecordReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAccessLogRecordReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterAccessLogRecordReconcileLoop) EXPECT() *MockMulticlusterAccessLogRecordReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterAccessLogRecordReconciler mocks base method
func (m *MockMulticlusterAccessLogRecordReconcileLoop) AddMulticlusterAccessLogRecordReconciler(ctx context.Context, rec controller.MulticlusterAccessLogRecordReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterAccessLogRecordReconciler", varargs...)
}

// AddMulticlusterAccessLogRecordReconciler indicates an expected call of AddMulticlusterAccessLogRecordReconciler
func (mr *MockMulticlusterAccessLogRecordReconcileLoopMockRecorder) AddMulticlusterAccessLogRecordReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterAccessLogRecordReconciler", reflect.TypeOf((*MockMulticlusterAccessLogRecordReconcileLoop)(nil).AddMulticlusterAccessLogRecordReconciler), varargs...)
}
