package app

import (
	"log"

	"github.com/child6yo/rago/services/generator/internal/app/client"
	"github.com/child6yo/rago/services/generator/internal/app/server"
	"github.com/child6yo/rago/services/generator/internal/app/usecase"
	"github.com/child6yo/rago/services/generator/internal/config"
)

// Application - структура приложения splitter.
type Application struct {
	config.Config // конфигурация
}

// CreateApplication создает новый экземпляр приложения.
func CreateApplication(config config.Config) *Application {
	return &Application{Config: config}
}

// StartApplication запускает приложение.
func (a *Application) StartApplication() {
	client := client.NewGRPCClient(a.Config)
	client.StartGRPCClient()

	usecase := usecase.NewGenerationService(client.Storage)

	server := server.NewGRPCServer(usecase, a.GRPCHost, a.GRPCPort)
	err := server.StartGRPCServer()
	if err != nil {
		log.Print(err)
		// обработка
	}
}

// TODO
func (a *Application) StopApplication() {

}
