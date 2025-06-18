package repository

import (
	"fmt"

	"github.com/child6yo/rago/services/user/internal"
	"github.com/jmoiron/sqlx"
)

// AuthorizationRepository имплементирует интерфейс Authorization.
type AuthorizationRepository struct {
	db *sqlx.DB
}

// NewAuthorizationRepository создает новый экземпляр AuthorizationRepository.
func NewAuthorizationRepository(db *sqlx.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db}
}

// CreateUser создает нового пользователя в базе данных.
// На вход принимает модель пользователя и коллекцию, которая будет ему принадлежать.
func (ar *AuthorizationRepository) CreateUser(collection string, user internal.User) error {
    exists, err := ar.UserExists(user.Login)
    if err != nil {
        return fmt.Errorf("repository (CreateUser): failed to check user existence: %w", err)
    }
    if exists {
        return fmt.Errorf("repository (CreateUser): user with login '%s' already exists", user.Login)
    }

    // Вставляем коллекцию и сразу получаем её ID
    var collectionID int64
    query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", collectionTable)
    err = ar.db.QueryRow(query, collection).Scan(&collectionID)
    if err != nil {
        return fmt.Errorf("repository (CreateUser): failed to create collection: %w", err)
    }

    // Вставляем пользователя
    query = fmt.Sprintf("INSERT INTO %s (login, password_hash, collection_id) VALUES ($1, $2, $3)", userTable)
    _, err = ar.db.Exec(query, user.Login, user.Password, collectionID)
    if err != nil {
        return fmt.Errorf("repository (CreateUser): failed to create user: %w", err)
    }

    return nil
}

// GetUser проверяет наличие пользователя с указанными атрибутами в базе данных.
// На вход принимает логин и пароль. Возвращает ошибку в случае неудачи.
func (ar *AuthorizationRepository) GetUser(login, password string) (_ internal.User, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("repository (GetUser): %w", err)
		}
	}()

	var user internal.User

	query := fmt.Sprintf("SELECT id, login, password_hash AS password FROM %s WHERE login=$1 AND password_hash=$2", userTable)
	err = ar.db.Get(&user, query, login, password)

	return user, err
}

// UserExists проверяет, существует ли пользователь с таким логином.
func (ar *AuthorizationRepository) UserExists(login string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE login = $1", userTable)
	var count int
	err := ar.db.QueryRow(query, login).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("repository (UserExists): %w", err)
	}
	return count > 0, nil
}
