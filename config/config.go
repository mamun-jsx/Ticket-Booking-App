package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// config file blue-print
type Config struct {
	AppPort string
	DbUrl   string
}

// create func to take env value from .env file
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env file failed to load")
	}
	return &Config{
		AppPort: getEnv("PORT", "8080"),
		DbUrl:   getEnv("DatabaseURL", "localhost"),
	}

}
