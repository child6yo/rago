package repository

import "github.com/jmoiron/sqlx"

// Repository содержит все интерфейсы пакета репозитория.
type Repository struct {
	Authorization
	APIKey
	Collection
}

// NewRepository создает новый репозиторий.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepository(db),
		APIKey:        NewAPIKeyRepository(db),
		Collection:    NewColletionRepository(db),
	}
}
