package repository

import (
	"fmt"
	"log"

	"github.com/child6yo/rago/services/user/internal"
	"github.com/jmoiron/sqlx"
)

// ApiKeyRepository имплементирует интерфейс ApiKey.
type ApiKeyRepository struct {
	db *sqlx.DB
}

// NewApiKeyRepository создает новый экземпляр ApiKeyRepository.
func NewApiKeyRepository(db *sqlx.DB) *ApiKeyRepository {
	return &ApiKeyRepository{db}
}

// CreateApiKey регистрирует новый апи ключ для конкретного пользователя.
// На вход принимает айди пользователя и ключ.
func (akr *ApiKeyRepository) CreateApiKey(userID int, key string) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, key) values ($1, $2)", apiKeyTable)
	_, err := akr.db.Exec(query, userID, key)

	return err
}

// CreateApiKey удаляет апи ключ по айди для конкретного пользователя.
// На вход принимает айди ключа и айди пользователя.
func (akr *ApiKeyRepository) DeleteApiKey(keyID int, userID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", apiKeyTable)
	_, err := akr.db.Exec(query, keyID, userID)

	return err
}

// GetApiKeys возвращает все апи ключи пользователя.
// На вход принимает айди пользователя.
func (akr *ApiKeyRepository) GetApiKeys(userID int) ([]internal.ApiKey, error) {
	var keys []internal.ApiKey

	query := fmt.Sprintf("SELECT id, key FROM %s WHERE user_id=$1", apiKeyTable)
	err := akr.db.Select(&keys, query, userID)

	log.Print("repo:", err)

	return keys, err
}
