package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type ServerConfigs struct {
	ServerAddress string
	DataDir       string

	AdminSecret string
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

	envDataDir := os.Getenv("DATA_DIR")
	if envDataDir != "" {
		config.DataDir = envDataDir
	} else {
		config.DataDir = "./data"
	}

	return &config, nil
}
