package project

import (
	"context"
	"errors"
	"task-trail/internal/customerrors"
	"task-trail/internal/repo"
	"task-trail/internal/usecase"
	"task-trail/internal/usecase/dto"
)

type UseCase struct {
	txManager   repo.TxManager
	authUC      usecase.Authentication
	projectRepo repo.ProjectRepository
	userRepo    repo.UserRepository
	errHandler  customerrors.ErrorHandler
}

func New(
	txManager repo.TxManager,
	authUC usecase.Authentication,
	projectRepo repo.ProjectRepository,
	userRepo repo.UserRepository,
	errHandler customerrors.ErrorHandler,
) *UseCase {
	return &UseCase{

		txManager:   txManager,
		authUC:      authUC,
		projectRepo: projectRepo,
		userRepo:    userRepo,
		errHandler:  errHandler,
	}
}

func (u *UseCase) Create(ctx context.Context, data *dto.ProjectCreate) (int, error) {
	id, err := u.projectRepo.Create(ctx, data)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return 0, u.errHandler.NotFound(err, "owner not found", "ownerID", data.OwnerID)
		}
		return 0, u.errHandler.InternalTrouble(err, "project creation failed", "ownerID", data.OwnerID)
	}
	return id, nil
}

func (u *UseCase) GetList(ctx context.Context, data *dto.ProjectList) ([]*dto.ProjectRes, error) {
	retVal, err := u.projectRepo.GetList(ctx, data)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return nil, u.errHandler.NotFound(err, "member not found", "memberID", data.MemberID)
		}
		return nil, u.errHandler.InternalTrouble(err, "get projects list failed", "memberID", data.MemberID)
	}
	return retVal, nil
}
