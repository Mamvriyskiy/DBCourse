// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepo is a mock of IUserRepo interface.
type MockIUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepoMockRecorder
}

// MockIUserRepoMockRecorder is the mock recorder for MockIUserRepo.
type MockIUserRepoMockRecorder struct {
	mock *MockIUserRepo
}

// NewMockIUserRepo creates a new mock instance.
func NewMockIUserRepo(ctrl *gomock.Controller) *MockIUserRepo {
	mock := &MockIUserRepo{ctrl: ctrl}
	mock.recorder = &MockIUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepo) EXPECT() *MockIUserRepoMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIUserRepo) CreateUser(user pkg.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserRepoMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserRepo)(nil).CreateUser), user)
}

// GetUser mocks base method.
func (m *MockIUserRepo) GetUser(login, password string) (pkg.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", login, password)
	ret0, _ := ret[0].(pkg.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIUserRepoMockRecorder) GetUser(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserRepo)(nil).GetUser), login, password)
}

// MockIHomeRepo is a mock of IHomeRepo interface.
type MockIHomeRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIHomeRepoMockRecorder
}

// MockIHomeRepoMockRecorder is the mock recorder for MockIHomeRepo.
type MockIHomeRepoMockRecorder struct {
	mock *MockIHomeRepo
}

// NewMockIHomeRepo creates a new mock instance.
func NewMockIHomeRepo(ctrl *gomock.Controller) *MockIHomeRepo {
	mock := &MockIHomeRepo{ctrl: ctrl}
	mock.recorder = &MockIHomeRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHomeRepo) EXPECT() *MockIHomeRepoMockRecorder {
	return m.recorder
}

// CreateHome mocks base method.
func (m *MockIHomeRepo) CreateHome(idUser int, home pkg.Home) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHome", idUser, home)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHome indicates an expected call of CreateHome.
func (mr *MockIHomeRepoMockRecorder) CreateHome(idUser, home interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHome", reflect.TypeOf((*MockIHomeRepo)(nil).CreateHome), idUser, home)
}

// DeleteHome mocks base method.
func (m *MockIHomeRepo) DeleteHome(homeID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHome", homeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHome indicates an expected call of DeleteHome.
func (mr *MockIHomeRepoMockRecorder) DeleteHome(homeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHome", reflect.TypeOf((*MockIHomeRepo)(nil).DeleteHome), homeID)
}

// GetHomeByID mocks base method.
func (m *MockIHomeRepo) GetHomeByID(homeID int) (pkg.Home, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHomeByID", homeID)
	ret0, _ := ret[0].(pkg.Home)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHomeByID indicates an expected call of GetHomeByID.
func (mr *MockIHomeRepoMockRecorder) GetHomeByID(homeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHomeByID", reflect.TypeOf((*MockIHomeRepo)(nil).GetHomeByID), homeID)
}

// ListUserHome mocks base method.
func (m *MockIHomeRepo) ListUserHome(userID int) ([]pkg.Home, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserHome", userID)
	ret0, _ := ret[0].([]pkg.Home)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserHome indicates an expected call of ListUserHome.
func (mr *MockIHomeRepoMockRecorder) ListUserHome(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserHome", reflect.TypeOf((*MockIHomeRepo)(nil).ListUserHome), userID)
}

// UpdateHome mocks base method.
func (m *MockIHomeRepo) UpdateHome(home pkg.Home) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHome", home)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHome indicates an expected call of UpdateHome.
func (mr *MockIHomeRepoMockRecorder) UpdateHome(home interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHome", reflect.TypeOf((*MockIHomeRepo)(nil).UpdateHome), home)
}

// MockIAccessHomeRepo is a mock of IAccessHomeRepo interface.
type MockIAccessHomeRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIAccessHomeRepoMockRecorder
}

// MockIAccessHomeRepoMockRecorder is the mock recorder for MockIAccessHomeRepo.
type MockIAccessHomeRepoMockRecorder struct {
	mock *MockIAccessHomeRepo
}

// NewMockIAccessHomeRepo creates a new mock instance.
func NewMockIAccessHomeRepo(ctrl *gomock.Controller) *MockIAccessHomeRepo {
	mock := &MockIAccessHomeRepo{ctrl: ctrl}
	mock.recorder = &MockIAccessHomeRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccessHomeRepo) EXPECT() *MockIAccessHomeRepoMockRecorder {
	return m.recorder
}

// AddOwner mocks base method.
func (m *MockIAccessHomeRepo) AddOwner(userID, homeID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOwner", userID, homeID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOwner indicates an expected call of AddOwner.
func (mr *MockIAccessHomeRepoMockRecorder) AddOwner(userID, homeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOwner", reflect.TypeOf((*MockIAccessHomeRepo)(nil).AddOwner), userID, homeID)
}

// AddUser mocks base method.
func (m *MockIAccessHomeRepo) AddUser(userID, accessLevel int, email string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", userID, accessLevel, email)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockIAccessHomeRepoMockRecorder) AddUser(userID, accessLevel, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockIAccessHomeRepo)(nil).AddUser), userID, accessLevel, email)
}

// DeleteUser mocks base method.
func (m *MockIAccessHomeRepo) DeleteUser(idUser int, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", idUser, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIAccessHomeRepoMockRecorder) DeleteUser(idUser, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIAccessHomeRepo)(nil).DeleteUser), idUser, email)
}

// GetListUserHome mocks base method.
func (m *MockIAccessHomeRepo) GetListUserHome(idHome int) ([]pkg.ClientHome, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListUserHome", idHome)
	ret0, _ := ret[0].([]pkg.ClientHome)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListUserHome indicates an expected call of GetListUserHome.
func (mr *MockIAccessHomeRepoMockRecorder) GetListUserHome(idHome interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListUserHome", reflect.TypeOf((*MockIAccessHomeRepo)(nil).GetListUserHome), idHome)
}

// UpdateLevel mocks base method.
func (m *MockIAccessHomeRepo) UpdateLevel(idUser int, access pkg.AddUserHome) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLevel", idUser, access)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLevel indicates an expected call of UpdateLevel.
func (mr *MockIAccessHomeRepoMockRecorder) UpdateLevel(idUser, access interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLevel", reflect.TypeOf((*MockIAccessHomeRepo)(nil).UpdateLevel), idUser, access)
}

// UpdateStatus mocks base method.
func (m *MockIAccessHomeRepo) UpdateStatus(idUser int, access pkg.AccessHome) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", idUser, access)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockIAccessHomeRepoMockRecorder) UpdateStatus(idUser, access interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockIAccessHomeRepo)(nil).UpdateStatus), idUser, access)
}

// MockIDeviceRepo is a mock of IDeviceRepo interface.
type MockIDeviceRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIDeviceRepoMockRecorder
}

// MockIDeviceRepoMockRecorder is the mock recorder for MockIDeviceRepo.
type MockIDeviceRepoMockRecorder struct {
	mock *MockIDeviceRepo
}

// NewMockIDeviceRepo creates a new mock instance.
func NewMockIDeviceRepo(ctrl *gomock.Controller) *MockIDeviceRepo {
	mock := &MockIDeviceRepo{ctrl: ctrl}
	mock.recorder = &MockIDeviceRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDeviceRepo) EXPECT() *MockIDeviceRepoMockRecorder {
	return m.recorder
}

// CreateDevice mocks base method.
func (m *MockIDeviceRepo) CreateDevice(homeID int, device *pkg.Devices) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevice", homeID, device)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockIDeviceRepoMockRecorder) CreateDevice(homeID, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockIDeviceRepo)(nil).CreateDevice), homeID, device)
}

// DeleteDevice mocks base method.
func (m *MockIDeviceRepo) DeleteDevice(idDevice int, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDevice", idDevice, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDevice indicates an expected call of DeleteDevice.
func (mr *MockIDeviceRepoMockRecorder) DeleteDevice(idDevice, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDevice", reflect.TypeOf((*MockIDeviceRepo)(nil).DeleteDevice), idDevice, name)
}

// GetDeviceByID mocks base method.
func (m *MockIDeviceRepo) GetDeviceByID(deviceID int) (pkg.Devices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByID", deviceID)
	ret0, _ := ret[0].(pkg.Devices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByID indicates an expected call of GetDeviceByID.
func (mr *MockIDeviceRepoMockRecorder) GetDeviceByID(deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByID", reflect.TypeOf((*MockIDeviceRepo)(nil).GetDeviceByID), deviceID)
}

// MockIHistoryDeviceRepo is a mock of IHistoryDeviceRepo interface.
type MockIHistoryDeviceRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIHistoryDeviceRepoMockRecorder
}

// MockIHistoryDeviceRepoMockRecorder is the mock recorder for MockIHistoryDeviceRepo.
type MockIHistoryDeviceRepoMockRecorder struct {
	mock *MockIHistoryDeviceRepo
}

// NewMockIHistoryDeviceRepo creates a new mock instance.
func NewMockIHistoryDeviceRepo(ctrl *gomock.Controller) *MockIHistoryDeviceRepo {
	mock := &MockIHistoryDeviceRepo{ctrl: ctrl}
	mock.recorder = &MockIHistoryDeviceRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHistoryDeviceRepo) EXPECT() *MockIHistoryDeviceRepoMockRecorder {
	return m.recorder
}

// CreateDeviceHistory mocks base method.
func (m *MockIHistoryDeviceRepo) CreateDeviceHistory(userID int, history pkg.AddHistory) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeviceHistory", userID, history)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeviceHistory indicates an expected call of CreateDeviceHistory.
func (mr *MockIHistoryDeviceRepoMockRecorder) CreateDeviceHistory(userID, history interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeviceHistory", reflect.TypeOf((*MockIHistoryDeviceRepo)(nil).CreateDeviceHistory), userID, history)
}

// GetDeviceHistory mocks base method.
func (m *MockIHistoryDeviceRepo) GetDeviceHistory(userID int, name string) ([]pkg.DevicesHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceHistory", userID, name)
	ret0, _ := ret[0].([]pkg.DevicesHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceHistory indicates an expected call of GetDeviceHistory.
func (mr *MockIHistoryDeviceRepoMockRecorder) GetDeviceHistory(userID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceHistory", reflect.TypeOf((*MockIHistoryDeviceRepo)(nil).GetDeviceHistory), userID, name)
}
