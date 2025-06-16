package app

import (
	"log"

	"github.com/child6yo/rago/services/splitter/internal/app/kafka/consumer"
	"github.com/child6yo/rago/services/splitter/internal/app/kafka/producer"
	"github.com/child6yo/rago/services/splitter/internal/app/usecase"
	"github.com/child6yo/rago/services/splitter/internal/config"
)

// Application - структура приложения splitter.
type Application struct {
	config.Config // конфигурация

	producer *producer.KafkaProducer
	consumer *consumer.Connection
}

// CreateApplication создает новый экземпляр приложения.
func CreateApplication(config config.Config) *Application {
	return &Application{Config: config}
}

// StartApplication запускает приложение.
func (a *Application) StartApplication() {
	a.producer = producer.NewKafkaProducer(a.KafkaBrokers, a.KafkaDocTopic)
	go func() {
		if err := a.producer.StartProducer(); err != nil {
			log.Fatal(err)
		}
	}()

	splitter := usecase.NewSplitService(a.SplitterWorkers, a.producer)

	a.consumer = consumer.NewConnection(a.KafkaBrokers, a.KafkaRawDocTopic, "group.splitter", 5, splitter)
	go func() {
		if err := a.consumer.RunConsumers(); err != nil {
			log.Fatal(err)
		}
	}()
}

// StopApplication завершает работу приложения.
func (a *Application) StopApplication() {
	a.consumer.StopConsumers()
	a.producer.StopProducer()
}
