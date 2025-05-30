package app

import (
	"log"

	"github.com/child6yo/rago/services/splitter/internal/app/kafka"
	"github.com/child6yo/rago/services/splitter/internal/app/server"
	"github.com/child6yo/rago/services/splitter/internal/app/usecase"
	"github.com/child6yo/rago/services/splitter/internal/config"
)

// Application - структура приложения splitter.
type Application struct {
	config.Config // конфигурация
}

// CreateApplication создает новый экземпляр приложения.
func CreateApplication(config config.Config) *Application {
	return &Application{config}
}

// StartApplication запускает приложение.
func (a *Application) StartApplication() {
	producer := kafka.NewKafkaProducer(a.KafkaBrokers, a.KafkaDocTopic)
	go producer.StartProducer()

	splitter := usecase.NewSplitService(a.SplitterWorkers, producer)

	server := server.NewGRPCServer(splitter, a.GRPCHost, a.GRPCPort)
	err := server.StartGRPCServer()
	if err != nil {
		log.Print(err)
		// обработка
	}
}

// TODO
func (a *Application) StopApplication() {

}
