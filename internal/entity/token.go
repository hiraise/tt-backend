package entity

import "time"

type EmailTokenPurpose string

const (
	PurposeConfirmation EmailTokenPurpose = "confirm"
	PurposeReset        EmailTokenPurpose = "reset"
)

type RefreshToken struct {
	ID        string
	UserId    int
	CreatedAt time.Time
	ExpiredAt time.Time
	RevokedAt *time.Time
}

type EmailToken struct {
	ID        string
	UserId    int
	Purpose   EmailTokenPurpose
	CreatedAt time.Time
	ExpiredAt time.Time
	UsedAt    *time.Time
}
