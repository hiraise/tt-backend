package response

import (
	"task-trail/internal/entity"
	"task-trail/internal/pkg/storage"
)

type AvatarRes struct {
	AvatarUrl string `json:"avatarUrl"`
}

func AvatarToAPI(id string, service storage.Service) *AvatarRes {
	return &AvatarRes{AvatarUrl: service.GetPath(id)}
}

type CurrentRes struct {
	ID        int     `json:"id"`
	Username  *string `json:"username"`
	Email     string  `json:"email"`
	AvatarUrl *string `json:"avatarUrl"`
}

func UserToAPI(u *entity.User, service storage.Service) *CurrentRes {
	var avatarUrl *string
	if u.AvatarID != nil {
		url := service.GetPath(*u.AvatarID)
		avatarUrl = &url
	}
	return &CurrentRes{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		AvatarUrl: avatarUrl,
	}
}
