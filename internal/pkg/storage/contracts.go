package storage

import "context"

type Service interface {
	Save(ctx context.Context, file []byte, name string, mimeType string) error
	Delete(ctx context.Context, name string) error
	GetPath(name string) string
}
