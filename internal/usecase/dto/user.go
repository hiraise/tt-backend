package dto

import "time"



type UserUpdate struct {
	ID           int
	Username     string
	Email        string
	AvatarID     string
	VerifiedAt   time.Time
	PasswordHash string
}

// Response DTO

type CurrentUser struct {
	ID        int
	Email     string
	Username  *string
	AvatarURL *string
}

type User struct {
	ID           int
	Email        string
	PasswordHash string
	VerifiedAt   *time.Time
	AvatarID     *string
	Username     *string
}
