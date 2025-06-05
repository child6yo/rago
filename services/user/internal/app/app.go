package app

import (
	"log"

	"github.com/child6yo/rago/services/auth/internal/app/repository"
	"github.com/child6yo/rago/services/auth/internal/app/server"
	"github.com/child6yo/rago/services/auth/internal/app/usecase"
	"github.com/child6yo/rago/services/auth/internal/config"
	"github.com/jmoiron/sqlx"
)

// Application - структура приложения splitter.
type Application struct {
	config.Config // конфигурация

	db *sqlx.DB
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
		log.Print(err)
		// обработка
	}
	a.db = db

	repo := repository.NewAuthorizationRepository(db)
	auth := usecase.NewAuthorizationService(repo)

	server := server.NewGRPCServer(auth, a.GRPCHost, a.GRPCPort)
	err = server.StartGRPCServer()
	if err != nil {
		log.Print(err)
		// обработка
	}
}

// TODO
func (a *Application) StopApplication() {

}
