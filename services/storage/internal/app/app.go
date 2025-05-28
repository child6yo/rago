package app

import (
	"context"
	"sync"

	"github.com/child6yo/rago/services/storage/internal/config"
)

// Application - структура приложения хранилища.
type Application struct {
	config.Config // конфигурация
	context       context.Context
	cancel        context.CancelFunc

	kfkWG *sync.WaitGroup // kafka consumer waitgroup
}

// CreateApplication создает новый экземпляр приложения.
func CreateApplication(config config.Config) *Application {
	ctx, cancel := context.WithCancel(context.Background())
	return &Application{Config: config, context: ctx, cancel: cancel}
}

// StartApplication запускает приложение.
func (a *Application) StartApplication() {
	runConsumer(a.KafkaAddr, "group.storage", []string{"document-topic"}, a)
}

// TODO
func (a *Application) StopApplication() {
	stopConsumer()
}