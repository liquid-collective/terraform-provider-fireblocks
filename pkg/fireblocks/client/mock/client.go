// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/fireblocks/client/client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddExternalWalletAsset mocks base method.
func (m *MockClient) AddExternalWalletAsset(ctx context.Context, walletID, assetID string, msg *client.AddExternalWalletAssetMsg) (*client.ExternalWalletAsset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddExternalWalletAsset", ctx, walletID, assetID, msg)
	ret0, _ := ret[0].(*client.ExternalWalletAsset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddExternalWalletAsset indicates an expected call of AddExternalWalletAsset.
func (mr *MockClientMockRecorder) AddExternalWalletAsset(ctx, walletID, assetID, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddExternalWalletAsset", reflect.TypeOf((*MockClient)(nil).AddExternalWalletAsset), ctx, walletID, assetID, msg)
}

// CancelTransaction mocks base method.
func (m *MockClient) CancelTransaction(ctx context.Context, txID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelTransaction", ctx, txID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelTransaction indicates an expected call of CancelTransaction.
func (mr *MockClientMockRecorder) CancelTransaction(ctx, txID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelTransaction", reflect.TypeOf((*MockClient)(nil).CancelTransaction), ctx, txID)
}

// CreateExternalWallet mocks base method.
func (m *MockClient) CreateExternalWallet(ctx context.Context, msg *client.CreateExternalWalletMsg) (*client.ExternalWallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExternalWallet", ctx, msg)
	ret0, _ := ret[0].(*client.ExternalWallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExternalWallet indicates an expected call of CreateExternalWallet.
func (mr *MockClientMockRecorder) CreateExternalWallet(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExternalWallet", reflect.TypeOf((*MockClient)(nil).CreateExternalWallet), ctx, msg)
}

// CreateTransaction mocks base method.
func (m *MockClient) CreateTransaction(ctx context.Context, msg *client.CreateTransactionMsg) (*client.CreateTransactionRespMsg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, msg)
	ret0, _ := ret[0].(*client.CreateTransactionRespMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockClientMockRecorder) CreateTransaction(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockClient)(nil).CreateTransaction), ctx, msg)
}

// CreateVaultAccount mocks base method.
func (m *MockClient) CreateVaultAccount(ctx context.Context, msg *client.CreateVaultAccountMsg) (*client.VaultAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVaultAccount", ctx, msg)
	ret0, _ := ret[0].(*client.VaultAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVaultAccount indicates an expected call of CreateVaultAccount.
func (mr *MockClientMockRecorder) CreateVaultAccount(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVaultAccount", reflect.TypeOf((*MockClient)(nil).CreateVaultAccount), ctx, msg)
}

// CreateVaultAccountAsset mocks base method.
func (m *MockClient) CreateVaultAccountAsset(ctx context.Context, vaultID, assetID string) (*client.CreateVaultAssetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVaultAccountAsset", ctx, vaultID, assetID)
	ret0, _ := ret[0].(*client.CreateVaultAssetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVaultAccountAsset indicates an expected call of CreateVaultAccountAsset.
func (mr *MockClientMockRecorder) CreateVaultAccountAsset(ctx, vaultID, assetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVaultAccountAsset", reflect.TypeOf((*MockClient)(nil).CreateVaultAccountAsset), ctx, vaultID, assetID)
}

// DeleteExternalWallet mocks base method.
func (m *MockClient) DeleteExternalWallet(ctx context.Context, walletID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExternalWallet", ctx, walletID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExternalWallet indicates an expected call of DeleteExternalWallet.
func (mr *MockClientMockRecorder) DeleteExternalWallet(ctx, walletID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExternalWallet", reflect.TypeOf((*MockClient)(nil).DeleteExternalWallet), ctx, walletID)
}

// DeleteExternalWalletAsset mocks base method.
func (m *MockClient) DeleteExternalWalletAsset(ctx context.Context, walletID, assetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExternalWalletAsset", ctx, walletID, assetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExternalWalletAsset indicates an expected call of DeleteExternalWalletAsset.
func (mr *MockClientMockRecorder) DeleteExternalWalletAsset(ctx, walletID, assetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExternalWalletAsset", reflect.TypeOf((*MockClient)(nil).DeleteExternalWalletAsset), ctx, walletID, assetID)
}

// GetExternalWallet mocks base method.
func (m *MockClient) GetExternalWallet(ctx context.Context, walletID string) (*client.ExternalWallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalWallet", ctx, walletID)
	ret0, _ := ret[0].(*client.ExternalWallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalWallet indicates an expected call of GetExternalWallet.
func (mr *MockClientMockRecorder) GetExternalWallet(ctx, walletID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalWallet", reflect.TypeOf((*MockClient)(nil).GetExternalWallet), ctx, walletID)
}

// GetExternalWalletAsset mocks base method.
func (m *MockClient) GetExternalWalletAsset(ctx context.Context, walletID, assetID string) (*client.ExternalWalletAsset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalWalletAsset", ctx, walletID, assetID)
	ret0, _ := ret[0].(*client.ExternalWalletAsset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalWalletAsset indicates an expected call of GetExternalWalletAsset.
func (mr *MockClientMockRecorder) GetExternalWalletAsset(ctx, walletID, assetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalWalletAsset", reflect.TypeOf((*MockClient)(nil).GetExternalWalletAsset), ctx, walletID, assetID)
}

// GetTransaction mocks base method.
func (m *MockClient) GetTransaction(ctx context.Context, txID string) (*client.TransactionMsg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", ctx, txID)
	ret0, _ := ret[0].(*client.TransactionMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockClientMockRecorder) GetTransaction(ctx, txID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockClient)(nil).GetTransaction), ctx, txID)
}

// GetVaultAccount mocks base method.
func (m *MockClient) GetVaultAccount(ctx context.Context, vaultAccountID string) (*client.VaultAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVaultAccount", ctx, vaultAccountID)
	ret0, _ := ret[0].(*client.VaultAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVaultAccount indicates an expected call of GetVaultAccount.
func (mr *MockClientMockRecorder) GetVaultAccount(ctx, vaultAccountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVaultAccount", reflect.TypeOf((*MockClient)(nil).GetVaultAccount), ctx, vaultAccountID)
}

// GetVaultAccountAssetBalance mocks base method.
func (m *MockClient) GetVaultAccountAssetBalance(ctx context.Context, vaultID, assetID string) (*client.VaultAsset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVaultAccountAssetBalance", ctx, vaultID, assetID)
	ret0, _ := ret[0].(*client.VaultAsset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVaultAccountAssetBalance indicates an expected call of GetVaultAccountAssetBalance.
func (mr *MockClientMockRecorder) GetVaultAccountAssetBalance(ctx, vaultID, assetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVaultAccountAssetBalance", reflect.TypeOf((*MockClient)(nil).GetVaultAccountAssetBalance), ctx, vaultID, assetID)
}

// HideVaultAccount mocks base method.
func (m *MockClient) HideVaultAccount(ctx context.Context, vaultAccountID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HideVaultAccount", ctx, vaultAccountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// HideVaultAccount indicates an expected call of HideVaultAccount.
func (mr *MockClientMockRecorder) HideVaultAccount(ctx, vaultAccountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HideVaultAccount", reflect.TypeOf((*MockClient)(nil).HideVaultAccount), ctx, vaultAccountID)
}

// ListUsers mocks base method.
func (m *MockClient) ListUsers(ctx context.Context) ([]*client.UserMsg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx)
	ret0, _ := ret[0].([]*client.UserMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockClientMockRecorder) ListUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockClient)(nil).ListUsers), ctx)
}

// ListVaultAccountAssetAddresses mocks base method.
func (m *MockClient) ListVaultAccountAssetAddresses(ctx context.Context, vaultID, assetID string) ([]*client.VaultAccountAssetAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVaultAccountAssetAddresses", ctx, vaultID, assetID)
	ret0, _ := ret[0].([]*client.VaultAccountAssetAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVaultAccountAssetAddresses indicates an expected call of ListVaultAccountAssetAddresses.
func (mr *MockClientMockRecorder) ListVaultAccountAssetAddresses(ctx, vaultID, assetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVaultAccountAssetAddresses", reflect.TypeOf((*MockClient)(nil).ListVaultAccountAssetAddresses), ctx, vaultID, assetID)
}

// UnhideVaultAccount mocks base method.
func (m *MockClient) UnhideVaultAccount(ctx context.Context, vaultAccountID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnhideVaultAccount", ctx, vaultAccountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnhideVaultAccount indicates an expected call of UnhideVaultAccount.
func (mr *MockClientMockRecorder) UnhideVaultAccount(ctx, vaultAccountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnhideVaultAccount", reflect.TypeOf((*MockClient)(nil).UnhideVaultAccount), ctx, vaultAccountID)
}

// UpdateVaultAccount mocks base method.
func (m *MockClient) UpdateVaultAccount(ctx context.Context, vaultAccountID string, msg *client.UpdateVaultAccountMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVaultAccount", ctx, vaultAccountID, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVaultAccount indicates an expected call of UpdateVaultAccount.
func (mr *MockClientMockRecorder) UpdateVaultAccount(ctx, vaultAccountID, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVaultAccount", reflect.TypeOf((*MockClient)(nil).UpdateVaultAccount), ctx, vaultAccountID, msg)
}
