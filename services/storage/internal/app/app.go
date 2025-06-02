package app

import (
	"errors"
	"log"

	"github.com/child6yo/rago/services/storage/internal/app/kafka"
	"github.com/child6yo/rago/services/storage/internal/app/usecase"
	"github.com/child6yo/rago/services/storage/internal/config"
)

// Application - структура приложения хранилища.
type Application struct {
	config.Config // конфигурация

	broker *kafka.Connection
}

// CreateApplication создает новый экземпляр приложения.
func CreateApplication(config config.Config) *Application {
	return &Application{Config: config}
}

// StartApplication запускает приложение. Работает идемпотентно.
func (a *Application) StartApplication() error {
	// обеспечивает идемпотентность
	if a.broker != nil {
		return errors.New("application already started")
	}

	DocHandler := usecase.NewDocHandlerService(a.Db)

	kConn := kafka.NewConnection(a.KafkaBrokers, a.KafkaDocTopic, a.KafkaGroupID, a.KafkaPartitions, DocHandler)
	a.broker = kConn
	go func() {
		if err := kConn.RunConsumers(); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

// StopApplication завершает работу приложения.
func (a *Application) StopApplication() {
	a.broker.StopConsumers()
}
