package persistent

import (
	"context"
	"errors"
	"task-trail/internal/entity"
	"task-trail/internal/repo"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgConn interface {
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type PgUserRepository struct {
	pg *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *PgUserRepository {
	return &PgUserRepository{pg: db}
}

func (r *PgUserRepository) getDb(ctx context.Context) pgConn {
	if tx := extractTx(ctx); tx != nil {
		return *tx
	}
	return r.pg
}

func (r *PgUserRepository) Create(ctx context.Context, user *entity.User) error {
	// TODO: remove verified_at after apply user verification
	query := `INSERT INTO users (email, password_hash, verified_at) VALUES ($1, $2, $3)`
	_, err := r.getDb(ctx).Exec(ctx, query, user.Email, user.PasswordHash, time.Now())
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return repo.Wrap(repo.ErrConflict, err)
			}
		}
		return repo.Wrap(repo.ErrInternal, err)
	}
	return nil
}

func (r *PgUserRepository) EmailIsTaken(ctx context.Context, email string) (bool, error) {
	var isTaken bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`
	if err := r.getDb(ctx).QueryRow(ctx, query, email).Scan(&isTaken); err != nil {
		return false, repo.Wrap(repo.ErrInternal, err)
	}
	if isTaken {
		return true, nil
	}
	return false, nil
}

func (r *PgUserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, email, password_hash, verified_at FROM users WHERE email = $1`
	var user entity.User
	if err := r.getDb(ctx).QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.VerifiedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {

			return nil, repo.Wrap(repo.ErrNotFound, err)
		}
		return nil, repo.Wrap(repo.ErrInternal, err)
	}
	return &user, nil
}
