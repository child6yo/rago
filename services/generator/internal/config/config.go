package config

import (
	"log"
	"os"
)

// Config - структура, определяющая конфигурацию приложения.
type Config struct {
	GRPCHost string
	GRPCPort string

	StorageGRPCHost string
	StorageGRPCPort string

	LLM       string
	OllamaURL string
}

// InitConfig инициализирует конфигурацию приложения,
// переданную через переменные окружения.
func InitConfig() Config {
	cfg := Config{}

	cfg.GRPCHost = getEnv("GRPC_HOST", "localhost")
	cfg.GRPCPort = getEnv("GRPC_PORT", "5003")

	cfg.StorageGRPCHost = getEnv("STORAGE_GRPC_HOST", "localhost")
	cfg.StorageGRPCPort = getEnv("STORAGE_GRPC_PORT", "5002")

	cfg.LLM = getEnv("LLM", "qwen3:1.7b")
	cfg.OllamaURL = getEnv("OLLAMA_URL", "http://localhost:11434")

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
