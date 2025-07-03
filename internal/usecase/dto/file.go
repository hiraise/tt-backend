package dto

type UploadFile struct {
	UserID int
	File   *File
}

type File struct {
	Data     []byte
	Name     string
	MimeType string
}
