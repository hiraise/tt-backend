package postgres

import (
	"errors"
	"task-trail/internal/pkg/logger"
	"time"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(dbUrl string, migrationPath string, l logger.Logger) error {
	var (
		attempts = 5
		timeout  = time.Second
		m        *migrate.Migrate
		err      error
	)
	for attempts > 0 {
		m, err = migrate.New(migrationPath, dbUrl+"?sslmode=disable")
		if err != nil {
			l.Info("try connect to postgresql db for migrate", "attempts", attempts)
			time.Sleep(timeout)
			attempts -= 1
			continue
		}
		break
	}
	if err != nil {
		return err
	}

	defer func() {
		serr, derr := m.Close()
		if serr != nil {
			l.Error("failed to close migration", "error", serr)
		}
		if derr != nil {
			l.Error("failed to close migration", "error", derr)
		}
	}()
	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			l.Info("migration complete without changes")
			return nil
		}
		return err
	}

	l.Info("migration complete")
	return nil
}
