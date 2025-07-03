package request

import "task-trail/internal/usecase/dto"

// UpdateReq represents the request payload for updating a user's information.
type UpdateReq struct {
	Username string `json:"username" binding:"max=100"`
}

// FormAPI converts the UpdateReq struct into an entity.User instance
func (u *UpdateReq) ToDTO(userID int) *dto.UserUpdate {
	return &dto.UserUpdate{ID: userID, Username: u.Username}
}
