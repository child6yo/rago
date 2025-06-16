package app

import (
	"log"

	"github.com/child6yo/rago/api-gateway/internal/app/client"
	"github.com/child6yo/rago/api-gateway/internal/app/handler"
	"github.com/child6yo/rago/api-gateway/internal/app/kafka/producer"
	"github.com/child6yo/rago/api-gateway/internal/app/server"
	"github.com/child6yo/rago/api-gateway/internal/config"
)

// Application - структура приложения splitter.
type Application struct {
	config.Config // конфигурация

	client   *client.GRPClient
	producer *producer.KafkaProducer
}

// CreateApplication создает новый экземпляр приложения.
func CreateApplication(config config.Config) *Application {
	return &Application{Config: config}
}

// StartApplication запускает приложение.
func (a *Application) StartApplication() {
	a.producer = producer.NewKafkaProducer(a.KafkaBrokers, a.KafkaTopic)
	go func() {
		if err := a.producer.StartProducer(); err != nil {
			log.Fatal(err)
		}
	}()

	a.client = client.NewGRPCClient(a.Config)
	go func() {
		a.client.StartGRPCClient()
	}()

	handler := handler.NewHandler(a.client, a.producer)

	srv := server.Server{}
	srv.Run(a.SrvPort, handler.InitRoutes())
}

// TODO
func (a *Application) StopApplication() {
	a.client.StopGRPCClient()
	a.producer.StopProducer()
}
