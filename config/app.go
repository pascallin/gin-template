package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// load .env
	godotenv.Load()
}

type AppConfig struct {
	AppEnv        string
	APIServerPort string
	AppApiKey     string
	AppJwtSecret  string
	AppWebPort    string
}

func GetAppConfig() AppConfig {
	return AppConfig{
		AppEnv:        os.Getenv("APP_ENV"),
		APIServerPort: os.Getenv("WEB_PORT"),
		AppApiKey:     os.Getenv("APP_API_KEY"),
		AppJwtSecret:  os.Getenv("JWT_SECRET"),
		AppWebPort:    os.Getenv("PORT"),
	}
}
