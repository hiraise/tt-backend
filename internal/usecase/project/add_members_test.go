package project

import (
	"context"
	"errors"
	"task-trail/internal/customerrors"
	"task-trail/internal/repo"
	"task-trail/internal/usecase/dto"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestUseCase_AddMembers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx  context.Context
		data *dto.ProjectAddMembers
	}
	ctx := context.Background()
	testArgs :=
		args{ctx: ctx, data: &dto.ProjectAddMembers{
			ProjectID:    1,
			OwnerID:      1,
			MemberEmails: []string{"test1@mail.com", "test2@mail.com", "test3@mail.com", "test4@mail.com"},
		}}
	testProject := &dto.Project{
		ID:          1,
		OwnerID:     1,
		Name:        "Test",
		Description: "Test",
		Members:     []*dto.UserEmailAndID{{ID: 1, Email: "test@mail.com"}},
	}
	tests := []struct {
		name        string
		uc          func(ctrl *gomock.Controller, args args) *UseCase
		args        args
		wantErr     bool
		wantErrType customerrors.ErrType
		wantErrMsg  string
	}{
		{
			name: "success",
			args: testArgs,
			uc: func(ctrl *gomock.Controller, args args) *UseCase {

				uc, deps := mockUseCase(ctrl)
				// mockTx(args.ctx, deps.txManager)
				deps.projectRepo.EXPECT().GetOwnedProject(args.ctx, args.data.ProjectID, args.data.OwnerID).Return(testProject, nil)
				deps.userRepo.EXPECT().GetIdsByEmails(args.ctx, args.data.MemberEmails).Return([]*dto.UserEmailAndID{}, nil)
				return uc
			},
			wantErr: false,
		},
		{
			name: "project not found",
			args: testArgs,
			uc: func(ctrl *gomock.Controller, args args) *UseCase {

				uc, deps := mockUseCase(ctrl)
				deps.projectRepo.EXPECT().GetOwnedProject(args.ctx, args.data.ProjectID, args.data.OwnerID).Return(nil, repo.ErrNotFound)
				return uc
			},
			wantErr:     true,
			wantErrType: customerrors.NotFoundErr,
			wantErrMsg:  "project not found",
		},
		{
			name: "project loading failed",
			args: testArgs,
			uc: func(ctrl *gomock.Controller, args args) *UseCase {

				uc, deps := mockUseCase(ctrl)
				deps.projectRepo.EXPECT().GetOwnedProject(args.ctx, args.data.ProjectID, args.data.OwnerID).Return(nil, repo.ErrInternal)
				return uc
			},
			wantErr:     true,
			wantErrType: customerrors.InternalErr,
			wantErrMsg:  "project loading failed",
		},
		{
			name: "member already in project",
			args: testArgs,
			uc: func(ctrl *gomock.Controller, args args) *UseCase {

				uc, deps := mockUseCase(ctrl)
				// mockTx(args.ctx, deps.txManager)
				prj := *testProject
				prj.Members = append(prj.Members, &dto.UserEmailAndID{ID: 2, Email: "test1@mail.com"})
				deps.projectRepo.EXPECT().GetOwnedProject(args.ctx, args.data.ProjectID, args.data.OwnerID).Return(&prj, nil)
				return uc
			},
			wantErr:     true,
			wantErrType: customerrors.ValidationErr,
			wantErrMsg:  "member already in project",
		},
		{
			name: "project members loading failed",
			args: testArgs,
			uc: func(ctrl *gomock.Controller, args args) *UseCase {

				uc, deps := mockUseCase(ctrl)
				// mockTx(args.ctx, deps.txManager)
				deps.projectRepo.EXPECT().GetOwnedProject(args.ctx, args.data.ProjectID, args.data.OwnerID).Return(testProject, nil)
				deps.userRepo.EXPECT().GetIdsByEmails(args.ctx, args.data.MemberEmails).Return(nil, repo.ErrInternal)
				return uc
			},
			wantErr:     true,
			wantErrType: customerrors.InternalErr,
			wantErrMsg:  "project members loading failed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.uc(ctrl, tt.args)
			err := u.AddMembers(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				var e *customerrors.Err
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if !errors.As(err, &e) {
					t.Errorf("expected custom error type, got %T", err)
					return
				}
				if e.Type != tt.wantErrType {
					t.Errorf("unexpected error type: got %d, want %d", e.Type, tt.wantErrType)
				}
				if e.Msg != tt.wantErrMsg {
					t.Errorf("unexpected error msg: got %s, want %s", e.Msg, tt.wantErrMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}
