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

	DbHost string
	DbPort int

	OllamaModel string
	OllamaURL   string

	KafkaBrokers    []string
	KafkaGroupID    string
	KafkaDocTopic   []string
	KafkaPartitions int
}

// InitConfig инициализирует конфигурацию приложения,
// переданную через переменные окружения.
func InitConfig() Config {
	cfg := Config{}

	cfg.GRPCHost = getEnv("GRPC_HOST", "localhost")
	cfg.GRPCPort = getEnv("GRPC_PORT", "5002")

	cfg.DbHost = getEnv("VECTORDB_HOST", "localhost")
	cfg.DbPort = getIntEnv("VECTORDB_PORT", 6333)

	cfg.OllamaModel = getEnv("OLLAMA_MODEL", "nomic-embed-text:v1.5")
	cfg.OllamaURL = getEnv("OLLAMA_ADDRES", "localhost:11434")

	cfg.KafkaBrokers = []string{getEnv("KAFKA_BROKER", "localhost:9092")}
	cfg.KafkaGroupID = getEnv("KAFKA_GROUP_ID", "group.storage")
	cfg.KafkaDocTopic = []string{getEnv("KAFKA_DOC_TOPIC", "document-topic")}
	cfg.KafkaPartitions = getIntEnv("KAFKA_PARTITIONS", 5)
	return cfg
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("config: failed to load env, defaul value = %s", defaultValue)
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
		log.Printf("config: failed to load env, defaul value = %d", defaultValue)
		return defaultValue
	}
	return value
}
