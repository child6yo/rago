package config

import (
	"log"
	"os"
)

// Config - структура, определяющая конфигурацию приложения.
type Config struct {
	SrvHost string
	SrvPort string

	KafkaBrokers []string
	KafkaTopic   string

	UserGRPCHost string
	UserGRPCPort string

	SplitterGRPCHost string
	SplitterGRPCPort string

	GeneratorGRPCHost string
	GeneratorGRPCPort string
}

func InitConfig() Config {
	cfg := Config{}

	cfg.SrvHost = getEnv("SERVER_HOST", "localhost")
	cfg.SrvPort = getEnv("SERVER_PORT", "8080")

	cfg.KafkaBrokers = []string{getEnv("KAFKA_BROKER", "localhost:9092")}
	cfg.KafkaTopic = getEnv("KAFKA_RAW_DOC_TOPIC", "raw-docs")

	cfg.UserGRPCHost = getEnv("GRPC_HOST", "localhost")
	cfg.UserGRPCPort = getEnv("GRPC_PORT", "5001")

	cfg.SplitterGRPCHost = getEnv("GRPC_HOST", "localhost")
	cfg.SplitterGRPCPort = getEnv("GRPC_PORT", "5000")

	cfg.GeneratorGRPCHost = getEnv("GRPC_HOST", "localhost")
	cfg.GeneratorGRPCPort = getEnv("GRPC_PORT", "5003")

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
