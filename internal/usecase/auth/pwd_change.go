package auth

import (
	"context"
	"errors"
	"task-trail/internal/entity"
	"task-trail/internal/repo"
	"task-trail/internal/usecase/dto"
)

func (u *UseCase) ChangePassword(ctx context.Context, dto dto.ChangePasswordDTO) error {
	if dto.NewPassword == dto.OldPassword {
		return u.errHandler.BadRequest(nil, "passwords are equal", "userID", dto.UserID)
	}
	user, err := u.userRepo.GetByID(ctx, dto.UserID)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return u.errHandler.BadRequest(err, "user not found", "userID", dto.UserID)
		}
		return u.errHandler.InternalTrouble(err, "password change failed", "userID", dto.UserID)
	}
	if err := u.passwordSvc.ComparePassword(dto.OldPassword, user.PasswordHash); err != nil {
		return u.errHandler.BadRequest(err, "incorrect old password", "userID", dto.UserID)
	}
	h, err := u.passwordSvc.HashPassword(dto.NewPassword)
	if err != nil {
		return u.errHandler.InternalTrouble(err, "hashing new password failed", "userID", dto.UserID)
	}
	if err := u.userRepo.Update(ctx, &entity.User{PasswordHash: h, ID: user.ID}); err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return u.errHandler.BadRequest(err, "user not found", "userID", dto.UserID)
		}
		return u.errHandler.InternalTrouble(err, "password change failed", "userID", dto.UserID)
	}
	return nil
}
