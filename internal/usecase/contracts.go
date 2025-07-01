package usecase

import (
	"context"
	"task-trail/internal/entity"
	"task-trail/internal/pkg/token"
	"task-trail/internal/usecase/dto"
)

// Authentication defines the contract for user authentication and authorization use cases.
// It provides methods for user login, registration, logout, token refresh, email verification,
// resending verification emails, sending password reset emails, and resetting passwords.
//
// Implementations of this interface should handle the necessary business logic for each operation,
// including token management and email communications.
type Authentication interface {
	Login(ctx context.Context, email string, password string) (int, *token.Token, *token.Token, error)
	Register(ctx context.Context, email string, password string) error
	Logout(ctx context.Context, rt string) error
	Refresh(ctx context.Context, rt string) (*token.Token, *token.Token, error)
	Verify(ctx context.Context, tokenID string) error
	ResendVerificationEmail(ctx context.Context, email string) error
	SendPasswordResetEmail(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, email string, password string) error
	ChangePassword(ctx context.Context, dto dto.ChangePasswordDTO) error
}

// User defines the contract for user-related operations in the application.
// It provides methods for updating a user's avatar, updating user information by ID,
// and retrieving a user by their ID.
type User interface {
	UpdateAvatar(ctx context.Context, userID int, file []byte, filename string, mimeType string) (string, error)
	UpdateByID(ctx context.Context, data *entity.User) (*entity.User, error)
	GetByID(ctx context.Context, ID int) (*entity.User, error)
}

// File defines the contract for file storage operations.
// It provides a method to save a file with associated metadata such as owner ID, filename, and MIME type.
// The Save method returns the identifier of the uploaded file, or an error if the operation fails.
type File interface {
	Save(ctx context.Context, ownerID int, file []byte, filename string, mimeType string) (string, error)
}
