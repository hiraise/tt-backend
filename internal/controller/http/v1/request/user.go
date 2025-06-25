package request

import "task-trail/internal/entity"

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type VerifyRequest struct {
	Token string `json:"token" binding:"required,uuid"`
}

type EmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token    string `json:"token" binding:"required,uuid"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type UpdateUser struct {
	Username string `json:"username" binding:"max=100"`
}

func (u *UpdateUser) FromUpdateUser() *entity.User {
	return &entity.User{Username: &u.Username}
}
