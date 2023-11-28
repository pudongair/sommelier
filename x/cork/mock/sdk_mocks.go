/* 
// Code generated by MockGen. DO NOT EDIT.

Interfaces from source: x/cork/types/expected_keepers.go
- StakingKeeper 
- GravityKeeper

Interfaces from source: github.com/cosmos/cosmos-sdk/x/staking/types/exported.go
- DelegatorI
- ValidatorI
*/
package mock

import (
	reflect "reflect"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "github.com/golang/mock/gomock"
	crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
)

// MockDelegationI is a mock of DelegationI interface.
type MockDelegationI struct {
	ctrl     *gomock.Controller
	recorder *MockDelegationIMockRecorder
}

// MockDelegationIMockRecorder is the mock recorder for MockDelegationI.
type MockDelegationIMockRecorder struct {
	mock *MockDelegationI
}

// NewMockDelegationI creates a new mock instance.
func NewMockDelegationI(ctrl *gomock.Controller) *MockDelegationI {
	mock := &MockDelegationI{ctrl: ctrl}
	mock.recorder = &MockDelegationIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelegationI) EXPECT() *MockDelegationIMockRecorder {
	return m.recorder
}

// GetDelegatorAddr mocks base method.
func (m *MockDelegationI) GetDelegatorAddr() sdk.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDelegatorAddr")
	ret0, _ := ret[0].(sdk.AccAddress)
	return ret0
}

// GetDelegatorAddr indicates an expected call of GetDelegatorAddr.
func (mr *MockDelegationIMockRecorder) GetDelegatorAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDelegatorAddr", reflect.TypeOf((*MockDelegationI)(nil).GetDelegatorAddr))
}

// GetShares mocks base method.
func (m *MockDelegationI) GetShares() sdk.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShares")
	ret0, _ := ret[0].(sdk.Dec)
	return ret0
}

// GetShares indicates an expected call of GetShares.
func (mr *MockDelegationIMockRecorder) GetShares() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShares", reflect.TypeOf((*MockDelegationI)(nil).GetShares))
}

// GetValidatorAddr mocks base method.
func (m *MockDelegationI) GetValidatorAddr() sdk.ValAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorAddr")
	ret0, _ := ret[0].(sdk.ValAddress)
	return ret0
}

// GetValidatorAddr indicates an expected call of GetValidatorAddr.
func (mr *MockDelegationIMockRecorder) GetValidatorAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorAddr", reflect.TypeOf((*MockDelegationI)(nil).GetValidatorAddr))
}

// MockValidatorI is a mock of ValidatorI interface.
type MockValidatorI struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorIMockRecorder
}

// MockValidatorIMockRecorder is the mock recorder for MockValidatorI.
type MockValidatorIMockRecorder struct {
	mock *MockValidatorI
}

// NewMockValidatorI creates a new mock instance.
func NewMockValidatorI(ctrl *gomock.Controller) *MockValidatorI {
	mock := &MockValidatorI{ctrl: ctrl}
	mock.recorder = &MockValidatorIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorI) EXPECT() *MockValidatorIMockRecorder {
	return m.recorder
}

// ConsPubKey mocks base method.
func (m *MockValidatorI) ConsPubKey() (cryptotypes.PubKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsPubKey")
	ret0, _ := ret[0].(cryptotypes.PubKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsPubKey indicates an expected call of ConsPubKey.
func (mr *MockValidatorIMockRecorder) ConsPubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsPubKey", reflect.TypeOf((*MockValidatorI)(nil).ConsPubKey))
}

// GetBondedTokens mocks base method.
func (m *MockValidatorI) GetBondedTokens() sdk.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBondedTokens")
	ret0, _ := ret[0].(sdk.Int)
	return ret0
}

// GetBondedTokens indicates an expected call of GetBondedTokens.
func (mr *MockValidatorIMockRecorder) GetBondedTokens() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBondedTokens", reflect.TypeOf((*MockValidatorI)(nil).GetBondedTokens))
}

// GetCommission mocks base method.
func (m *MockValidatorI) GetCommission() sdk.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommission")
	ret0, _ := ret[0].(sdk.Dec)
	return ret0
}

// GetCommission indicates an expected call of GetCommission.
func (mr *MockValidatorIMockRecorder) GetCommission() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommission", reflect.TypeOf((*MockValidatorI)(nil).GetCommission))
}

// GetConsAddr mocks base method.
func (m *MockValidatorI) GetConsAddr() (sdk.ConsAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsAddr")
	ret0, _ := ret[0].(sdk.ConsAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsAddr indicates an expected call of GetConsAddr.
func (mr *MockValidatorIMockRecorder) GetConsAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsAddr", reflect.TypeOf((*MockValidatorI)(nil).GetConsAddr))
}

// GetConsensusPower mocks base method.
func (m *MockValidatorI) GetConsensusPower(arg0 sdk.Int) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsensusPower", arg0)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetConsensusPower indicates an expected call of GetConsensusPower.
func (mr *MockValidatorIMockRecorder) GetConsensusPower(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsensusPower", reflect.TypeOf((*MockValidatorI)(nil).GetConsensusPower), arg0)
}

// GetDelegatorShares mocks base method.
func (m *MockValidatorI) GetDelegatorShares() sdk.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDelegatorShares")
	ret0, _ := ret[0].(sdk.Dec)
	return ret0
}

// GetDelegatorShares indicates an expected call of GetDelegatorShares.
func (mr *MockValidatorIMockRecorder) GetDelegatorShares() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDelegatorShares", reflect.TypeOf((*MockValidatorI)(nil).GetDelegatorShares))
}

// GetMinSelfDelegation mocks base method.
func (m *MockValidatorI) GetMinSelfDelegation() sdk.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinSelfDelegation")
	ret0, _ := ret[0].(sdk.Int)
	return ret0
}

// GetMinSelfDelegation indicates an expected call of GetMinSelfDelegation.
func (mr *MockValidatorIMockRecorder) GetMinSelfDelegation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinSelfDelegation", reflect.TypeOf((*MockValidatorI)(nil).GetMinSelfDelegation))
}

// GetMoniker mocks base method.
func (m *MockValidatorI) GetMoniker() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMoniker")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMoniker indicates an expected call of GetMoniker.
func (mr *MockValidatorIMockRecorder) GetMoniker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMoniker", reflect.TypeOf((*MockValidatorI)(nil).GetMoniker))
}

// GetOperator mocks base method.
func (m *MockValidatorI) GetOperator() sdk.ValAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperator")
	ret0, _ := ret[0].(sdk.ValAddress)
	return ret0
}

// GetOperator indicates an expected call of GetOperator.
func (mr *MockValidatorIMockRecorder) GetOperator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperator", reflect.TypeOf((*MockValidatorI)(nil).GetOperator))
}

// GetStatus mocks base method.
func (m *MockValidatorI) GetStatus() stakingtypes.BondStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(stakingtypes.BondStatus)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockValidatorIMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockValidatorI)(nil).GetStatus))
}

// GetTokens mocks base method.
func (m *MockValidatorI) GetTokens() sdk.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokens")
	ret0, _ := ret[0].(sdk.Int)
	return ret0
}

// GetTokens indicates an expected call of GetTokens.
func (mr *MockValidatorIMockRecorder) GetTokens() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokens", reflect.TypeOf((*MockValidatorI)(nil).GetTokens))
}

// IsBonded mocks base method.
func (m *MockValidatorI) IsBonded() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBonded")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBonded indicates an expected call of IsBonded.
func (mr *MockValidatorIMockRecorder) IsBonded() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBonded", reflect.TypeOf((*MockValidatorI)(nil).IsBonded))
}

// IsJailed mocks base method.
func (m *MockValidatorI) IsJailed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsJailed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsJailed indicates an expected call of IsJailed.
func (mr *MockValidatorIMockRecorder) IsJailed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsJailed", reflect.TypeOf((*MockValidatorI)(nil).IsJailed))
}

// IsUnbonded mocks base method.
func (m *MockValidatorI) IsUnbonded() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUnbonded")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUnbonded indicates an expected call of IsUnbonded.
func (mr *MockValidatorIMockRecorder) IsUnbonded() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnbonded", reflect.TypeOf((*MockValidatorI)(nil).IsUnbonded))
}

// IsUnbonding mocks base method.
func (m *MockValidatorI) IsUnbonding() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUnbonding")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUnbonding indicates an expected call of IsUnbonding.
func (mr *MockValidatorIMockRecorder) IsUnbonding() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnbonding", reflect.TypeOf((*MockValidatorI)(nil).IsUnbonding))
}

// SharesFromTokens mocks base method.
func (m *MockValidatorI) SharesFromTokens(amt sdk.Int) (sdk.Dec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SharesFromTokens", amt)
	ret0, _ := ret[0].(sdk.Dec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SharesFromTokens indicates an expected call of SharesFromTokens.
func (mr *MockValidatorIMockRecorder) SharesFromTokens(amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SharesFromTokens", reflect.TypeOf((*MockValidatorI)(nil).SharesFromTokens), amt)
}

// SharesFromTokensTruncated mocks base method.
func (m *MockValidatorI) SharesFromTokensTruncated(amt sdk.Int) (sdk.Dec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SharesFromTokensTruncated", amt)
	ret0, _ := ret[0].(sdk.Dec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SharesFromTokensTruncated indicates an expected call of SharesFromTokensTruncated.
func (mr *MockValidatorIMockRecorder) SharesFromTokensTruncated(amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SharesFromTokensTruncated", reflect.TypeOf((*MockValidatorI)(nil).SharesFromTokensTruncated), amt)
}

// TmConsPublicKey mocks base method.
func (m *MockValidatorI) TmConsPublicKey() (crypto.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TmConsPublicKey")
	ret0, _ := ret[0].(crypto.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TmConsPublicKey indicates an expected call of TmConsPublicKey.
func (mr *MockValidatorIMockRecorder) TmConsPublicKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TmConsPublicKey", reflect.TypeOf((*MockValidatorI)(nil).TmConsPublicKey))
}

// TokensFromShares mocks base method.
func (m *MockValidatorI) TokensFromShares(arg0 sdk.Dec) sdk.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokensFromShares", arg0)
	ret0, _ := ret[0].(sdk.Dec)
	return ret0
}

// TokensFromShares indicates an expected call of TokensFromShares.
func (mr *MockValidatorIMockRecorder) TokensFromShares(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokensFromShares", reflect.TypeOf((*MockValidatorI)(nil).TokensFromShares), arg0)
}

// TokensFromSharesRoundUp mocks base method.
func (m *MockValidatorI) TokensFromSharesRoundUp(arg0 sdk.Dec) sdk.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokensFromSharesRoundUp", arg0)
	ret0, _ := ret[0].(sdk.Dec)
	return ret0
}

// TokensFromSharesRoundUp indicates an expected call of TokensFromSharesRoundUp.
func (mr *MockValidatorIMockRecorder) TokensFromSharesRoundUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokensFromSharesRoundUp", reflect.TypeOf((*MockValidatorI)(nil).TokensFromSharesRoundUp), arg0)
}

// TokensFromSharesTruncated mocks base method.
func (m *MockValidatorI) TokensFromSharesTruncated(arg0 sdk.Dec) sdk.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokensFromSharesTruncated", arg0)
	ret0, _ := ret[0].(sdk.Dec)
	return ret0
}

// TokensFromSharesTruncated indicates an expected call of TokensFromSharesTruncated.
func (mr *MockValidatorIMockRecorder) TokensFromSharesTruncated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokensFromSharesTruncated", reflect.TypeOf((*MockValidatorI)(nil).TokensFromSharesTruncated), arg0)
}