package auth

import (
	"context"
	"errors"
	"task-trail/internal/entity"
	"task-trail/internal/repo"
	"time"
)

func (u *UseCase) Register(ctx context.Context, email string, password string) error {

	f := func(ctx context.Context) error {
		hash, err := u.passwordSvc.HashPassword(password)
		if err != nil {
			return u.errHandler.InternalTrouble(err, "password hashing failed")
		}

		user := &entity.User{Email: email, PasswordHash: string(hash)}
		id, err := u.userRepo.Create(ctx, user)
		if err != nil {
			if errors.Is(err, repo.ErrConflict) {
				return u.errHandler.Conflict(err, "email already taken", "email", email)
			}
			return u.errHandler.InternalTrouble(err, "failed to create new user", "email", email)
		}

		et := entity.EmailToken{
			ID:        u.uuid.Generate(),
			ExpiredAt: time.Now().Add(time.Minute * 10),
			UserId:    id,
			Purpose:   entity.PurposeConfirmation,
		}
		u.etRepo.Create(ctx, et)
		u.notificationRepo.SendConfirmationEmail(ctx, user.Email, et.ID)
		return nil
	}

	return u.txManager.DoWithTx(ctx, f)
}
