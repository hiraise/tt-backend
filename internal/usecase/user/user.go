package user

import (
	"context"
	"errors"
	"task-trail/internal/customerrors"
	"task-trail/internal/pkg/password"
	"task-trail/internal/pkg/storage"
	"task-trail/internal/pkg/uuid"
	"task-trail/internal/repo"
	"task-trail/internal/usecase/dto"
	"task-trail/internal/usecase/file"
)

var avatarAllowedMimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
}

type UseCase struct {
	txManager   repo.TxManager
	userRepo    repo.UserRepository
	fileUseCase *file.UseCase
	storage     storage.Service
	pwdService  password.Service
	errHandler  customerrors.ErrorHandler
	uuidGen     uuid.Generator
}

func New(
	txManager repo.TxManager,
	repo repo.UserRepository,
	fileUseCase *file.UseCase,
	storage storage.Service,
	pwdService password.Service,
	errHandler customerrors.ErrorHandler,
	uuidGen uuid.Generator,
) *UseCase {
	return &UseCase{
		txManager:   txManager,
		userRepo:    repo,
		fileUseCase: fileUseCase,
		storage:     storage,
		pwdService:  pwdService,
		errHandler:  errHandler,
		uuidGen:     uuidGen,
	}
}

func (u *UseCase) UpdateAvatar(
	ctx context.Context,
	data *dto.UploadFile,
) (string, error) {
	if !avatarAllowedMimeTypes[data.File.MimeType] {
		return "", u.errHandler.BadRequest(nil, "invalid mime type", "mimeType", data.File.MimeType)
	}
	var avatarID string
	var err error
	fn := func(ctx context.Context) error {
		avatarID, err = u.fileUseCase.Save(ctx, data)
		if err != nil {
			return err
		}
		err = u.userRepo.Update(ctx, &dto.UserUpdate{ID: data.UserID, AvatarID: avatarID})
		if err != nil {

			if errors.Is(err, repo.ErrNotFound) {
				return u.errHandler.BadRequest(err, "user not found", "userID", data.UserID)
			}
			return u.errHandler.InternalTrouble(err, "user update failed", "userID", data.UserID)
		}
		return nil
	}

	if err := u.txManager.DoWithTx(ctx, fn); err != nil {
		return "", err
	}
	return avatarID, nil
}

func (u *UseCase) UpdateByID(ctx context.Context, data *dto.UserUpdate) (*dto.CurrentUser, error) {
	err := u.userRepo.Update(ctx, data)
	if err != nil {

		if errors.Is(err, repo.ErrNotFound) {
			return nil, u.errHandler.BadRequest(err, "user not found", "userID", data.ID)
		}
		return nil, u.errHandler.InternalTrouble(err, "user update failed", "userID", data.ID)
	}
	user, err := u.userRepo.GetByID(ctx, data.ID)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return nil, u.errHandler.BadRequest(err, "user not found", "userID", data.ID)
		}
		return nil, u.errHandler.InternalTrouble(err, "user update failed", "userID", data.ID)
	}
	return u.toCurrentUser(user), nil
}
func (u *UseCase) GetByID(ctx context.Context, ID int) (*dto.CurrentUser, error) {
	user, err := u.userRepo.GetByID(ctx, ID)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return nil, u.errHandler.BadRequest(err, "user not found", "userID", ID)
		}
		return nil, u.errHandler.InternalTrouble(err, "user update failed", "userID", ID)
	}
	return u.toCurrentUser(user), nil

}

func (u *UseCase) toCurrentUser(data *dto.User) *dto.CurrentUser {
	retVal := &dto.CurrentUser{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
	}
	if data.AvatarID != nil {
		avatarURL := u.storage.GetPath(*data.AvatarID)
		retVal.AvatarURL = &avatarURL
	}
	return retVal
}
