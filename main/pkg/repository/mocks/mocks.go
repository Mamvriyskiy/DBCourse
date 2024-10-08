// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
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

// AddCode mocks base method.
func (m *MockIUserRepo) AddCode(email pkg.Email) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCode", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCode indicates an expected call of AddCode.
func (mr *MockIUserRepoMockRecorder) AddCode(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCode", reflect.TypeOf((*MockIUserRepo)(nil).AddCode), email)
}

// ChangePassword mocks base method.
func (m *MockIUserRepo) ChangePassword(password, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", password, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockIUserRepoMockRecorder) ChangePassword(password, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockIUserRepo)(nil).ChangePassword), password, token)
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

// GetCode mocks base method.
func (m *MockIUserRepo) GetCode(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCode", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCode indicates an expected call of GetCode.
func (mr *MockIUserRepoMockRecorder) GetCode(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCode", reflect.TypeOf((*MockIUserRepo)(nil).GetCode), token)
}

// GetUser mocks base method.
func (m *MockIUserRepo) GetUser(email, password string) (pkg.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", email, password)
	ret0, _ := ret[0].(pkg.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIUserRepoMockRecorder) GetUser(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserRepo)(nil).GetUser), email, password)
}

// GetUserByEmail mocks base method.
func (m *MockIUserRepo) GetUserByEmail(email string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockIUserRepoMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockIUserRepo)(nil).GetUserByEmail), email)
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
func (m *MockIHomeRepo) CreateHome(home pkg.Home) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHome", home)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHome indicates an expected call of CreateHome.
func (mr *MockIHomeRepoMockRecorder) CreateHome(home interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHome", reflect.TypeOf((*MockIHomeRepo)(nil).CreateHome), home)
}

// DeleteHome mocks base method.
func (m *MockIHomeRepo) DeleteHome(homeID int, homeName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHome", homeID, homeName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHome indicates an expected call of DeleteHome.
func (mr *MockIHomeRepoMockRecorder) DeleteHome(homeID, homeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHome", reflect.TypeOf((*MockIHomeRepo)(nil).DeleteHome), homeID, homeName)
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
func (m *MockIHomeRepo) UpdateHome(home pkg.UpdateNameHome) error {
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
func (m *MockIAccessHomeRepo) AddUser(userID int, access pkg.Access) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", userID, access)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockIAccessHomeRepoMockRecorder) AddUser(userID, access interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockIAccessHomeRepo)(nil).AddUser), userID, access)
}

// DeleteUser mocks base method.
func (m *MockIAccessHomeRepo) DeleteUser(idUser int, access pkg.Access) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", idUser, access)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIAccessHomeRepoMockRecorder) DeleteUser(idUser, access interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIAccessHomeRepo)(nil).DeleteUser), idUser, access)
}

// GetListUserHome mocks base method.
func (m *MockIAccessHomeRepo) GetListUserHome(userID int) ([]pkg.ClientHome, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListUserHome", userID)
	ret0, _ := ret[0].([]pkg.ClientHome)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListUserHome indicates an expected call of GetListUserHome.
func (mr *MockIAccessHomeRepoMockRecorder) GetListUserHome(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListUserHome", reflect.TypeOf((*MockIAccessHomeRepo)(nil).GetListUserHome), userID)
}

// UpdateLevel mocks base method.
func (m *MockIAccessHomeRepo) UpdateLevel(userID int, updateAccess pkg.Access) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLevel", userID, updateAccess)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLevel indicates an expected call of UpdateLevel.
func (mr *MockIAccessHomeRepoMockRecorder) UpdateLevel(userID, updateAccess interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLevel", reflect.TypeOf((*MockIAccessHomeRepo)(nil).UpdateLevel), userID, updateAccess)
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
func (m *MockIDeviceRepo) CreateDevice(userID int, device *pkg.Devices, character pkg.DeviceCharacteristics, typeCharacter pkg.TypeCharacter) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevice", userID, device, character, typeCharacter)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockIDeviceRepoMockRecorder) CreateDevice(userID, device, character, typeCharacter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockIDeviceRepo)(nil).CreateDevice), userID, device, character, typeCharacter)
}

// DeleteDevice mocks base method.
func (m *MockIDeviceRepo) DeleteDevice(idDevice int, name, home string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDevice", idDevice, name, home)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDevice indicates an expected call of DeleteDevice.
func (mr *MockIDeviceRepoMockRecorder) DeleteDevice(idDevice, name, home interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDevice", reflect.TypeOf((*MockIDeviceRepo)(nil).DeleteDevice), idDevice, name, home)
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

// GetListDevices mocks base method.
func (m *MockIDeviceRepo) GetListDevices(userID int) ([]pkg.Devices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListDevices", userID)
	ret0, _ := ret[0].([]pkg.Devices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListDevices indicates an expected call of GetListDevices.
func (mr *MockIDeviceRepoMockRecorder) GetListDevices(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListDevices", reflect.TypeOf((*MockIDeviceRepo)(nil).GetListDevices), userID)
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
func (m *MockIHistoryDeviceRepo) GetDeviceHistory(userID int, name, home string) ([]pkg.DevicesHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceHistory", userID, name, home)
	ret0, _ := ret[0].([]pkg.DevicesHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceHistory indicates an expected call of GetDeviceHistory.
func (mr *MockIHistoryDeviceRepoMockRecorder) GetDeviceHistory(userID, name, home interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceHistory", reflect.TypeOf((*MockIHistoryDeviceRepo)(nil).GetDeviceHistory), userID, name, home)
}
