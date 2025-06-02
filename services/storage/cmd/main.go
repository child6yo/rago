package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/child6yo/rago/services/storage/internal/app"
	"github.com/child6yo/rago/services/storage/internal/config"
)

func main() {
	// инициализация конфигурации
	cfg := config.InitConfig()

	// создание экземпляра приложения
	app := app.CreateApplication(cfg)

	// запуск приложения
	app.StartApplication()

	// получение сигнала на остановку приложения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// остановка приложения
	app.StopApplication()
}
