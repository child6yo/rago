package app

import (
	"log"

	"github.com/child6yo/rago/services/user/internal/app/repository"
	"github.com/child6yo/rago/services/user/internal/app/server"
	"github.com/child6yo/rago/services/user/internal/app/usecase"
	"github.com/child6yo/rago/services/user/internal/config"
	"github.com/jmoiron/sqlx"
)

// Application - структура приложения splitter.
type Application struct {
	config.Config // конфигурация

	db     *sqlx.DB // соденинение с базой данных
	server *server.GRPCServer
}

// CreateApplication создает новый экземпляр приложения.
func CreateApplication(config config.Config) *Application {
	return &Application{Config: config}
}

// StartApplication запускает приложение.
func (a *Application) StartApplication() {
	db, err := repository.NewPostgresDB(
		a.PgHost, a.PgPort, a.PgUsername,
		a.PgDBName, a.PgPassword, a.PgSSLMode,
	)
	if err != nil {
		log.Panic(err)
	}
	a.db = db

	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo)

	a.server = server.NewGRPCServer(usecase, a.GRPCHost, a.GRPCPort)
	go func() {
		if err := a.server.StartGRPCServer(); err != nil {
			log.Fatal(err)
		}
	}()
}

// StopApplication завершает работу приложения.
func (a *Application) StopApplication() {
	a.server.ShutdownGRPCServer()
	a.db.Close()
}
