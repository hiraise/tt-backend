package auth

import (
	"context"
	"task-trail/internal/usecase/dto"
)

func (u *UseCase) SendPasswordResetEmail(ctx context.Context, email string) error {
	f := func(ctx context.Context) error {
		user, err := u.getUserByEmail(ctx, email)
		if err != nil {
			return err
		}
		if user.VerifiedAt == nil {
			return u.errHandler.BadRequest(nil, "user is not verified", "userID", user.ID)
		}
		// create email token
		tokenID, err := u.createEmailToken(ctx, user.ID, dto.PurposeVerification)
		if err != nil {
			return err
		}
		// send email
		if err := u.notificationRepo.SendResetPasswordEmail(ctx, email, tokenID); err != nil {
			return u.errHandler.InternalTrouble(err, "failed to send reset password email", "userID", user.ID)
		}
		return nil
	}
	return u.txManager.DoWithTx(ctx, f)
}
