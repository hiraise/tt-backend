package response

import (
	"task-trail/internal/entity"
	"task-trail/internal/pkg/storage"
)

type UserAvatar struct {
	AvatarUrl string `json:"avatarUrl"`
}

func NewUserAvatar(id string, service storage.Service) *UserAvatar {
	return &UserAvatar{AvatarUrl: service.GetPath(id)}
}

type CurrentUser struct {
	ID        int     `json:"id"`
	Username  *string `json:"username"`
	Email     string  `json:"email"`
	AvatarUrl *string `json:"avatarUrl"`
}

func NewCurrentUser(u *entity.User, service storage.Service) *CurrentUser {
	var avatarUrl *string
	if u.AvatarID != nil {
		url := service.GetPath(*u.AvatarID)
		avatarUrl = &url
	}
	return &CurrentUser{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		AvatarUrl: avatarUrl,
	}
}
