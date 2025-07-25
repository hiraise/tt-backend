// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecase/contracts.go
//
// Generated by this command:
//
//	mockgen -source=internal/usecase/contracts.go -destination=test/mocks/mock_usecase.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	dto "task-trail/internal/usecase/dto"

	gomock "go.uber.org/mock/gomock"
)

// MockAuthentication is a mock of Authentication interface.
type MockAuthentication struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationMockRecorder
	isgomock struct{}
}

// MockAuthenticationMockRecorder is the mock recorder for MockAuthentication.
type MockAuthenticationMockRecorder struct {
	mock *MockAuthentication
}

// NewMockAuthentication creates a new mock instance.
func NewMockAuthentication(ctrl *gomock.Controller) *MockAuthentication {
	mock := &MockAuthentication{ctrl: ctrl}
	mock.recorder = &MockAuthenticationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthentication) EXPECT() *MockAuthenticationMockRecorder {
	return m.recorder
}

// AutoRegister mocks base method.
func (m *MockAuthentication) AutoRegister(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoRegister", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoRegister indicates an expected call of AutoRegister.
func (mr *MockAuthenticationMockRecorder) AutoRegister(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoRegister", reflect.TypeOf((*MockAuthentication)(nil).AutoRegister), ctx, email)
}

// ChangePassword mocks base method.
func (m *MockAuthentication) ChangePassword(ctx context.Context, data *dto.PasswordChange) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockAuthenticationMockRecorder) ChangePassword(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockAuthentication)(nil).ChangePassword), ctx, data)
}

// Login mocks base method.
func (m *MockAuthentication) Login(ctx context.Context, data *dto.Credentials) (*dto.LoginRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, data)
	ret0, _ := ret[0].(*dto.LoginRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthenticationMockRecorder) Login(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthentication)(nil).Login), ctx, data)
}

// Logout mocks base method.
func (m *MockAuthentication) Logout(ctx context.Context, refreshToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", ctx, refreshToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockAuthenticationMockRecorder) Logout(ctx, refreshToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthentication)(nil).Logout), ctx, refreshToken)
}

// Refresh mocks base method.
func (m *MockAuthentication) Refresh(ctx context.Context, refreshToken string) (*dto.RefreshRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", ctx, refreshToken)
	ret0, _ := ret[0].(*dto.RefreshRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh.
func (mr *MockAuthenticationMockRecorder) Refresh(ctx, refreshToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockAuthentication)(nil).Refresh), ctx, refreshToken)
}

// Register mocks base method.
func (m *MockAuthentication) Register(ctx context.Context, data *dto.Credentials) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockAuthenticationMockRecorder) Register(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthentication)(nil).Register), ctx, data)
}

// ResendVerificationEmail mocks base method.
func (m *MockAuthentication) ResendVerificationEmail(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResendVerificationEmail", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResendVerificationEmail indicates an expected call of ResendVerificationEmail.
func (mr *MockAuthenticationMockRecorder) ResendVerificationEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResendVerificationEmail", reflect.TypeOf((*MockAuthentication)(nil).ResendVerificationEmail), ctx, email)
}

// ResetPassword mocks base method.
func (m *MockAuthentication) ResetPassword(ctx context.Context, data *dto.PasswordReset) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockAuthenticationMockRecorder) ResetPassword(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockAuthentication)(nil).ResetPassword), ctx, data)
}

// SendPasswordResetEmail mocks base method.
func (m *MockAuthentication) SendPasswordResetEmail(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendPasswordResetEmail", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendPasswordResetEmail indicates an expected call of SendPasswordResetEmail.
func (mr *MockAuthenticationMockRecorder) SendPasswordResetEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPasswordResetEmail", reflect.TypeOf((*MockAuthentication)(nil).SendPasswordResetEmail), ctx, email)
}

// Verify mocks base method.
func (m *MockAuthentication) Verify(ctx context.Context, tokenID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, tokenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockAuthenticationMockRecorder) Verify(ctx, tokenID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockAuthentication)(nil).Verify), ctx, tokenID)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
	isgomock struct{}
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// GetCurrentByID mocks base method.
func (m *MockUser) GetCurrentByID(ctx context.Context, ID int) (*dto.CurrentUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentByID", ctx, ID)
	ret0, _ := ret[0].(*dto.CurrentUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentByID indicates an expected call of GetCurrentByID.
func (mr *MockUserMockRecorder) GetCurrentByID(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentByID", reflect.TypeOf((*MockUser)(nil).GetCurrentByID), ctx, ID)
}

// UpdateAvatar mocks base method.
func (m *MockUser) UpdateAvatar(ctx context.Context, data *dto.FileUpload) (*dto.UserAvatar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", ctx, data)
	ret0, _ := ret[0].(*dto.UserAvatar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockUserMockRecorder) UpdateAvatar(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockUser)(nil).UpdateAvatar), ctx, data)
}

// UpdateByID mocks base method.
func (m *MockUser) UpdateByID(ctx context.Context, data *dto.UserUpdate) (*dto.CurrentUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", ctx, data)
	ret0, _ := ret[0].(*dto.CurrentUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockUserMockRecorder) UpdateByID(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockUser)(nil).UpdateByID), ctx, data)
}

// MockFile is a mock of File interface.
type MockFile struct {
	ctrl     *gomock.Controller
	recorder *MockFileMockRecorder
	isgomock struct{}
}

// MockFileMockRecorder is the mock recorder for MockFile.
type MockFileMockRecorder struct {
	mock *MockFile
}

// NewMockFile creates a new mock instance.
func NewMockFile(ctrl *gomock.Controller) *MockFile {
	mock := &MockFile{ctrl: ctrl}
	mock.recorder = &MockFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFile) EXPECT() *MockFileMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockFile) Save(ctx context.Context, data *dto.FileUpload) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockFileMockRecorder) Save(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockFile)(nil).Save), ctx, data)
}

// MockProject is a mock of Project interface.
type MockProject struct {
	ctrl     *gomock.Controller
	recorder *MockProjectMockRecorder
	isgomock struct{}
}

// MockProjectMockRecorder is the mock recorder for MockProject.
type MockProjectMockRecorder struct {
	mock *MockProject
}

// NewMockProject creates a new mock instance.
func NewMockProject(ctrl *gomock.Controller) *MockProject {
	mock := &MockProject{ctrl: ctrl}
	mock.recorder = &MockProjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProject) EXPECT() *MockProjectMockRecorder {
	return m.recorder
}

// AddMembers mocks base method.
func (m *MockProject) AddMembers(ctx context.Context, data *dto.ProjectAddMembers) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMembers", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMembers indicates an expected call of AddMembers.
func (mr *MockProjectMockRecorder) AddMembers(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMembers", reflect.TypeOf((*MockProject)(nil).AddMembers), ctx, data)
}

// Create mocks base method.
func (m *MockProject) Create(ctx context.Context, data *dto.ProjectCreate) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, data)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProjectMockRecorder) Create(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProject)(nil).Create), ctx, data)
}

// GetByID mocks base method.
func (m *MockProject) GetByID(ctx context.Context, projectID, memberID int) (*dto.ProjectRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, projectID, memberID)
	ret0, _ := ret[0].(*dto.ProjectRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockProjectMockRecorder) GetByID(ctx, projectID, memberID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProject)(nil).GetByID), ctx, projectID, memberID)
}

// GetCandidates mocks base method.
func (m *MockProject) GetCandidates(ctx context.Context, ownerID, projectID int) ([]*dto.UserSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCandidates", ctx, ownerID, projectID)
	ret0, _ := ret[0].([]*dto.UserSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandidates indicates an expected call of GetCandidates.
func (mr *MockProjectMockRecorder) GetCandidates(ctx, ownerID, projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandidates", reflect.TypeOf((*MockProject)(nil).GetCandidates), ctx, ownerID, projectID)
}

// GetList mocks base method.
func (m *MockProject) GetList(ctx context.Context, data *dto.ProjectList) ([]*dto.ProjectRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, data)
	ret0, _ := ret[0].([]*dto.ProjectRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockProjectMockRecorder) GetList(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockProject)(nil).GetList), ctx, data)
}
