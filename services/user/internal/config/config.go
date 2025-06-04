package config

import (
	"log"
	"os"
)

// Config - структура, определяющая конфигурацию приложения.
type Config struct {
	PgHost     string
	PgPort     string
	PgUsername string
	PgDBName   string
	PgPassword string
	PgSSLMode  string
}

// InitConfig инициализирует конфигурацию приложения,
// переданную через переменные окружения.
func InitConfig() Config {
	cfg := Config{}

	cfg.PgHost = getEnv("PG_HOST", "localhost")
	cfg.PgPort = getEnv("PG_PORT", "5432")
	cfg.PgUsername = getEnv("PG_USERNAME", "postgres")
	cfg.PgDBName = getEnv("PG_DATABSE", "postgres")
	cfg.PgPassword = getEnv("PG_PASSWORD", "Qwerty")
	cfg.PgSSLMode = getEnv("PG_SSLMODE", "disable")

	return cfg
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Print("Failed to load env")
		return defaultValue
	}
	return value
}
