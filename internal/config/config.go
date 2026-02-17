package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	DATABASE_NAME        string
	DATABASE_USER        string
	DATABASE_PASSWORD    string
	DATABASE_HOST        string
	JWT_SECRET_KEY       string
	JWT_EXPIRATION_HOURS string
	DATABASE_PORT        string
	DATABASE_URL         string
}

func Load() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
		return nil, err
	}

	var JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	if JWT_SECRET_KEY == "" {
		log.Fatal("JWT_SECRET_KEY is not set in environment variables")
	}

	var config *config = &config{
		DATABASE_NAME:        os.Getenv("DATABASE_NAME"),
		DATABASE_USER:        os.Getenv("DATABASE_USER"),
		DATABASE_PASSWORD:    os.Getenv("DATABASE_PASSWORD"),
		DATABASE_HOST:        os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:        os.Getenv("DATABASE_PORT"),
		DATABASE_URL:         os.Getenv("DATABASE_URL"),
		JWT_SECRET_KEY:       os.Getenv("JWT_SECRET_KEY"),
		JWT_EXPIRATION_HOURS: os.Getenv("JWT_EXPIRATION_HOURS"),
	}
	return config, nil
}
