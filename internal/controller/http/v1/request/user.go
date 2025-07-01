package request

import (
	"task-trail/internal/entity"
	"task-trail/internal/usecase/dto"
)

// AuthReq represents the request payload for user authentication.
// It contains the user's email and password, both of which are required.
type AuthReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

// VerifyReq represents the request payload for verifying a user using a verify token.
type VerifyReq struct {
	Token string `json:"token" binding:"required,uuid"`
}

// EmailReq represents a request payload containing an email address.
type EmailReq struct {
	Email string `json:"email" binding:"required,email"`
}

// ResetPasswordReq represents the request payload for resetting a user's password.
type ResetPasswordReq struct {
	Token    string `json:"token" binding:"required,uuid"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type ChangePasswordReq struct {
	OldPassword string `json:"oldPassword" binding:"required,min=8,max=50"`
	NewPassword string `json:"newPassword" binding:"required,min=8,max=50"`
}

func (r *ChangePasswordReq) ToDTO(userID int) dto.ChangePasswordDTO {
	return dto.ChangePasswordDTO{UserID: userID, OldPassword: r.OldPassword, NewPassword: r.NewPassword}
}

// UpdateReq represents the request payload for updating a user's information.
type UpdateReq struct {
	Username string `json:"username" binding:"max=100"`
}

// FormAPI converts the UpdateReq struct into an entity.User instance
func (u *UpdateReq) ToEntity() *entity.User {
	return &entity.User{Username: &u.Username}
}
