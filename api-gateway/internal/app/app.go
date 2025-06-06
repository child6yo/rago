package app

import (
	"github.com/child6yo/rago/api-gateway/internal/app/client"
	"github.com/child6yo/rago/api-gateway/internal/app/handler"
	"github.com/child6yo/rago/api-gateway/internal/app/server"
	"github.com/child6yo/rago/api-gateway/internal/config"
)

// Application - структура приложения splitter.
type Application struct {
	config.Config // конфигурация

	client *client.GRPClient
}

// CreateApplication создает новый экземпляр приложения.
func CreateApplication(config config.Config) *Application {
	return &Application{Config: config}
}

// StartApplication запускает приложение.
func (a *Application) StartApplication() {
	client := client.NewGRPCClient(a.Config)
	client.StartGRPCClient()
	
	handler := handler.NewHandler(client)

	srv := server.Server{}
	srv.Run(a.SrvPort, handler.InitRoutes())
}

// TODO
func (a *Application) StopApplication() {
	a.client.StopGRPCClient()
}
