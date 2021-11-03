package config

import (
	"github.com/joho/godotenv"

	"os"
)

func Config(key string) string {

	env := os.Getenv("FOO_ENV")
	if env == "" {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if env != "test" {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load() // The Original .env
	return os.Getenv(key)
}
