package dto

import "time"

//

type Credentials struct {
	Email    string
	Password string
}

type UserCreate struct {
	Email        string
	PasswordHash string
}

type PasswordChange struct {
	UserID      int
	OldPassword string
	NewPassword string
}

type PasswordReset struct {
	TokenID     string
	NewPassword string
}

// Response

type AccessToken struct {
	Token string
	Exp   time.Time
}
type RefreshToken struct {
	Token string
	Exp   time.Time
	Jti   string
}

type LoginRes struct {
	UserID int
	AT     *AccessToken
	RT     *RefreshToken
}

type RefreshRes struct {
	AT *AccessToken
	RT *RefreshToken
}
