package persistent

import (
	"context"
	"task-trail/internal/entity"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgUserRepository struct {
	PgRepostitory
}

func NewUserRepo(db *pgxpool.Pool) *PgUserRepository {
	return &PgUserRepository{PgRepostitory{pg: db}}
}

func (r *PgUserRepository) Create(ctx context.Context, user *entity.User) error {
	// TODO: remove verified_at after apply user verification
	query := `INSERT INTO users (email, password_hash, verified_at) VALUES ($1, $2, $3)`
	_, err := r.getDb(ctx).Exec(ctx, query, user.Email, user.PasswordHash, time.Now())
	if err != nil {
		return r.handleError(err)
	}
	return nil
}

func (r *PgUserRepository) EmailIsTaken(ctx context.Context, email string) (bool, error) {
	var isTaken bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`
	if err := r.getDb(ctx).QueryRow(ctx, query, email).Scan(&isTaken); err != nil {
		return false, r.handleError(err)
	}
	if isTaken {
		return true, nil
	}
	return false, nil
}

func (r *PgUserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, email, password_hash, verified_at FROM users WHERE email = $1`
	var user entity.User
	if err := r.getDb(ctx).
		QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Email, &user.PasswordHash, &user.VerifiedAt); err != nil {
		return nil, r.handleError(err)
	}
	return &user, nil
}
