package storage

import (
	"context"
	"task-trail/internal/usecase/dto"
)

type Service interface {
	Save(ctx context.Context, dto *dto.File) error
	Delete(ctx context.Context, name string) error
	GetPath(name string) string
}
