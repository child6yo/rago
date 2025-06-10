package app

import (
	"errors"
	"log"

	"github.com/child6yo/rago/services/storage/internal/app/kafka"
	"github.com/child6yo/rago/services/storage/internal/app/server"
	"github.com/child6yo/rago/services/storage/internal/app/usecase"
	"github.com/child6yo/rago/services/storage/internal/config"
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

	DocHandler := usecase.NewDocHandlerService(a.Db)

	a.broker = kafka.NewConnection(a.KafkaBrokers, a.KafkaDocTopic, a.KafkaGroupID, a.KafkaPartitions, DocHandler)
	go func() {
		if err := a.broker.RunConsumers(); err != nil {
			log.Fatal(err)
		}
	}()

	usecase := usecase.NewStorageService(a.Db)

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
