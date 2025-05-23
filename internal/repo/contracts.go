package repo

import (
	"context"
	"errors"
	"fmt"
	"task-trail/internal/entity"
)

var ErrNotFound = errors.New("entity not found")
var ErrConflict = errors.New("entity already exists")
var ErrInternal = errors.New("something went wrong")

func Wrap(err error, background error) error {
	return fmt.Errorf("%w, error: %w", err, background)
}

type TxManager interface {
	DoWithTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	EmailIsTaken(ctx context.Context, email string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}
type VerificationRepository interface {
	Create(ctx context.Context, userId int, code int) error
	Verify(ctx context.Context, code int) error
}

type TokenRepository interface {
	Create(ctx context.Context, token *entity.RefreshToken) error
	GetById(ctx context.Context, tokenId string, userId int) (*entity.RefreshToken, error)
	Revoke(ctx context.Context, tokenId string) error
	RevokeAllUsersTokens(ctx context.Context, userId int) (int, error)
	DeleteRevokedAndOldTokens(ctx context.Context, olderThan int) (int, error)
}

type EmailTokenRepository interface {
	GetById(ctx context.Context, tokenId string) (*entity.EmailToken, error)
	Create(ctx context.Context, token entity.EmailToken) error
	Use(ctx context.Context, tokenId string) error
}

type NotificationRepository interface {
	SendConfirmationEmail(ctx context.Context, email string, token string) error
	// SendResetPasswordEmail(ctx context.Context)
}
