package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	// postgres driver
	_ "github.com/lib/pq"
)

const (
	userTable = "users"
	apiKeyTable = "api_keys"
)

// NewPostgresDB создает новое подключение к базе данных postgres.
func NewPostgresDB(host, port, username, dbName, password, sslMode string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, username, dbName, password, sslMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
