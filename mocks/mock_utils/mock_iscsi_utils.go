// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/utils (interfaces: IscsiReconcileUtils)

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/netapp/trident/utils/models"
)

// MockIscsiReconcileUtils is a mock of IscsiReconcileUtils interface.
type MockIscsiReconcileUtils struct {
	ctrl     *gomock.Controller
	recorder *MockIscsiReconcileUtilsMockRecorder
}

// MockIscsiReconcileUtilsMockRecorder is the mock recorder for MockIscsiReconcileUtils.
type MockIscsiReconcileUtilsMockRecorder struct {
	mock *MockIscsiReconcileUtils
}

// NewMockIscsiReconcileUtils creates a new mock instance.
func NewMockIscsiReconcileUtils(ctrl *gomock.Controller) *MockIscsiReconcileUtils {
	mock := &MockIscsiReconcileUtils{ctrl: ctrl}
	mock.recorder = &MockIscsiReconcileUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIscsiReconcileUtils) EXPECT() *MockIscsiReconcileUtilsMockRecorder {
	return m.recorder
}

// GetDevicesForLUN mocks base method.
func (m *MockIscsiReconcileUtils) GetDevicesForLUN(arg0 []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesForLUN", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesForLUN indicates an expected call of GetDevicesForLUN.
func (mr *MockIscsiReconcileUtilsMockRecorder) GetDevicesForLUN(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesForLUN", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).GetDevicesForLUN), arg0)
}

// GetISCSIHostSessionMapForTarget mocks base method.
func (m *MockIscsiReconcileUtils) GetISCSIHostSessionMapForTarget(arg0 context.Context, arg1 string) map[int]int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetISCSIHostSessionMapForTarget", arg0, arg1)
	ret0, _ := ret[0].(map[int]int)
	return ret0
}

// GetISCSIHostSessionMapForTarget indicates an expected call of GetISCSIHostSessionMapForTarget.
func (mr *MockIscsiReconcileUtilsMockRecorder) GetISCSIHostSessionMapForTarget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetISCSIHostSessionMapForTarget", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).GetISCSIHostSessionMapForTarget), arg0, arg1)
}

// GetMultipathDeviceBySerial mocks base method.
func (m *MockIscsiReconcileUtils) GetMultipathDeviceBySerial(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMultipathDeviceBySerial", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultipathDeviceBySerial indicates an expected call of GetMultipathDeviceBySerial.
func (mr *MockIscsiReconcileUtilsMockRecorder) GetMultipathDeviceBySerial(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultipathDeviceBySerial", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).GetMultipathDeviceBySerial), arg0, arg1)
}

// GetMultipathDeviceDisks mocks base method.
func (m *MockIscsiReconcileUtils) GetMultipathDeviceDisks(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMultipathDeviceDisks", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultipathDeviceDisks indicates an expected call of GetMultipathDeviceDisks.
func (mr *MockIscsiReconcileUtilsMockRecorder) GetMultipathDeviceDisks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultipathDeviceDisks", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).GetMultipathDeviceDisks), arg0, arg1)
}

// GetMultipathDeviceUUID mocks base method.
func (m *MockIscsiReconcileUtils) GetMultipathDeviceUUID(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMultipathDeviceUUID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultipathDeviceUUID indicates an expected call of GetMultipathDeviceUUID.
func (mr *MockIscsiReconcileUtilsMockRecorder) GetMultipathDeviceUUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultipathDeviceUUID", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).GetMultipathDeviceUUID), arg0)
}

// GetSysfsBlockDirsForLUN mocks base method.
func (m *MockIscsiReconcileUtils) GetSysfsBlockDirsForLUN(arg0 int, arg1 map[int]int) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSysfsBlockDirsForLUN", arg0, arg1)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetSysfsBlockDirsForLUN indicates an expected call of GetSysfsBlockDirsForLUN.
func (mr *MockIscsiReconcileUtilsMockRecorder) GetSysfsBlockDirsForLUN(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSysfsBlockDirsForLUN", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).GetSysfsBlockDirsForLUN), arg0, arg1)
}

// ReconcileISCSIVolumeInfo mocks base method.
func (m *MockIscsiReconcileUtils) ReconcileISCSIVolumeInfo(arg0 context.Context, arg1 *models.VolumeTrackingInfo) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileISCSIVolumeInfo", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileISCSIVolumeInfo indicates an expected call of ReconcileISCSIVolumeInfo.
func (mr *MockIscsiReconcileUtilsMockRecorder) ReconcileISCSIVolumeInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileISCSIVolumeInfo", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).ReconcileISCSIVolumeInfo), arg0, arg1)
}
