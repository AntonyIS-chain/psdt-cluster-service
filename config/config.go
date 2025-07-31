package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	CLUSTERS_TABLE      string
	TRIBES_TABLE        string
	SQUADS_TABLE        string
	USERS_TABLE         string
	SQUAD_USERS_TABLE   string
	SERVER_PORT         string
	POSTGRES_DB         string
	POSTGRES_USER       string
	POSTGRES_HOST       string
	POSTGRES_PORT       string
	POSTGRES_PASSWORD   string
	LOGGING_SERVICE_URL string
	DEBUG               bool
	TEST                bool
	ENV                 string
}

func NewConfig() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// Load environment file for development
	if env == "development" {
		if err := godotenv.Load(".env.development"); err != nil {
			return nil, err
		}
	}

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	test, _ := strconv.ParseBool(os.Getenv("TEST"))

	return &Config{
		CLUSTERS_TABLE:      os.Getenv("CLUSTERS_TABLE"),
		TRIBES_TABLE:        os.Getenv("TRIBES_TABLE"),
		SQUADS_TABLE:        os.Getenv("SQUADS_TABLE"),
		USERS_TABLE:         os.Getenv("USERS_TABLE"),
		SQUAD_USERS_TABLE:   os.Getenv("SQUAD_USERS_TABLE"),
		SERVER_PORT:         os.Getenv("SERVER_PORT"),
		POSTGRES_DB:         os.Getenv("POSTGRES_DB"),
		POSTGRES_USER:       os.Getenv("POSTGRES_USER"),
		POSTGRES_HOST:       os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:       os.Getenv("POSTGRES_PORT"),
		POSTGRES_PASSWORD:   os.Getenv("POSTGRES_PASSWORD"),
		LOGGING_SERVICE_URL: os.Getenv("LOGGING_SERVICE_URL"),
		DEBUG:               debug,
		TEST:                test,
		ENV:                 env,
	}, nil
}
