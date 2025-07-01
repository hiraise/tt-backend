package dto

type FileDTO struct {
	Data     []byte
	Name     string
	MimeType string
}

type UpdateAvatarDTO struct {
	UserID int
	File   FileDTO
}

type ChangePasswordDTO struct {
	UserID      int
	OldPassword string
	NewPassword string
}
