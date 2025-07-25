// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repo/contracts.go
//
// Generated by this command:
//
//	mockgen -source=internal/repo/contracts.go -destination=test/mocks/mock_repo.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	dto "task-trail/internal/usecase/dto"

	gomock "go.uber.org/mock/gomock"
)

// MockTxManager is a mock of TxManager interface.
type MockTxManager struct {
	ctrl     *gomock.Controller
	recorder *MockTxManagerMockRecorder
	isgomock struct{}
}

// MockTxManagerMockRecorder is the mock recorder for MockTxManager.
type MockTxManagerMockRecorder struct {
	mock *MockTxManager
}

// NewMockTxManager creates a new mock instance.
func NewMockTxManager(ctrl *gomock.Controller) *MockTxManager {
	mock := &MockTxManager{ctrl: ctrl}
	mock.recorder = &MockTxManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxManager) EXPECT() *MockTxManagerMockRecorder {
	return m.recorder
}

// DoWithTx mocks base method.
func (m *MockTxManager) DoWithTx(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoWithTx", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoWithTx indicates an expected call of DoWithTx.
func (mr *MockTxManagerMockRecorder) DoWithTx(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoWithTx", reflect.TypeOf((*MockTxManager)(nil).DoWithTx), ctx, fn)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
	isgomock struct{}
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
func (m *MockUserRepository) Create(ctx context.Context, arg1 *dto.UserCreate) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(ctx, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), ctx, arg1)
}

// EmailIsTaken mocks base method.
func (m *MockUserRepository) EmailIsTaken(ctx context.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmailIsTaken", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EmailIsTaken indicates an expected call of EmailIsTaken.
func (mr *MockUserRepositoryMockRecorder) EmailIsTaken(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmailIsTaken", reflect.TypeOf((*MockUserRepository)(nil).EmailIsTaken), ctx, email)
}

// GetByEmail mocks base method.
func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(*dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUserRepositoryMockRecorder) GetByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetByEmail), ctx, email)
}

// GetByID mocks base method.
func (m *MockUserRepository) GetByID(ctx context.Context, ID int) (*dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, ID)
	ret0, _ := ret[0].(*dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserRepositoryMockRecorder) GetByID(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserRepository)(nil).GetByID), ctx, ID)
}

// GetIdsByEmails mocks base method.
func (m *MockUserRepository) GetIdsByEmails(ctx context.Context, emails []string) ([]*dto.UserEmailAndID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdsByEmails", ctx, emails)
	ret0, _ := ret[0].([]*dto.UserEmailAndID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdsByEmails indicates an expected call of GetIdsByEmails.
func (mr *MockUserRepositoryMockRecorder) GetIdsByEmails(ctx, emails any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdsByEmails", reflect.TypeOf((*MockUserRepository)(nil).GetIdsByEmails), ctx, emails)
}

// Update mocks base method.
func (m *MockUserRepository) Update(ctx context.Context, arg1 *dto.UserUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(ctx, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), ctx, arg1)
}

// MockVerificationRepository is a mock of VerificationRepository interface.
type MockVerificationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVerificationRepositoryMockRecorder
	isgomock struct{}
}

// MockVerificationRepositoryMockRecorder is the mock recorder for MockVerificationRepository.
type MockVerificationRepositoryMockRecorder struct {
	mock *MockVerificationRepository
}

// NewMockVerificationRepository creates a new mock instance.
func NewMockVerificationRepository(ctrl *gomock.Controller) *MockVerificationRepository {
	mock := &MockVerificationRepository{ctrl: ctrl}
	mock.recorder = &MockVerificationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerificationRepository) EXPECT() *MockVerificationRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVerificationRepository) Create(ctx context.Context, userID, code int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userID, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockVerificationRepositoryMockRecorder) Create(ctx, userID, code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVerificationRepository)(nil).Create), ctx, userID, code)
}

// Verify mocks base method.
func (m *MockVerificationRepository) Verify(ctx context.Context, code int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockVerificationRepositoryMockRecorder) Verify(ctx, code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockVerificationRepository)(nil).Verify), ctx, code)
}

// MockRefreshTokenRepository is a mock of RefreshTokenRepository interface.
type MockRefreshTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenRepositoryMockRecorder
	isgomock struct{}
}

// MockRefreshTokenRepositoryMockRecorder is the mock recorder for MockRefreshTokenRepository.
type MockRefreshTokenRepositoryMockRecorder struct {
	mock *MockRefreshTokenRepository
}

// NewMockRefreshTokenRepository creates a new mock instance.
func NewMockRefreshTokenRepository(ctrl *gomock.Controller) *MockRefreshTokenRepository {
	mock := &MockRefreshTokenRepository{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshTokenRepository) EXPECT() *MockRefreshTokenRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRefreshTokenRepository) Create(ctx context.Context, data *dto.RefreshTokenCreate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRefreshTokenRepositoryMockRecorder) Create(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRefreshTokenRepository)(nil).Create), ctx, data)
}

// DeleteRevokedAndOldTokens mocks base method.
func (m *MockRefreshTokenRepository) DeleteRevokedAndOldTokens(ctx context.Context, olderThan int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRevokedAndOldTokens", ctx, olderThan)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRevokedAndOldTokens indicates an expected call of DeleteRevokedAndOldTokens.
func (mr *MockRefreshTokenRepositoryMockRecorder) DeleteRevokedAndOldTokens(ctx, olderThan any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRevokedAndOldTokens", reflect.TypeOf((*MockRefreshTokenRepository)(nil).DeleteRevokedAndOldTokens), ctx, olderThan)
}

// GetByID mocks base method.
func (m *MockRefreshTokenRepository) GetByID(ctx context.Context, tokenID string, userID int) (*dto.RefreshToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, tokenID, userID)
	ret0, _ := ret[0].(*dto.RefreshToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRefreshTokenRepositoryMockRecorder) GetByID(ctx, tokenID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRefreshTokenRepository)(nil).GetByID), ctx, tokenID, userID)
}

// Revoke mocks base method.
func (m *MockRefreshTokenRepository) Revoke(ctx context.Context, tokenID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revoke", ctx, tokenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Revoke indicates an expected call of Revoke.
func (mr *MockRefreshTokenRepositoryMockRecorder) Revoke(ctx, tokenID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockRefreshTokenRepository)(nil).Revoke), ctx, tokenID)
}

// RevokeAllUsersTokens mocks base method.
func (m *MockRefreshTokenRepository) RevokeAllUsersTokens(ctx context.Context, userID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAllUsersTokens", ctx, userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAllUsersTokens indicates an expected call of RevokeAllUsersTokens.
func (mr *MockRefreshTokenRepositoryMockRecorder) RevokeAllUsersTokens(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAllUsersTokens", reflect.TypeOf((*MockRefreshTokenRepository)(nil).RevokeAllUsersTokens), ctx, userID)
}

// MockEmailTokenRepository is a mock of EmailTokenRepository interface.
type MockEmailTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEmailTokenRepositoryMockRecorder
	isgomock struct{}
}

// MockEmailTokenRepositoryMockRecorder is the mock recorder for MockEmailTokenRepository.
type MockEmailTokenRepositoryMockRecorder struct {
	mock *MockEmailTokenRepository
}

// NewMockEmailTokenRepository creates a new mock instance.
func NewMockEmailTokenRepository(ctrl *gomock.Controller) *MockEmailTokenRepository {
	mock := &MockEmailTokenRepository{ctrl: ctrl}
	mock.recorder = &MockEmailTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailTokenRepository) EXPECT() *MockEmailTokenRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEmailTokenRepository) Create(ctx context.Context, data *dto.EmailTokenCreate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockEmailTokenRepositoryMockRecorder) Create(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEmailTokenRepository)(nil).Create), ctx, data)
}

// DeleteUsedAndOldTokens mocks base method.
func (m *MockEmailTokenRepository) DeleteUsedAndOldTokens(ctx context.Context, olderThan int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsedAndOldTokens", ctx, olderThan)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUsedAndOldTokens indicates an expected call of DeleteUsedAndOldTokens.
func (mr *MockEmailTokenRepositoryMockRecorder) DeleteUsedAndOldTokens(ctx, olderThan any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsedAndOldTokens", reflect.TypeOf((*MockEmailTokenRepository)(nil).DeleteUsedAndOldTokens), ctx, olderThan)
}

// GetByID mocks base method.
func (m *MockEmailTokenRepository) GetByID(ctx context.Context, tokenID string) (*dto.EmailToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, tokenID)
	ret0, _ := ret[0].(*dto.EmailToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockEmailTokenRepositoryMockRecorder) GetByID(ctx, tokenID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockEmailTokenRepository)(nil).GetByID), ctx, tokenID)
}

// Use mocks base method.
func (m *MockEmailTokenRepository) Use(ctx context.Context, tokenID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Use", ctx, tokenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Use indicates an expected call of Use.
func (mr *MockEmailTokenRepositoryMockRecorder) Use(ctx, tokenID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockEmailTokenRepository)(nil).Use), ctx, tokenID)
}

// MockNotificationRepository is a mock of NotificationRepository interface.
type MockNotificationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationRepositoryMockRecorder
	isgomock struct{}
}

// MockNotificationRepositoryMockRecorder is the mock recorder for MockNotificationRepository.
type MockNotificationRepositoryMockRecorder struct {
	mock *MockNotificationRepository
}

// NewMockNotificationRepository creates a new mock instance.
func NewMockNotificationRepository(ctrl *gomock.Controller) *MockNotificationRepository {
	mock := &MockNotificationRepository{ctrl: ctrl}
	mock.recorder = &MockNotificationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationRepository) EXPECT() *MockNotificationRepositoryMockRecorder {
	return m.recorder
}

// SendAutoRegisterEmail mocks base method.
func (m *MockNotificationRepository) SendAutoRegisterEmail(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAutoRegisterEmail", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAutoRegisterEmail indicates an expected call of SendAutoRegisterEmail.
func (mr *MockNotificationRepositoryMockRecorder) SendAutoRegisterEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAutoRegisterEmail", reflect.TypeOf((*MockNotificationRepository)(nil).SendAutoRegisterEmail), ctx, email)
}

// SendInvintationInProject mocks base method.
func (m *MockNotificationRepository) SendInvintationInProject(ctx context.Context, data *dto.NotificationProjectInvite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendInvintationInProject", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendInvintationInProject indicates an expected call of SendInvintationInProject.
func (mr *MockNotificationRepositoryMockRecorder) SendInvintationInProject(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendInvintationInProject", reflect.TypeOf((*MockNotificationRepository)(nil).SendInvintationInProject), ctx, data)
}

// SendResetPasswordEmail mocks base method.
func (m *MockNotificationRepository) SendResetPasswordEmail(ctx context.Context, email, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendResetPasswordEmail", ctx, email, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendResetPasswordEmail indicates an expected call of SendResetPasswordEmail.
func (mr *MockNotificationRepositoryMockRecorder) SendResetPasswordEmail(ctx, email, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendResetPasswordEmail", reflect.TypeOf((*MockNotificationRepository)(nil).SendResetPasswordEmail), ctx, email, token)
}

// SendVerificationEmail mocks base method.
func (m *MockNotificationRepository) SendVerificationEmail(ctx context.Context, email, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerificationEmail", ctx, email, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerificationEmail indicates an expected call of SendVerificationEmail.
func (mr *MockNotificationRepositoryMockRecorder) SendVerificationEmail(ctx, email, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerificationEmail", reflect.TypeOf((*MockNotificationRepository)(nil).SendVerificationEmail), ctx, email, token)
}

// MockFileRepository is a mock of FileRepository interface.
type MockFileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFileRepositoryMockRecorder
	isgomock struct{}
}

// MockFileRepositoryMockRecorder is the mock recorder for MockFileRepository.
type MockFileRepositoryMockRecorder struct {
	mock *MockFileRepository
}

// NewMockFileRepository creates a new mock instance.
func NewMockFileRepository(ctrl *gomock.Controller) *MockFileRepository {
	mock := &MockFileRepository{ctrl: ctrl}
	mock.recorder = &MockFileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileRepository) EXPECT() *MockFileRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFileRepository) Create(ctx context.Context, file *dto.FileCreate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, file)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockFileRepositoryMockRecorder) Create(ctx, file any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFileRepository)(nil).Create), ctx, file)
}

// MockProjectRepository is a mock of ProjectRepository interface.
type MockProjectRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProjectRepositoryMockRecorder
	isgomock struct{}
}

// MockProjectRepositoryMockRecorder is the mock recorder for MockProjectRepository.
type MockProjectRepositoryMockRecorder struct {
	mock *MockProjectRepository
}

// NewMockProjectRepository creates a new mock instance.
func NewMockProjectRepository(ctrl *gomock.Controller) *MockProjectRepository {
	mock := &MockProjectRepository{ctrl: ctrl}
	mock.recorder = &MockProjectRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectRepository) EXPECT() *MockProjectRepositoryMockRecorder {
	return m.recorder
}

// AddMembers mocks base method.
func (m *MockProjectRepository) AddMembers(ctx context.Context, data *dto.ProjectAddMembersDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMembers", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMembers indicates an expected call of AddMembers.
func (mr *MockProjectRepositoryMockRecorder) AddMembers(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMembers", reflect.TypeOf((*MockProjectRepository)(nil).AddMembers), ctx, data)
}

// Create mocks base method.
func (m *MockProjectRepository) Create(ctx context.Context, data *dto.ProjectCreate) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, data)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProjectRepositoryMockRecorder) Create(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProjectRepository)(nil).Create), ctx, data)
}

// GetByID mocks base method.
func (m *MockProjectRepository) GetByID(ctx context.Context, projectID int) (*dto.ProjectRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, projectID)
	ret0, _ := ret[0].(*dto.ProjectRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockProjectRepositoryMockRecorder) GetByID(ctx, projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProjectRepository)(nil).GetByID), ctx, projectID)
}

// GetCandidates mocks base method.
func (m *MockProjectRepository) GetCandidates(ctx context.Context, ownerID, projectID int) ([]*dto.UserSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCandidates", ctx, ownerID, projectID)
	ret0, _ := ret[0].([]*dto.UserSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandidates indicates an expected call of GetCandidates.
func (mr *MockProjectRepositoryMockRecorder) GetCandidates(ctx, ownerID, projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandidates", reflect.TypeOf((*MockProjectRepository)(nil).GetCandidates), ctx, ownerID, projectID)
}

// GetList mocks base method.
func (m *MockProjectRepository) GetList(ctx context.Context, data *dto.ProjectList) ([]*dto.ProjectRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, data)
	ret0, _ := ret[0].([]*dto.ProjectRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockProjectRepositoryMockRecorder) GetList(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockProjectRepository)(nil).GetList), ctx, data)
}

// GetOwned mocks base method.
func (m *MockProjectRepository) GetOwned(ctx context.Context, projectID, ownerID int) (*dto.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwned", ctx, projectID, ownerID)
	ret0, _ := ret[0].(*dto.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOwned indicates an expected call of GetOwned.
func (mr *MockProjectRepositoryMockRecorder) GetOwned(ctx, projectID, ownerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwned", reflect.TypeOf((*MockProjectRepository)(nil).GetOwned), ctx, projectID, ownerID)
}

// IsMember mocks base method.
func (m *MockProjectRepository) IsMember(ctx context.Context, projectID, memberID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMember", ctx, projectID, memberID)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsMember indicates an expected call of IsMember.
func (mr *MockProjectRepositoryMockRecorder) IsMember(ctx, projectID, memberID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMember", reflect.TypeOf((*MockProjectRepository)(nil).IsMember), ctx, projectID, memberID)
}
