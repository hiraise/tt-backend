package auth

import "context"

func (u *UseCase) Logout(ctx context.Context, rt string) error {
	userId, tokenId, err := u.verifyRT(ctx, rt)
	if err != nil {
		return err
	}
	return u.revokeRT(ctx, tokenId, userId)
}
