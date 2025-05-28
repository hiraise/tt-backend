package auth

import (
	"context"
	"errors"
	"task-trail/internal/entity"
	"task-trail/internal/repo"
)

func (u *UseCase) Resend(ctx context.Context, email string) error {

	f := func(ctx context.Context) error {
		user, err := u.userRepo.GetByEmail(ctx, email)
		if err != nil {
			if errors.Is(err, repo.ErrNotFound) {
				return u.errHandler.Ok(err, "user not found", "email", email)
			}
			return u.errHandler.InternalTrouble(err, "failed to get user", "email", email)
		}
		// create email token
		tokenID, err := u.createEmailToken(ctx, user.ID, entity.PurposeVerification)
		if err != nil {
			return err
		}
		// send verification
		if err := u.notificationRepo.SendVerificationEmail(ctx, email, tokenID); err != nil {
			return u.errHandler.InternalTrouble(err, "verification email sending failed", "userID", user.ID)
		}
		return nil
	}

	return u.txManager.DoWithTx(ctx, f)
}
