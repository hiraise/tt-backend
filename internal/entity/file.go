package entity

import "time"

type File struct {
	ID            string
	OriginalName  string
	MimeType      string
	OwnerID       int
	CreatedAt     time.Time
	SoftDeletedAt *time.Time
	DeletedAt     *time.Time
}
