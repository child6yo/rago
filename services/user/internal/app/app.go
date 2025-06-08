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

	db *sqlx.DB // соденинение с базой данных
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

	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo)

	server := server.NewGRPCServer(usecase, a.GRPCHost, a.GRPCPort)
	err = server.StartGRPCServer()
	if err != nil {
		log.Print(err)
		// обработка
	}
}

// TODO
func (a *Application) StopApplication() {

}
