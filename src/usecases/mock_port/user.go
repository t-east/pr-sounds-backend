// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/usecases/port/user.go

// Package mock_port is a generated GoMock package.
package mock_port

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/t-east/pr-sounds-backend/src/domains/entities"
)

// MockUserInputPort is a mock of UserInputPort interface.
type MockUserInputPort struct {
	ctrl     *gomock.Controller
	recorder *MockUserInputPortMockRecorder
}

// MockUserInputPortMockRecorder is the mock recorder for MockUserInputPort.
type MockUserInputPortMockRecorder struct {
	mock *MockUserInputPort
}

// NewMockUserInputPort creates a new mock instance.
func NewMockUserInputPort(ctrl *gomock.Controller) *MockUserInputPort {
	mock := &MockUserInputPort{ctrl: ctrl}
	mock.recorder = &MockUserInputPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserInputPort) EXPECT() *MockUserInputPortMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserInputPort) Create(arg0 *entities.User) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserInputPortMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserInputPort)(nil).Create), arg0)
}

// FindByID mocks base method.
func (m *MockUserInputPort) FindByID(arg0 uint) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserInputPortMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserInputPort)(nil).FindByID), arg0)
}

// Update mocks base method.
func (m *MockUserInputPort) Update(arg0 *entities.User) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserInputPortMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserInputPort)(nil).Update), arg0)
}

// MockUserOutputPort is a mock of UserOutputPort interface.
type MockUserOutputPort struct {
	ctrl     *gomock.Controller
	recorder *MockUserOutputPortMockRecorder
}

// MockUserOutputPortMockRecorder is the mock recorder for MockUserOutputPort.
type MockUserOutputPortMockRecorder struct {
	mock *MockUserOutputPort
}

// NewMockUserOutputPort creates a new mock instance.
func NewMockUserOutputPort(ctrl *gomock.Controller) *MockUserOutputPort {
	mock := &MockUserOutputPort{ctrl: ctrl}
	mock.recorder = &MockUserOutputPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserOutputPort) EXPECT() *MockUserOutputPortMockRecorder {
	return m.recorder
}

// Render mocks base method.
func (m *MockUserOutputPort) Render(arg0 *entities.User, arg1 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Render", arg0, arg1)
}

// Render indicates an expected call of Render.
func (mr *MockUserOutputPortMockRecorder) Render(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockUserOutputPort)(nil).Render), arg0, arg1)
}

// RenderError mocks base method.
func (m *MockUserOutputPort) RenderError(arg0 error, arg1 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RenderError", arg0, arg1)
}

// RenderError indicates an expected call of RenderError.
func (mr *MockUserOutputPortMockRecorder) RenderError(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderError", reflect.TypeOf((*MockUserOutputPort)(nil).RenderError), arg0, arg1)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(arg0 *entities.User) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), arg0)
}

// FindByEmail mocks base method.
func (m *MockUserRepository) FindByEmail(arg0 string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserRepositoryMockRecorder) FindByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), arg0)
}

// FindByID mocks base method.
func (m *MockUserRepository) FindByID(arg0 uint) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserRepositoryMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserRepository)(nil).FindByID), arg0)
}

// Update mocks base method.
func (m *MockUserRepository) Update(arg0 *entities.User) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), arg0)
}
