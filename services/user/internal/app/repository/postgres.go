package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	// postgres driver
	_ "github.com/lib/pq"
)

const (
	userTable       = "users"
	apiKeyTable     = "api_keys"
	collectionTable = "collections"
)

// NewPostgresDB создает новое подключение к базе данных postgres.
func NewPostgresDB(host, port, username, dbName, password, sslMode string) (*sqlx.DB, error) {
	for {
		db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			host, port, username, dbName, password, sslMode))
		if err != nil {
			log.Printf("repository: invalid sqlx connection data: %v", err)
		}

		err = db.Ping()

		if err == nil {
			return db, nil
		}

		log.Print("repository: waiting for postgres")
		time.Sleep(4 * time.Second)
	}

}
