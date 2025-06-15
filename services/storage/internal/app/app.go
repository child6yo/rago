package app

import (
	"errors"
	"log"
	"time"

	"github.com/child6yo/rago/services/storage/internal/app/kafka"
	qdrantrepo "github.com/child6yo/rago/services/storage/internal/app/repository/qdrant"
	"github.com/child6yo/rago/services/storage/internal/app/server"
	"github.com/child6yo/rago/services/storage/internal/app/usecase"
	"github.com/child6yo/rago/services/storage/internal/config"
	"github.com/child6yo/rago/services/storage/internal/pkg/embedding"
)

// Application - структура приложения хранилища.
type Application struct {
	config.Config // конфигурация

	broker *kafka.Connection
	server *server.GRPCServer
}

// CreateApplication создает новый экземпляр приложения.
func CreateApplication(config config.Config) *Application {
	return &Application{Config: config}
}

// StartApplication запускает приложение. Работает идемпотентно.
func (a *Application) StartApplication() error {
	if a.broker != nil {
		return errors.New("application already started")
	}

	vectorDBClient, err := qdrantrepo.NewQdrantClient(a.DbHost, a.DbPort)
	if err != nil {
		log.Fatal(err)
	}

	ollamaEmbedder, err := embedding.NewOllamaEmbedder(a.OllamaModel, a.OllamaURL, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	loader := usecase.NewLoader(vectorDBClient, ollamaEmbedder)
	usecase := usecase.NewStorageService(vectorDBClient, ollamaEmbedder)

	a.broker = kafka.NewConnection(a.KafkaBrokers, a.KafkaDocTopic, a.KafkaGroupID, a.KafkaPartitions, loader)
	go func() {
		if err := a.broker.RunConsumers(); err != nil {
			log.Fatal(err)
		}
	}()

	a.server = server.NewGRPCServer(usecase, a.GRPCHost, a.GRPCPort)
	go func() {
		if err := a.server.StartGRPCServer(); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

// StopApplication завершает работу приложения.
func (a *Application) StopApplication() {
	a.broker.StopConsumers()
	a.server.ShutdownGRPCServer()
}
