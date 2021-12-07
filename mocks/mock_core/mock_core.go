// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/core (interfaces: Orchestrator)

// Package mock_core is a generated GoMock package.
package mock_core

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	config "github.com/netapp/trident/config"
	core "github.com/netapp/trident/core"
	frontend "github.com/netapp/trident/frontend"
	storage "github.com/netapp/trident/storage"
	storageclass "github.com/netapp/trident/storage_class"
	utils "github.com/netapp/trident/utils"
)

// MockOrchestrator is a mock of Orchestrator interface.
type MockOrchestrator struct {
	ctrl     *gomock.Controller
	recorder *MockOrchestratorMockRecorder
}

// MockOrchestratorMockRecorder is the mock recorder for MockOrchestrator.
type MockOrchestratorMockRecorder struct {
	mock *MockOrchestrator
}

// NewMockOrchestrator creates a new mock instance.
func NewMockOrchestrator(ctrl *gomock.Controller) *MockOrchestrator {
	mock := &MockOrchestrator{ctrl: ctrl}
	mock.recorder = &MockOrchestratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrchestrator) EXPECT() *MockOrchestratorMockRecorder {
	return m.recorder
}

// AddBackend mocks base method.
func (m *MockOrchestrator) AddBackend(arg0 context.Context, arg1, arg2 string) (*storage.BackendExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBackend", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.BackendExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBackend indicates an expected call of AddBackend.
func (mr *MockOrchestratorMockRecorder) AddBackend(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBackend", reflect.TypeOf((*MockOrchestrator)(nil).AddBackend), arg0, arg1, arg2)
}

// AddFrontend mocks base method.
func (m *MockOrchestrator) AddFrontend(arg0 frontend.Plugin) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddFrontend", arg0)
}

// AddFrontend indicates an expected call of AddFrontend.
func (mr *MockOrchestratorMockRecorder) AddFrontend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFrontend", reflect.TypeOf((*MockOrchestrator)(nil).AddFrontend), arg0)
}

// AddNode mocks base method.
func (m *MockOrchestrator) AddNode(arg0 context.Context, arg1 *utils.Node, arg2 core.NodeEventCallback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNode", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNode indicates an expected call of AddNode.
func (mr *MockOrchestratorMockRecorder) AddNode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNode", reflect.TypeOf((*MockOrchestrator)(nil).AddNode), arg0, arg1, arg2)
}

// AddStorageClass mocks base method.
func (m *MockOrchestrator) AddStorageClass(arg0 context.Context, arg1 *storageclass.Config) (*storageclass.External, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStorageClass", arg0, arg1)
	ret0, _ := ret[0].(*storageclass.External)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddStorageClass indicates an expected call of AddStorageClass.
func (mr *MockOrchestratorMockRecorder) AddStorageClass(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStorageClass", reflect.TypeOf((*MockOrchestrator)(nil).AddStorageClass), arg0, arg1)
}

// AddVolume mocks base method.
func (m *MockOrchestrator) AddVolume(arg0 context.Context, arg1 *storage.VolumeConfig) (*storage.VolumeExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVolume", arg0, arg1)
	ret0, _ := ret[0].(*storage.VolumeExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddVolume indicates an expected call of AddVolume.
func (mr *MockOrchestratorMockRecorder) AddVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVolume", reflect.TypeOf((*MockOrchestrator)(nil).AddVolume), arg0, arg1)
}

// AddVolumeTransaction mocks base method.
func (m *MockOrchestrator) AddVolumeTransaction(arg0 context.Context, arg1 *storage.VolumeTransaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVolumeTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVolumeTransaction indicates an expected call of AddVolumeTransaction.
func (mr *MockOrchestratorMockRecorder) AddVolumeTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVolumeTransaction", reflect.TypeOf((*MockOrchestrator)(nil).AddVolumeTransaction), arg0, arg1)
}

// AttachVolume mocks base method.
func (m *MockOrchestrator) AttachVolume(arg0 context.Context, arg1, arg2 string, arg3 *utils.VolumePublishInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachVolume", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachVolume indicates an expected call of AttachVolume.
func (mr *MockOrchestratorMockRecorder) AttachVolume(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachVolume", reflect.TypeOf((*MockOrchestrator)(nil).AttachVolume), arg0, arg1, arg2, arg3)
}

// Bootstrap mocks base method.
func (m *MockOrchestrator) Bootstrap() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrap")
	ret0, _ := ret[0].(error)
	return ret0
}

// Bootstrap indicates an expected call of Bootstrap.
func (mr *MockOrchestratorMockRecorder) Bootstrap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockOrchestrator)(nil).Bootstrap))
}

// CanBackendMirror mocks base method.
func (m *MockOrchestrator) CanBackendMirror(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanBackendMirror", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanBackendMirror indicates an expected call of CanBackendMirror.
func (mr *MockOrchestratorMockRecorder) CanBackendMirror(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanBackendMirror", reflect.TypeOf((*MockOrchestrator)(nil).CanBackendMirror), arg0, arg1)
}

// CloneVolume mocks base method.
func (m *MockOrchestrator) CloneVolume(arg0 context.Context, arg1 *storage.VolumeConfig) (*storage.VolumeExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloneVolume", arg0, arg1)
	ret0, _ := ret[0].(*storage.VolumeExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloneVolume indicates an expected call of CloneVolume.
func (mr *MockOrchestratorMockRecorder) CloneVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneVolume", reflect.TypeOf((*MockOrchestrator)(nil).CloneVolume), arg0, arg1)
}

// CreateSnapshot mocks base method.
func (m *MockOrchestrator) CreateSnapshot(arg0 context.Context, arg1 *storage.SnapshotConfig) (*storage.SnapshotExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", arg0, arg1)
	ret0, _ := ret[0].(*storage.SnapshotExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockOrchestratorMockRecorder) CreateSnapshot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockOrchestrator)(nil).CreateSnapshot), arg0, arg1)
}

// DeleteBackend mocks base method.
func (m *MockOrchestrator) DeleteBackend(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBackend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBackend indicates an expected call of DeleteBackend.
func (mr *MockOrchestratorMockRecorder) DeleteBackend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBackend", reflect.TypeOf((*MockOrchestrator)(nil).DeleteBackend), arg0, arg1)
}

// DeleteBackendByBackendUUID mocks base method.
func (m *MockOrchestrator) DeleteBackendByBackendUUID(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBackendByBackendUUID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBackendByBackendUUID indicates an expected call of DeleteBackendByBackendUUID.
func (mr *MockOrchestratorMockRecorder) DeleteBackendByBackendUUID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBackendByBackendUUID", reflect.TypeOf((*MockOrchestrator)(nil).DeleteBackendByBackendUUID), arg0, arg1, arg2)
}

// DeleteNode mocks base method.
func (m *MockOrchestrator) DeleteNode(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNode indicates an expected call of DeleteNode.
func (mr *MockOrchestratorMockRecorder) DeleteNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNode", reflect.TypeOf((*MockOrchestrator)(nil).DeleteNode), arg0, arg1)
}

// DeleteSnapshot mocks base method.
func (m *MockOrchestrator) DeleteSnapshot(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockOrchestratorMockRecorder) DeleteSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockOrchestrator)(nil).DeleteSnapshot), arg0, arg1, arg2)
}

// DeleteStorageClass mocks base method.
func (m *MockOrchestrator) DeleteStorageClass(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStorageClass", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStorageClass indicates an expected call of DeleteStorageClass.
func (mr *MockOrchestratorMockRecorder) DeleteStorageClass(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStorageClass", reflect.TypeOf((*MockOrchestrator)(nil).DeleteStorageClass), arg0, arg1)
}

// DeleteVolume mocks base method.
func (m *MockOrchestrator) DeleteVolume(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolume indicates an expected call of DeleteVolume.
func (mr *MockOrchestratorMockRecorder) DeleteVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockOrchestrator)(nil).DeleteVolume), arg0, arg1)
}

// DeleteVolumeTransaction mocks base method.
func (m *MockOrchestrator) DeleteVolumeTransaction(arg0 context.Context, arg1 *storage.VolumeTransaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolumeTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolumeTransaction indicates an expected call of DeleteVolumeTransaction.
func (mr *MockOrchestratorMockRecorder) DeleteVolumeTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolumeTransaction", reflect.TypeOf((*MockOrchestrator)(nil).DeleteVolumeTransaction), arg0, arg1)
}

// DetachVolume mocks base method.
func (m *MockOrchestrator) DetachVolume(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachVolume indicates an expected call of DetachVolume.
func (mr *MockOrchestratorMockRecorder) DetachVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachVolume", reflect.TypeOf((*MockOrchestrator)(nil).DetachVolume), arg0, arg1, arg2)
}

// EstablishMirror mocks base method.
func (m *MockOrchestrator) EstablishMirror(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstablishMirror", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EstablishMirror indicates an expected call of EstablishMirror.
func (mr *MockOrchestratorMockRecorder) EstablishMirror(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstablishMirror", reflect.TypeOf((*MockOrchestrator)(nil).EstablishMirror), arg0, arg1, arg2, arg3)
}

// GetBackend mocks base method.
func (m *MockOrchestrator) GetBackend(arg0 context.Context, arg1 string) (*storage.BackendExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackend", arg0, arg1)
	ret0, _ := ret[0].(*storage.BackendExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackend indicates an expected call of GetBackend.
func (mr *MockOrchestratorMockRecorder) GetBackend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackend", reflect.TypeOf((*MockOrchestrator)(nil).GetBackend), arg0, arg1)
}

// GetBackendByBackendUUID mocks base method.
func (m *MockOrchestrator) GetBackendByBackendUUID(arg0 context.Context, arg1 string) (*storage.BackendExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackendByBackendUUID", arg0, arg1)
	ret0, _ := ret[0].(*storage.BackendExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackendByBackendUUID indicates an expected call of GetBackendByBackendUUID.
func (mr *MockOrchestratorMockRecorder) GetBackendByBackendUUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendByBackendUUID", reflect.TypeOf((*MockOrchestrator)(nil).GetBackendByBackendUUID), arg0, arg1)
}

// GetDriverTypeForVolume mocks base method.
func (m *MockOrchestrator) GetDriverTypeForVolume(arg0 context.Context, arg1 *storage.VolumeExternal) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverTypeForVolume", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDriverTypeForVolume indicates an expected call of GetDriverTypeForVolume.
func (mr *MockOrchestratorMockRecorder) GetDriverTypeForVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverTypeForVolume", reflect.TypeOf((*MockOrchestrator)(nil).GetDriverTypeForVolume), arg0, arg1)
}

// GetFrontend mocks base method.
func (m *MockOrchestrator) GetFrontend(arg0 context.Context, arg1 string) (frontend.Plugin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFrontend", arg0, arg1)
	ret0, _ := ret[0].(frontend.Plugin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFrontend indicates an expected call of GetFrontend.
func (mr *MockOrchestratorMockRecorder) GetFrontend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFrontend", reflect.TypeOf((*MockOrchestrator)(nil).GetFrontend), arg0, arg1)
}

// GetMirrorStatus mocks base method.
func (m *MockOrchestrator) GetMirrorStatus(arg0 context.Context, arg1, arg2, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMirrorStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMirrorStatus indicates an expected call of GetMirrorStatus.
func (mr *MockOrchestratorMockRecorder) GetMirrorStatus(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMirrorStatus", reflect.TypeOf((*MockOrchestrator)(nil).GetMirrorStatus), arg0, arg1, arg2, arg3)
}

// GetNode mocks base method.
func (m *MockOrchestrator) GetNode(arg0 context.Context, arg1 string) (*utils.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNode", arg0, arg1)
	ret0, _ := ret[0].(*utils.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNode indicates an expected call of GetNode.
func (mr *MockOrchestratorMockRecorder) GetNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNode", reflect.TypeOf((*MockOrchestrator)(nil).GetNode), arg0, arg1)
}

// GetSnapshot mocks base method.
func (m *MockOrchestrator) GetSnapshot(arg0 context.Context, arg1, arg2 string) (*storage.SnapshotExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.SnapshotExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshot indicates an expected call of GetSnapshot.
func (mr *MockOrchestratorMockRecorder) GetSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshot", reflect.TypeOf((*MockOrchestrator)(nil).GetSnapshot), arg0, arg1, arg2)
}

// GetStorageClass mocks base method.
func (m *MockOrchestrator) GetStorageClass(arg0 context.Context, arg1 string) (*storageclass.External, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageClass", arg0, arg1)
	ret0, _ := ret[0].(*storageclass.External)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageClass indicates an expected call of GetStorageClass.
func (mr *MockOrchestratorMockRecorder) GetStorageClass(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageClass", reflect.TypeOf((*MockOrchestrator)(nil).GetStorageClass), arg0, arg1)
}

// GetVersion mocks base method.
func (m *MockOrchestrator) GetVersion(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockOrchestratorMockRecorder) GetVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockOrchestrator)(nil).GetVersion), arg0)
}

// GetVolume mocks base method.
func (m *MockOrchestrator) GetVolume(arg0 context.Context, arg1 string) (*storage.VolumeExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolume", arg0, arg1)
	ret0, _ := ret[0].(*storage.VolumeExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolume indicates an expected call of GetVolume.
func (mr *MockOrchestratorMockRecorder) GetVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolume", reflect.TypeOf((*MockOrchestrator)(nil).GetVolume), arg0, arg1)
}

// GetVolumeExternal mocks base method.
func (m *MockOrchestrator) GetVolumeExternal(arg0 context.Context, arg1, arg2 string) (*storage.VolumeExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeExternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.VolumeExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeExternal indicates an expected call of GetVolumeExternal.
func (mr *MockOrchestratorMockRecorder) GetVolumeExternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeExternal", reflect.TypeOf((*MockOrchestrator)(nil).GetVolumeExternal), arg0, arg1, arg2)
}

// GetVolumeByInternalName mocks base method.
func (m *MockOrchestrator) GetVolumeByInternalName(volumeInternal string, ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeByInternalName", volumeInternal, ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeByInternalName indicates an expected call of GetVolumeByInternalName.
func (mr *MockOrchestratorMockRecorder) GetVolumeByInternalName(volumeInternal, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeByInternalName", reflect.TypeOf((*MockOrchestrator)(nil).GetVolumeByInternalName), volumeInternal, ctx)
}

// GetVolumeTransaction mocks base method.
func (m *MockOrchestrator) GetVolumeTransaction(arg0 context.Context, arg1 *storage.VolumeTransaction) (*storage.VolumeTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeTransaction", arg0, arg1)
	ret0, _ := ret[0].(*storage.VolumeTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeTransaction indicates an expected call of GetVolumeTransaction.
func (mr *MockOrchestratorMockRecorder) GetVolumeTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeTransaction", reflect.TypeOf((*MockOrchestrator)(nil).GetVolumeTransaction), arg0, arg1)
}

// GetVolumeType mocks base method.
func (m *MockOrchestrator) GetVolumeType(arg0 context.Context, arg1 *storage.VolumeExternal) (config.VolumeType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeType", arg0, arg1)
	ret0, _ := ret[0].(config.VolumeType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeType indicates an expected call of GetVolumeType.
func (mr *MockOrchestratorMockRecorder) GetVolumeType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeType", reflect.TypeOf((*MockOrchestrator)(nil).GetVolumeType), arg0, arg1)
}

// ImportVolume mocks base method.
func (m *MockOrchestrator) ImportVolume(arg0 context.Context, arg1 *storage.VolumeConfig) (*storage.VolumeExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportVolume", arg0, arg1)
	ret0, _ := ret[0].(*storage.VolumeExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportVolume indicates an expected call of ImportVolume.
func (mr *MockOrchestratorMockRecorder) ImportVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportVolume", reflect.TypeOf((*MockOrchestrator)(nil).ImportVolume), arg0, arg1)
}

// LegacyImportVolume mocks base method.
func (m *MockOrchestrator) LegacyImportVolume(arg0 context.Context, arg1 *storage.VolumeConfig, arg2 string, arg3 bool, arg4 core.VolumeCallback) (*storage.VolumeExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LegacyImportVolume", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*storage.VolumeExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LegacyImportVolume indicates an expected call of LegacyImportVolume.
func (mr *MockOrchestratorMockRecorder) LegacyImportVolume(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LegacyImportVolume", reflect.TypeOf((*MockOrchestrator)(nil).LegacyImportVolume), arg0, arg1, arg2, arg3, arg4)
}

// ListBackends mocks base method.
func (m *MockOrchestrator) ListBackends(arg0 context.Context) ([]*storage.BackendExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBackends", arg0)
	ret0, _ := ret[0].([]*storage.BackendExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBackends indicates an expected call of ListBackends.
func (mr *MockOrchestratorMockRecorder) ListBackends(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackends", reflect.TypeOf((*MockOrchestrator)(nil).ListBackends), arg0)
}

// ListNodes mocks base method.
func (m *MockOrchestrator) ListNodes(arg0 context.Context) ([]*utils.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNodes", arg0)
	ret0, _ := ret[0].([]*utils.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNodes indicates an expected call of ListNodes.
func (mr *MockOrchestratorMockRecorder) ListNodes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodes", reflect.TypeOf((*MockOrchestrator)(nil).ListNodes), arg0)
}

// ListSnapshots mocks base method.
func (m *MockOrchestrator) ListSnapshots(arg0 context.Context) ([]*storage.SnapshotExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSnapshots", arg0)
	ret0, _ := ret[0].([]*storage.SnapshotExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSnapshots indicates an expected call of ListSnapshots.
func (mr *MockOrchestratorMockRecorder) ListSnapshots(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnapshots", reflect.TypeOf((*MockOrchestrator)(nil).ListSnapshots), arg0)
}

// ListSnapshotsByName mocks base method.
func (m *MockOrchestrator) ListSnapshotsByName(arg0 context.Context, arg1 string) ([]*storage.SnapshotExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSnapshotsByName", arg0, arg1)
	ret0, _ := ret[0].([]*storage.SnapshotExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSnapshotsByName indicates an expected call of ListSnapshotsByName.
func (mr *MockOrchestratorMockRecorder) ListSnapshotsByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnapshotsByName", reflect.TypeOf((*MockOrchestrator)(nil).ListSnapshotsByName), arg0, arg1)
}

// ListSnapshotsForVolume mocks base method.
func (m *MockOrchestrator) ListSnapshotsForVolume(arg0 context.Context, arg1 string) ([]*storage.SnapshotExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSnapshotsForVolume", arg0, arg1)
	ret0, _ := ret[0].([]*storage.SnapshotExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSnapshotsForVolume indicates an expected call of ListSnapshotsForVolume.
func (mr *MockOrchestratorMockRecorder) ListSnapshotsForVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnapshotsForVolume", reflect.TypeOf((*MockOrchestrator)(nil).ListSnapshotsForVolume), arg0, arg1)
}

// ListStorageClasses mocks base method.
func (m *MockOrchestrator) ListStorageClasses(arg0 context.Context) ([]*storageclass.External, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStorageClasses", arg0)
	ret0, _ := ret[0].([]*storageclass.External)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStorageClasses indicates an expected call of ListStorageClasses.
func (mr *MockOrchestratorMockRecorder) ListStorageClasses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStorageClasses", reflect.TypeOf((*MockOrchestrator)(nil).ListStorageClasses), arg0)
}

// ListVolumes mocks base method.
func (m *MockOrchestrator) ListVolumes(arg0 context.Context) ([]*storage.VolumeExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumes", arg0)
	ret0, _ := ret[0].([]*storage.VolumeExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVolumes indicates an expected call of ListVolumes.
func (mr *MockOrchestratorMockRecorder) ListVolumes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumes", reflect.TypeOf((*MockOrchestrator)(nil).ListVolumes), arg0)
}

// ListVolumesByPlugin mocks base method.
func (m *MockOrchestrator) ListVolumesByPlugin(arg0 context.Context, arg1 string) ([]*storage.VolumeExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumesByPlugin", arg0, arg1)
	ret0, _ := ret[0].([]*storage.VolumeExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVolumesByPlugin indicates an expected call of ListVolumesByPlugin.
func (mr *MockOrchestratorMockRecorder) ListVolumesByPlugin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumesByPlugin", reflect.TypeOf((*MockOrchestrator)(nil).ListVolumesByPlugin), arg0, arg1)
}

// PeriodicallyReconcileNodeAccessOnBackends mocks base method.
func (m *MockOrchestrator) PeriodicallyReconcileNodeAccessOnBackends() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PeriodicallyReconcileNodeAccessOnBackends")
}

// PeriodicallyReconcileNodeAccessOnBackends indicates an expected call of PeriodicallyReconcileNodeAccessOnBackends.
func (mr *MockOrchestratorMockRecorder) PeriodicallyReconcileNodeAccessOnBackends() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeriodicallyReconcileNodeAccessOnBackends", reflect.TypeOf((*MockOrchestrator)(nil).PeriodicallyReconcileNodeAccessOnBackends))
}

// PromoteMirror mocks base method.
func (m *MockOrchestrator) PromoteMirror(arg0 context.Context, arg1, arg2, arg3, arg4 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PromoteMirror", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PromoteMirror indicates an expected call of PromoteMirror.
func (mr *MockOrchestratorMockRecorder) PromoteMirror(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PromoteMirror", reflect.TypeOf((*MockOrchestrator)(nil).PromoteMirror), arg0, arg1, arg2, arg3, arg4)
}

// PublishVolume mocks base method.
func (m *MockOrchestrator) PublishVolume(arg0 context.Context, arg1 string, arg2 *utils.VolumePublishInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishVolume indicates an expected call of PublishVolume.
func (mr *MockOrchestratorMockRecorder) PublishVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishVolume", reflect.TypeOf((*MockOrchestrator)(nil).PublishVolume), arg0, arg1, arg2)
}

// ReadSnapshotsForVolume mocks base method.
func (m *MockOrchestrator) ReadSnapshotsForVolume(arg0 context.Context, arg1 string) ([]*storage.SnapshotExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadSnapshotsForVolume", arg0, arg1)
	ret0, _ := ret[0].([]*storage.SnapshotExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadSnapshotsForVolume indicates an expected call of ReadSnapshotsForVolume.
func (mr *MockOrchestratorMockRecorder) ReadSnapshotsForVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSnapshotsForVolume", reflect.TypeOf((*MockOrchestrator)(nil).ReadSnapshotsForVolume), arg0, arg1)
}

// ReestablishMirror mocks base method.
func (m *MockOrchestrator) ReestablishMirror(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReestablishMirror", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReestablishMirror indicates an expected call of ReestablishMirror.
func (mr *MockOrchestratorMockRecorder) ReestablishMirror(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReestablishMirror", reflect.TypeOf((*MockOrchestrator)(nil).ReestablishMirror), arg0, arg1, arg2, arg3)
}

// ReloadVolumes mocks base method.
func (m *MockOrchestrator) ReloadVolumes(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadVolumes", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadVolumes indicates an expected call of ReloadVolumes.
func (mr *MockOrchestratorMockRecorder) ReloadVolumes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadVolumes", reflect.TypeOf((*MockOrchestrator)(nil).ReloadVolumes), arg0)
}

// RemoveBackendConfigRef mocks base method.
func (m *MockOrchestrator) RemoveBackendConfigRef(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBackendConfigRef", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveBackendConfigRef indicates an expected call of RemoveBackendConfigRef.
func (mr *MockOrchestratorMockRecorder) RemoveBackendConfigRef(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBackendConfigRef", reflect.TypeOf((*MockOrchestrator)(nil).RemoveBackendConfigRef), arg0, arg1, arg2)
}

// ResizeVolume mocks base method.
func (m *MockOrchestrator) ResizeVolume(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResizeVolume indicates an expected call of ResizeVolume.
func (mr *MockOrchestratorMockRecorder) ResizeVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizeVolume", reflect.TypeOf((*MockOrchestrator)(nil).ResizeVolume), arg0, arg1, arg2)
}

// SetVolumeState mocks base method.
func (m *MockOrchestrator) SetVolumeState(arg0 context.Context, arg1 string, arg2 storage.VolumeState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVolumeState", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVolumeState indicates an expected call of SetVolumeState.
func (mr *MockOrchestratorMockRecorder) SetVolumeState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVolumeState", reflect.TypeOf((*MockOrchestrator)(nil).SetVolumeState), arg0, arg1, arg2)
}

// UnpublishVolume mocks base method.
func (m *MockOrchestrator) UnpublishVolume(arg0 context.Context, arg1 string, arg2 *utils.VolumePublishInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpublishVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnpublishVolume indicates an expected call of UnpublishVolume.
func (mr *MockOrchestratorMockRecorder) UnpublishVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpublishVolume", reflect.TypeOf((*MockOrchestrator)(nil).UnpublishVolume), arg0, arg1, arg2)
}

// UpdateBackend mocks base method.
func (m *MockOrchestrator) UpdateBackend(arg0 context.Context, arg1, arg2, arg3 string) (*storage.BackendExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBackend", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*storage.BackendExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBackend indicates an expected call of UpdateBackend.
func (mr *MockOrchestratorMockRecorder) UpdateBackend(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBackend", reflect.TypeOf((*MockOrchestrator)(nil).UpdateBackend), arg0, arg1, arg2, arg3)
}

// UpdateBackendByBackendUUID mocks base method.
func (m *MockOrchestrator) UpdateBackendByBackendUUID(arg0 context.Context, arg1, arg2, arg3, arg4 string) (*storage.BackendExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBackendByBackendUUID", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*storage.BackendExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBackendByBackendUUID indicates an expected call of UpdateBackendByBackendUUID.
func (mr *MockOrchestratorMockRecorder) UpdateBackendByBackendUUID(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBackendByBackendUUID", reflect.TypeOf((*MockOrchestrator)(nil).UpdateBackendByBackendUUID), arg0, arg1, arg2, arg3, arg4)
}

// UpdateBackendState mocks base method.
func (m *MockOrchestrator) UpdateBackendState(arg0 context.Context, arg1, arg2 string) (*storage.BackendExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBackendState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.BackendExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBackendState indicates an expected call of UpdateBackendState.
func (mr *MockOrchestratorMockRecorder) UpdateBackendState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBackendState", reflect.TypeOf((*MockOrchestrator)(nil).UpdateBackendState), arg0, arg1, arg2)
}
