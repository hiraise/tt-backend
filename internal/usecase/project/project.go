package project

import (
	"task-trail/internal/customerrors"
	"task-trail/internal/repo"
	"task-trail/internal/usecase"
)

type UseCase struct {
	txManager        repo.TxManager
	authUC           usecase.Authentication
	projectRepo      repo.ProjectRepository
	userRepo         repo.UserRepository
	notificationRepo repo.NotificationRepository
	errHandler       customerrors.ErrorHandler
}

func New(
	txManager repo.TxManager,
	authUC usecase.Authentication,
	projectRepo repo.ProjectRepository,
	userRepo repo.UserRepository,
	notificationRepo repo.NotificationRepository,
	errHandler customerrors.ErrorHandler,
) *UseCase {
	return &UseCase{

		txManager:        txManager,
		authUC:           authUC,
		projectRepo:      projectRepo,
		userRepo:         userRepo,
		notificationRepo: notificationRepo,
		errHandler:       errHandler,
	}
}
