package dto

import "time"

// entity
type User struct {
	ID           int
	Email        string
	PasswordHash string
	VerifiedAt   *time.Time
	AvatarID     *string
	Username     *string
}

// request

type UserUpdate struct {
	ID           int
	Username     string
	Email        string
	AvatarID     string
	VerifiedAt   time.Time
	PasswordHash string
}

// response

type CurrentUser struct {
	ID        int
	Email     string
	Username  *string
	AvatarURL *string
}

type UserAvatar struct {
	AvatarURL string
}
