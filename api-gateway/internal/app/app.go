package app

import (
	"context"
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

	client     *client.GRPClient
	producer   *producer.KafkaProducer
	httpServer *server.Server
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

	a.httpServer = &server.Server{}
	go func() {
		a.httpServer.Run(a.SrvPort, handler.InitRoutes())
	}()
}

// StopApplication останавливает работу приложения.
func (a *Application) StopApplication() {
	a.client.StopGRPCClient()
	a.producer.StopProducer()
	a.httpServer.Shutdown(context.Background())
}
