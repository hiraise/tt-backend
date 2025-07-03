package response

import (
	"task-trail/internal/pkg/storage"
	"task-trail/internal/usecase/dto"
)

type AvatarRes struct {
	AvatarUrl string `json:"avatarUrl"`
}

func AvatarToAPI(id string, service storage.Service) *AvatarRes {
	return &AvatarRes{AvatarUrl: service.GetPath(id)}
}

type CurrentRes struct {
	ID        int     `json:"id"`
	Email     string  `json:"email"`
	Username  *string `json:"username"`
	AvatarUrl *string `json:"avatarUrl"`
}

func CurrentUserFromDTO(data *dto.CurrentUser) *CurrentRes {
	return &CurrentRes{
		ID:        data.ID,
		Username:  data.Username,
		Email:     data.Email,
		AvatarUrl: data.AvatarURL,
	}
}
