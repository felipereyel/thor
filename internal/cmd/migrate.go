package cmd

import (
	"fmt"

	"github.com/felipereyel/thor/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrateDb(cfg *config.ServerConfigs) error {
	sourceURL := fmt.Sprintf("file://%s", cfg.MigrationsDir)
	databaseURL := fmt.Sprintf("sqlite://%s", cfg.DatabaseFile)

	m, err := migrate.New(sourceURL, databaseURL)
	defer m.Close()
	if err != nil {
		return fmt.Errorf("Failed to get migrate: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Failed to migrate up: %w", err)
	}

	return nil
}
