// Code generated by MockGen. DO NOT EDIT.
// Source: internal/persistence/account_dao.go
//
// Generated by this command:
//
//	mockgen -source=internal/persistence/account_dao.go -destination=internal/persistence/mock/account_dao_mock.go
//

// Package mock_persistence is a generated GoMock package.
package mock_persistence

import (
	model "account-service/internal/model"
	persistence "account-service/internal/persistence"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAccountDAO is a mock of AccountDAO interface.
type MockAccountDAO struct {
	ctrl     *gomock.Controller
	recorder *MockAccountDAOMockRecorder
}

// MockAccountDAOMockRecorder is the mock recorder for MockAccountDAO.
type MockAccountDAOMockRecorder struct {
	mock *MockAccountDAO
}

// NewMockAccountDAO creates a new mock instance.
func NewMockAccountDAO(ctrl *gomock.Controller) *MockAccountDAO {
	mock := &MockAccountDAO{ctrl: ctrl}
	mock.recorder = &MockAccountDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountDAO) EXPECT() *MockAccountDAOMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountDAO) Create(account model.Account) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", account)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountDAOMockRecorder) Create(account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountDAO)(nil).Create), account)
}

// Delete mocks base method.
func (m *MockAccountDAO) Delete(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAccountDAOMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccountDAO)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockAccountDAO) Get(id int64) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAccountDAOMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAccountDAO)(nil).Get), id)
}

// GetByUserId mocks base method.
func (m *MockAccountDAO) GetByUserId(userId int64, page, size int) (persistence.Page[model.Account], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserId", userId, page, size)
	ret0, _ := ret[0].(persistence.Page[model.Account])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserId indicates an expected call of GetByUserId.
func (mr *MockAccountDAOMockRecorder) GetByUserId(userId, page, size any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserId", reflect.TypeOf((*MockAccountDAO)(nil).GetByUserId), userId, page, size)
}

// Update mocks base method.
func (m *MockAccountDAO) Update(account model.Account) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", account)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAccountDAOMockRecorder) Update(account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccountDAO)(nil).Update), account)
}
