package auth

import (
	"context"
	"errors"
	"task-trail/internal/entity"
	"task-trail/internal/repo"
	"time"
)

func (u *UseCase) Verify(ctx context.Context, tokenID string) error {

	f := func(ctx context.Context) error {
		now := time.Now()
		token, err := u.etRepo.GetByID(ctx, tokenID)
		if err != nil {
			if errors.Is(err, repo.ErrNotFound) {
				return u.errHandler.NotFound(err, "email token not found", "token", tokenID)
			}
			return u.errHandler.InternalTrouble(err, "email token loading failed", "token", tokenID)
		}
		if token.UsedAt != nil {
			return u.errHandler.BadRequest(err, "email token already used", "token", tokenID)
		}
		if token.ExpiredAt.Unix() <= now.Unix() {
			return u.errHandler.BadRequest(err, "email token is expired", "token", tokenID)
		}
		if err := u.userRepo.Update(ctx, &entity.User{ID: token.UserID, VerifiedAt: &now}); err != nil {
			if errors.Is(err, repo.ErrNotFound) {
				return u.errHandler.NotFound(err, "user not found", "userID", token.UserID)
			}
			return u.errHandler.InternalTrouble(err, "user verification failed", "userID", token.UserID)
		}
		if err := u.etRepo.Use(ctx, tokenID); err != nil {
			if errors.Is(err, repo.ErrNotFound) {
				return u.errHandler.NotFound(err, "email token not found", "token", token.ID)
			}
			return u.errHandler.InternalTrouble(err, "email token update failed", "token", token.ID)
		}
		return nil
	}

	return u.txManager.DoWithTx(ctx, f)

}
