// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/utils/iscsi (interfaces: ISCSI)

// Package mock_iscsi is a generated GoMock package.
package mock_iscsi

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/netapp/trident/utils/models"
)

// MockISCSI is a mock of ISCSI interface.
type MockISCSI struct {
	ctrl     *gomock.Controller
	recorder *MockISCSIMockRecorder
}

// MockISCSIMockRecorder is the mock recorder for MockISCSI.
type MockISCSIMockRecorder struct {
	mock *MockISCSI
}

// NewMockISCSI creates a new mock instance.
func NewMockISCSI(ctrl *gomock.Controller) *MockISCSI {
	mock := &MockISCSI{ctrl: ctrl}
	mock.recorder = &MockISCSIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISCSI) EXPECT() *MockISCSIMockRecorder {
	return m.recorder
}

// AddSession mocks base method.
func (m *MockISCSI) AddSession(arg0 context.Context, arg1 *models.ISCSISessions, arg2 *models.VolumePublishInfo, arg3, arg4 string, arg5 models.PortalInvalid) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSession", arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddSession indicates an expected call of AddSession.
func (mr *MockISCSIMockRecorder) AddSession(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSession", reflect.TypeOf((*MockISCSI)(nil).AddSession), arg0, arg1, arg2, arg3, arg4, arg5)
}

// AttachVolumeRetry mocks base method.
func (m *MockISCSI) AttachVolumeRetry(arg0 context.Context, arg1, arg2 string, arg3 *models.VolumePublishInfo, arg4 map[string]string, arg5 time.Duration) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachVolumeRetry", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachVolumeRetry indicates an expected call of AttachVolumeRetry.
func (mr *MockISCSIMockRecorder) AttachVolumeRetry(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachVolumeRetry", reflect.TypeOf((*MockISCSI)(nil).AttachVolumeRetry), arg0, arg1, arg2, arg3, arg4, arg5)
}

// IsAlreadyAttached mocks base method.
func (m *MockISCSI) IsAlreadyAttached(arg0 context.Context, arg1 int, arg2 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAlreadyAttached", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAlreadyAttached indicates an expected call of IsAlreadyAttached.
func (mr *MockISCSIMockRecorder) IsAlreadyAttached(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAlreadyAttached", reflect.TypeOf((*MockISCSI)(nil).IsAlreadyAttached), arg0, arg1, arg2)
}

// PreChecks mocks base method.
func (m *MockISCSI) PreChecks(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreChecks", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PreChecks indicates an expected call of PreChecks.
func (mr *MockISCSIMockRecorder) PreChecks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreChecks", reflect.TypeOf((*MockISCSI)(nil).PreChecks), arg0)
}

// RescanDevices mocks base method.
func (m *MockISCSI) RescanDevices(arg0 context.Context, arg1 string, arg2 int32, arg3 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RescanDevices", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RescanDevices indicates an expected call of RescanDevices.
func (mr *MockISCSIMockRecorder) RescanDevices(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RescanDevices", reflect.TypeOf((*MockISCSI)(nil).RescanDevices), arg0, arg1, arg2, arg3)
}
