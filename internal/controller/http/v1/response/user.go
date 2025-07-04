package response

import (
	"task-trail/internal/usecase/dto"
)

type AvatarRes struct {
	AvatarUrl string `json:"avatarUrl"`
}

func NewAvatarResFromDTO(data *dto.UserAvatar) *AvatarRes {
	return &AvatarRes{AvatarUrl: data.AvatarURL}
}

type CurrentRes struct {
	ID        int     `json:"id"`
	Email     string  `json:"email"`
	Username  *string `json:"username"`
	AvatarUrl *string `json:"avatarUrl"`
}

func NewCurrentResFromDTO(data *dto.CurrentUser) *CurrentRes {
	return &CurrentRes{
		ID:        data.ID,
		Username:  data.Username,
		Email:     data.Email,
		AvatarUrl: data.AvatarURL,
	}
}
