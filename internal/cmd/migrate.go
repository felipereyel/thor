package cmd

import (
	"fmt"

	"thor/internal/config"
	"thor/internal/embeded"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func migrateDb(cfg *config.ServerConfigs) error {
	databaseURL := fmt.Sprintf("sqlite://%s", cfg.DatabaseFile)

	d, err := iofs.New(embeded.Migrations, "migrations")
	if err != nil {
		return fmt.Errorf("failed to get embeded migrations: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, databaseURL)
	if err != nil {
		return fmt.Errorf("failed to get migrate: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to migrate up: %w", err)
	}

	return nil
}
