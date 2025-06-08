package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Authorization
	ApiKey
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepository(db),
		ApiKey: NewApiKeyRepository(db),
	}
}
