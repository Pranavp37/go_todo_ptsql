package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	DATABASE_NAME     string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_URL      string
}

func Load() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
		return nil, err
	}

	var config *config = &config{
		DATABASE_NAME:     os.Getenv("DATABSE_NAME"),
		DATABASE_USER:     os.Getenv("DATABSE_USER"),
		DATABASE_PASSWORD: os.Getenv("DATABSE_PASSWORD"),
		DATABASE_HOST:     os.Getenv("DATABSE_HOST"),
		DATABASE_PORT:     os.Getenv("DATABSE_PORT"),
		DATABASE_URL:      os.Getenv("DATABSE_URL"),
	}
	return config, nil
}
