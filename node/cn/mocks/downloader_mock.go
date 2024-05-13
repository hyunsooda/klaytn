// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/klaytn/klaytn/node/cn (interfaces: ProtocolManagerDownloader)

// Package mocks is a generated GoMock package.
package mocks

import (
	big "math/big"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kaia "github.com/klaytn/klaytn"
	types "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	downloader "github.com/klaytn/klaytn/datasync/downloader"
	snap "github.com/klaytn/klaytn/node/cn/snap"
	reward "github.com/klaytn/klaytn/reward"
)

// MockProtocolManagerDownloader is a mock of ProtocolManagerDownloader interface.
type MockProtocolManagerDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockProtocolManagerDownloaderMockRecorder
}

// MockProtocolManagerDownloaderMockRecorder is the mock recorder for MockProtocolManagerDownloader.
type MockProtocolManagerDownloaderMockRecorder struct {
	mock *MockProtocolManagerDownloader
}

// NewMockProtocolManagerDownloader creates a new mock instance.
func NewMockProtocolManagerDownloader(ctrl *gomock.Controller) *MockProtocolManagerDownloader {
	mock := &MockProtocolManagerDownloader{ctrl: ctrl}
	mock.recorder = &MockProtocolManagerDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProtocolManagerDownloader) EXPECT() *MockProtocolManagerDownloaderMockRecorder {
	return m.recorder
}

// Cancel mocks base method.
func (m *MockProtocolManagerDownloader) Cancel() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cancel")
}

// Cancel indicates an expected call of Cancel.
func (mr *MockProtocolManagerDownloaderMockRecorder) Cancel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).Cancel))
}

// DeliverBodies mocks base method.
func (m *MockProtocolManagerDownloader) DeliverBodies(arg0 string, arg1 [][]*types.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverBodies", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeliverBodies indicates an expected call of DeliverBodies.
func (mr *MockProtocolManagerDownloaderMockRecorder) DeliverBodies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverBodies", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).DeliverBodies), arg0, arg1)
}

// DeliverHeaders mocks base method.
func (m *MockProtocolManagerDownloader) DeliverHeaders(arg0 string, arg1 []*types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverHeaders", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeliverHeaders indicates an expected call of DeliverHeaders.
func (mr *MockProtocolManagerDownloaderMockRecorder) DeliverHeaders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverHeaders", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).DeliverHeaders), arg0, arg1)
}

// DeliverNodeData mocks base method.
func (m *MockProtocolManagerDownloader) DeliverNodeData(arg0 string, arg1 [][]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverNodeData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeliverNodeData indicates an expected call of DeliverNodeData.
func (mr *MockProtocolManagerDownloaderMockRecorder) DeliverNodeData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverNodeData", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).DeliverNodeData), arg0, arg1)
}

// DeliverReceipts mocks base method.
func (m *MockProtocolManagerDownloader) DeliverReceipts(arg0 string, arg1 [][]*types.Receipt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverReceipts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeliverReceipts indicates an expected call of DeliverReceipts.
func (mr *MockProtocolManagerDownloaderMockRecorder) DeliverReceipts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverReceipts", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).DeliverReceipts), arg0, arg1)
}

// DeliverSnapPacket mocks base method.
func (m *MockProtocolManagerDownloader) DeliverSnapPacket(arg0 *snap.Peer, arg1 snap.Packet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverSnapPacket", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeliverSnapPacket indicates an expected call of DeliverSnapPacket.
func (mr *MockProtocolManagerDownloaderMockRecorder) DeliverSnapPacket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverSnapPacket", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).DeliverSnapPacket), arg0, arg1)
}

// DeliverStakingInfos mocks base method.
func (m *MockProtocolManagerDownloader) DeliverStakingInfos(arg0 string, arg1 []*reward.StakingInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverStakingInfos", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeliverStakingInfos indicates an expected call of DeliverStakingInfos.
func (mr *MockProtocolManagerDownloaderMockRecorder) DeliverStakingInfos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverStakingInfos", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).DeliverStakingInfos), arg0, arg1)
}

// GetSnapSyncer mocks base method.
func (m *MockProtocolManagerDownloader) GetSnapSyncer() *snap.Syncer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapSyncer")
	ret0, _ := ret[0].(*snap.Syncer)
	return ret0
}

// GetSnapSyncer indicates an expected call of GetSnapSyncer.
func (mr *MockProtocolManagerDownloaderMockRecorder) GetSnapSyncer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapSyncer", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).GetSnapSyncer))
}

// Progress mocks base method.
func (m *MockProtocolManagerDownloader) Progress() kaia.SyncProgress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Progress")
	ret0, _ := ret[0].(kaia.SyncProgress)
	return ret0
}

// Progress indicates an expected call of Progress.
func (mr *MockProtocolManagerDownloaderMockRecorder) Progress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Progress", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).Progress))
}

// RegisterPeer mocks base method.
func (m *MockProtocolManagerDownloader) RegisterPeer(arg0 string, arg1 int, arg2 downloader.Peer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterPeer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterPeer indicates an expected call of RegisterPeer.
func (mr *MockProtocolManagerDownloaderMockRecorder) RegisterPeer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPeer", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).RegisterPeer), arg0, arg1, arg2)
}

// SyncStakingInfo mocks base method.
func (m *MockProtocolManagerDownloader) SyncStakingInfo(arg0 string, arg1, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStakingInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStakingInfo indicates an expected call of SyncStakingInfo.
func (mr *MockProtocolManagerDownloaderMockRecorder) SyncStakingInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStakingInfo", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).SyncStakingInfo), arg0, arg1, arg2)
}

// SyncStakingInfoStatus mocks base method.
func (m *MockProtocolManagerDownloader) SyncStakingInfoStatus() *downloader.SyncingStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStakingInfoStatus")
	ret0, _ := ret[0].(*downloader.SyncingStatus)
	return ret0
}

// SyncStakingInfoStatus indicates an expected call of SyncStakingInfoStatus.
func (mr *MockProtocolManagerDownloaderMockRecorder) SyncStakingInfoStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStakingInfoStatus", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).SyncStakingInfoStatus))
}

// Synchronise mocks base method.
func (m *MockProtocolManagerDownloader) Synchronise(arg0 string, arg1 common.Hash, arg2 *big.Int, arg3 downloader.SyncMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Synchronise", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Synchronise indicates an expected call of Synchronise.
func (mr *MockProtocolManagerDownloaderMockRecorder) Synchronise(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Synchronise", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).Synchronise), arg0, arg1, arg2, arg3)
}

// Terminate mocks base method.
func (m *MockProtocolManagerDownloader) Terminate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Terminate")
}

// Terminate indicates an expected call of Terminate.
func (mr *MockProtocolManagerDownloaderMockRecorder) Terminate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terminate", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).Terminate))
}

// UnregisterPeer mocks base method.
func (m *MockProtocolManagerDownloader) UnregisterPeer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterPeer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterPeer indicates an expected call of UnregisterPeer.
func (mr *MockProtocolManagerDownloaderMockRecorder) UnregisterPeer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterPeer", reflect.TypeOf((*MockProtocolManagerDownloader)(nil).UnregisterPeer), arg0)
}
