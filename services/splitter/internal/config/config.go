package config

import (
	"log"
	"os"
	"strconv"
)

// Config - структура, определяющая конфигурацию приложения.
type Config struct {
	GRPCHost string
	GRPCPort string

	KafkaBrokers   []string
	KafkaDocTopic string

	SplitterWorkers int
}

// InitConfig инициализирует конфигурацию приложения,
// переданную через переменные окружения.
func InitConfig() Config {
	grpcHost := getEnv("GRPC_HOST", "localhost")
	grpcPort := getEnv("GRPC_PORT", "5000")

	kafkaBrokers := []string{
		getEnv("KAFKA_BROKER", "localhost:9092"),
	}
	kafkaDocTopic := getEnv("KAFKA_DOC_TOPIC", "document-topic")

	splitWorkers := getIntEnv("SPLITTER_WORKERS", 10)

	return Config{grpcHost, grpcPort, kafkaBrokers, kafkaDocTopic, splitWorkers}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Print("Failed to load env")
		return defaultValue
	}
	return value
}

func getIntEnv(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Print("Failed to load env: ", err)
		return defaultValue
	}
	return value
}
