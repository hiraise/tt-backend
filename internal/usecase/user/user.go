package user

import (
	"context"
	"errors"
	"task-trail/internal/customerrors"
	"task-trail/internal/entity"
	"task-trail/internal/pkg/password"
	"task-trail/internal/pkg/uuid"
	"task-trail/internal/repo"
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
	pwdService  password.Service
	errHandler  customerrors.ErrorHandler
	uuidGen     uuid.Generator
}

func New(
	txManager repo.TxManager,
	repo repo.UserRepository,
	fileUseCase *file.UseCase,
	pwdService password.Service,
	errHandler customerrors.ErrorHandler,
	uuidGen uuid.Generator,
) *UseCase {
	return &UseCase{
		txManager:   txManager,
		userRepo:    repo,
		fileUseCase: fileUseCase,
		pwdService:  pwdService,
		errHandler:  errHandler,
		uuidGen:     uuidGen,
	}
}

func (u *UseCase) UpdateAvatar(
	ctx context.Context,
	userID int,
	file []byte,
	filename string,
	mimeType string,
) (string, error) {
	if !avatarAllowedMimeTypes[mimeType] {
		return "", u.errHandler.BadRequest(nil, "invalid mime type", "mimeType", mimeType)
	}
	var avatarID string
	var err error
	fn := func(ctx context.Context) error {
		avatarID, err = u.fileUseCase.Save(ctx, userID, file, filename, mimeType)
		if err != nil {
			return err
		}
		err = u.userRepo.Update(ctx, &entity.User{ID: userID, AvatarID: &avatarID})
		if err != nil {

			if errors.Is(err, repo.ErrNotFound) {
				return u.errHandler.BadRequest(err, "user not found", "userID", userID)
			}
			return u.errHandler.InternalTrouble(err, "user update failed", "userID", userID)
		}
		return nil
	}

	if err := u.txManager.DoWithTx(ctx, fn); err != nil {
		return "", err
	}
	return avatarID, nil
}

func (u *UseCase) UpdateByID(ctx context.Context, data *entity.User) (*entity.User, error) {
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
	return user, nil
}
func (u *UseCase) GetByID(ctx context.Context, ID int) (*entity.User, error) {
	user, err := u.userRepo.GetByID(ctx, ID)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return nil, u.errHandler.BadRequest(err, "user not found", "userID", ID)
		}
		return nil, u.errHandler.InternalTrouble(err, "user update failed", "userID", ID)
	}
	return user, nil

}
