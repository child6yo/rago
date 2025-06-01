package config

import (
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/child6yo/rago/services/storage/internal/pkg/database"
	"github.com/child6yo/rago/services/storage/internal/pkg/database/qdrant"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
)

// Config - структура, определяющая конфигурацию приложения.
type Config struct {
	Db database.VectorDB

	KafkaBrokers    []string
	KafkaGroupID    string
	KafkaDocTopic   []string
	KafkaPartitions int
}

// InitConfig инициализирует конфигурацию приложения,
// переданную через переменные окружения.
func InitConfig() *Config {
	cfg := &Config{}

	cfg.Db = switchDatabase()

	cfg.KafkaBrokers = []string{getEnv("KAFKA_BROKER", "localhost:9092")}
	cfg.KafkaGroupID = getEnv("KAFKA_GROUP_ID", "group.storage")
	cfg.KafkaDocTopic = []string{getEnv("KAFKA_DOC_TOPIC", "document-topic")}
	cfg.KafkaPartitions = getIntEnv("KAFKA_PARTITIONS", 5)
	return cfg
}

func switchDatabase() database.VectorDB {
	dbURL, err := url.Parse(getEnv("DB_URL", "http://localhost:6333"))
	if err != nil {
		log.Fatal(err)
	}

	embedder := connectToLLM()
	collection := getEnv("DB_COLLECTION", "dev_coll")

	switch getEnv("DB", "qdrant") {
	case "qdrant":
		db, err := qdrant.NewQdrantConnection(dbURL, embedder, collection)
		if err != nil {
			log.Fatal(err)
		}
		return db
	}

	return nil
}

func connectToLLM() embeddings.Embedder {
	ollmaURL := getEnv("OLLAMA_URL", "http://localhost:11434")
	model := getEnv("OLLAMA_EMB_MODEL", "qwen3:0.6b")
	llm, err := ollama.New(ollama.WithModel(model), ollama.WithServerURL(ollmaURL))
	if err != nil {
		log.Fatal(err)
	}

	embedder, err := embeddings.NewEmbedder(llm)
	if err != nil {
		log.Fatal(err)
	}

	return embedder
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
