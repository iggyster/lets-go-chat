package app

import (
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "dev"
	}

	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}

	godotenv.Load(".env." + env)
	godotenv.Load()
}
