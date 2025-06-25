package entity

import "time"

type User struct {
	ID           int
	Email        string
	PasswordHash string
	VerifiedAt   *time.Time
	AvatarID     *string
	Username     *string
}
