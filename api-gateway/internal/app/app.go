package app

import "github.com/child6yo/rago/api-gateway/internal/config"

// Application - структура приложения splitter.
type Application struct {
	config.Config // конфигурация
}

// CreateApplication создает новый экземпляр приложения.
func CreateApplication(config config.Config) *Application {
	return &Application{config}
}

// StartApplication запускает приложение.
func (a *Application) StartApplication() {}

// TODO
func (a *Application) StopApplication() {

}
