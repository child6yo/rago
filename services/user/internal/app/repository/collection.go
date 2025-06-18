package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// ColletionRepository имплементирует интерфейс Collection.
type ColletionRepository struct {
	db *sqlx.DB
}

// NewColletionRepository создает новый экземпляр ColletionRepository.
func NewColletionRepository(db *sqlx.DB) *ColletionRepository {
	return &ColletionRepository{db}
}

// GetCollection возвращает коллекцию, принадлежащую пользователю.
func (cr *ColletionRepository) GetCollection(userID int) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("repository (GetCollection): %w", err)
		}
	}()

	query := fmt.Sprintf(`SELECT collections.name FROM %s
						INNER JOIN %s 
						ON users.collection_id = collections.id
						WHERE users.id = $1`, userTable, collectionTable)

	var collection string
	err = cr.db.Get(&collection, query, userID)

	return collection, err
}
