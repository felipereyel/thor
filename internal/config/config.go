package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type ServerConfigs struct {
	ServerAddress string

	AdminSecret string
}

func GetServerConfigs() ServerConfigs {
	config := ServerConfigs{}

	// mandatory

	config.AdminSecret = os.Getenv("ADMIN_SECRET")
	if config.AdminSecret == "" {
		panic("Missing ADMIN_SECRET")
	}

	// optional - with defaults

	envPort := os.Getenv("PORT")
	if envPort != "" {
		config.ServerAddress = ":" + envPort
	} else {
		config.ServerAddress = ":3000"
	}

	return config
}
