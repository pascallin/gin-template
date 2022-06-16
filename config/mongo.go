package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// load .env
	godotenv.Load()
}

type Mongo struct {
	URI      string `json:"uri,omitempty"`
	DATABASE string `json:"database"`
}

func GetMongoConfig() Mongo {
	return Mongo{
		URI:      os.Getenv("MONGODB_URI"),
		DATABASE: os.Getenv("MONGODB_DATABASE"),
	}
}
