package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/child6yo/rago/services/user/internal/app"
	"github.com/child6yo/rago/services/user/internal/config"
)

func main() {
	// инициализация конфигурации
	cfg := config.InitConfig()

	// создание экземпляра приложения
	app := app.CreateApplication(cfg)

	// запуск приложения
	app.StartApplication()
	log.Printf("user service successfully started")

	// получение сигнала на остановку приложения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// остановка приложения
	app.StopApplication()
}
