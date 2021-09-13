package config

import (
	"github.com/JeremyLoy/config"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPServer string
	HTTPPort   string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPass     string
}

func Parse() (cfg Config, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	err = config.FromEnv().To(&cfg)
	if err != nil {
		return
	}

	return
}
