package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/child6yo/rago/services/splitter/internal/app"
	"github.com/child6yo/rago/services/splitter/internal/config"
)

func main() {
	// инициализация конфигурации
	cfg := config.InitConfig()

	// создание экземпляра приложения
	app := app.CreateApplication(cfg)

	// запуск приложения
	app.StartApplication()
	log.Printf("splitter service successfully started")

	// получение сигнала на остановку приложения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// остановка приложения
	app.StopApplication()
}
