package config

import "os"

type Config struct {
	Port string

	DBAuthenticationData
}

type DBAuthenticationData struct {
	DBUsername string
	DBPassword string
	DBDatabase string
	DBHost     string
	DBPort     string
}

func New() (*Config, error) {
	return &Config{
		Port: os.Getenv("PORT"),
		DBAuthenticationData: DBAuthenticationData{
			DBUsername: os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBDatabase: os.Getenv("DB_HOST"),
			DBHost:     os.Getenv("DB_PORT"),
			DBPort:     os.Getenv("DB_BASE"),
		},
	}, nil
}
