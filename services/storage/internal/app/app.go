package app

import (
	"context"
	"log"

	"github.com/child6yo/rago/services/storage/internal/app/kafka"
	"github.com/child6yo/rago/services/storage/internal/app/usecase"
	"github.com/child6yo/rago/services/storage/internal/config"
)

// Application - структура приложения хранилища.
type Application struct {
	config.Config // конфигурация

	context context.Context
	cancel  context.CancelFunc
}

// CreateApplication создает новый экземпляр приложения.
func CreateApplication(config config.Config) *Application {
	ctx, cancel := context.WithCancel(context.Background())
	return &Application{Config: config, context: ctx, cancel: cancel}
}

// StartApplication запускает приложение.
func (a *Application) StartApplication() {
	DocHandler := usecase.NewDocHandlerService(a.Db)

	kConn := kafka.NewKafkaConn(a.KafkaBrokers, a.KafkaDocTopic, a.KafkaGroupID, a.KafkaPartitions, DocHandler)
	if err := kConn.RunConsumers(); err != nil {
		log.Print(err)
	}
}

// TODO
// func (a *Application) StopApplication() {
// 	stopConsumer()
// }
