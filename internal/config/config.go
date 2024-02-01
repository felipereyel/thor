package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type ServerConfigs struct {
	AdminSecret string

	ServerAddress string

	MigrationsDir string
	DataDir       string
	ConfigDir     string

	DatabaseFile string
}

func GetServerConfigs() (*ServerConfigs, error) {
	config := ServerConfigs{}

	// mandatory

	config.AdminSecret = os.Getenv("ADMIN_SECRET")
	if config.AdminSecret == "" {
		return nil, fmt.Errorf("ADMIN_SECRET is not set")
	}

	// optional - with defaults

	envPort := os.Getenv("PORT")
	if envPort != "" {
		config.ServerAddress = ":" + envPort
	} else {
		config.ServerAddress = ":3000"
	}

	envMigrationsDir := os.Getenv("MIGRATIONS_DIR")
	if envMigrationsDir != "" {
		config.MigrationsDir = envMigrationsDir
	} else {
		config.MigrationsDir = "/migrations"
	}

	envDataDir := os.Getenv("DATA_DIR")
	if envDataDir != "" {
		config.DataDir = envDataDir
	} else {
		config.DataDir = "/data"
	}

	envConfigDir := os.Getenv("CONFIG_DIR")
	if envConfigDir != "" {
		config.ConfigDir = envConfigDir
	} else {
		config.ConfigDir = "/config"
	}

	config.DatabaseFile = fmt.Sprintf("%s/%s", config.ConfigDir, "thor.sqlite3")

	return &config, nil
}
