package config

import "os"

type Config struct {
	PostgresDatabase    string
	AuthSecret          string
	DisableRegistration string
	BaseURL             string
}

var CurrentConfig Config

func InitializeConfig() {
	CurrentConfig = Config{
		PostgresDatabase:    os.Getenv("POSTGRES_DATABASE"),
		AuthSecret:          os.Getenv("AUTH_SECRET"),
		DisableRegistration: os.Getenv("DISABLE_REGISTRATION"),
		BaseURL:             os.Getenv("BASE_URL"),
	}
}
