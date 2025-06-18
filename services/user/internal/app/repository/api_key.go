package repository

import (
	"errors"
	"fmt"

	"github.com/child6yo/rago/services/user/internal"
	"github.com/jmoiron/sqlx"
)

// APIKeyRepository имплементирует интерфейс APIKey.
type APIKeyRepository struct {
	db *sqlx.DB
}

// NewAPIKeyRepository создает новый экземпляр APIKeyRepository.
func NewAPIKeyRepository(db *sqlx.DB) *APIKeyRepository {
	return &APIKeyRepository{db}
}

// CreateAPIKey регистрирует новый апи ключ для конкретного пользователя.
// На вход принимает айди пользователя и ключ.
func (akr *APIKeyRepository) CreateAPIKey(userID int, key string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("repository (CreateAPIKey): %w", err)
		}
	}()
	query := fmt.Sprintf("INSERT INTO %s (user_id, key) values ($1, $2)", apiKeyTable)
	_, err = akr.db.Exec(query, userID, key)

	return err
}

// DeleteAPIKey удаляет из базы данных апи ключ по айди для конкретного пользователя.
// На вход принимает айди ключа и айди пользователя.
func (akr *APIKeyRepository) DeleteAPIKey(keyID int, userID int) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("repository (DeleteAPIKey): %w", err)
		}
	}()
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", apiKeyTable)
	_, err = akr.db.Exec(query, keyID, userID)

	return err
}

// GetAPIKeys возвращает из базы данных все апи ключи пользователя.
// На вход принимает айди пользователя.
func (akr *APIKeyRepository) GetAPIKeys(userID int) (_ []internal.APIKey, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("repository (GetAPIKeys): %w", err)
		}
	}()
	var keys []internal.APIKey

	query := fmt.Sprintf("SELECT id, key FROM %s WHERE user_id=$1", apiKeyTable)
	err = akr.db.Select(&keys, query, userID)

	return keys, err
}

// CheckAPIKey проверяет существование ключа апи в базе данных.
// На вход принимает ключ. Возвращает ошибку, если его не существует.
func (akr *APIKeyRepository) CheckAPIKey(key string) error {
	var id int

	query := fmt.Sprintf("SELECT id FROM %s WHERE key=$1", apiKeyTable)
	err := akr.db.Get(&id, query, key)
	if err != nil {
		return fmt.Errorf("repository (CheckAPIKey): %w", err)
	}

	if id == 0 {
		return errors.New("repository (CheckAPIKey): api key doesn't exists")
	}

	return nil
}
