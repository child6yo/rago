package config

import (
	"log"
	"os"
	"strconv"
)

// Config - структура, определяющая конфигурацию приложения.
type Config struct {
	KafkaBrokers          []string
	KafkaRawDocTopic      []string
	KafkaRawDocPartitions int
	KafkaDocTopic         string

	SplitterWorkers int
}

// InitConfig инициализирует конфигурацию приложения,
// переданную через переменные окружения.
func InitConfig() Config {
	cfg := Config{}

	cfg.KafkaBrokers = []string{getEnv("KAFKA_BROKER", "localhost:9092")}
	cfg.KafkaRawDocTopic = []string{getEnv("KAFKA_RAW_DOC_TOPIC", "raw-docs")}
	cfg.KafkaRawDocPartitions = getIntEnv("KAFKA_RAW_DOC_PARTITIONS", 5)
	cfg.KafkaDocTopic = getEnv("KAFKA_DOC_TOPIC", "document-topic")

	cfg.SplitterWorkers = getIntEnv("SPLITTER_WORKERS", 10)

	return cfg
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		if defaultValue != "" {
			log.Printf("config: failed to load env key = %s, defaul value = %s", key, defaultValue)
		}
		return defaultValue
	}
	return value
}

func getIntEnv(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		log.Printf("config: failed to load env key = %s, defaul value = %d", key, defaultValue)
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("config: failed to load env key = %s, defaul value = %d", key, defaultValue)
		return defaultValue
	}
	return value
}
