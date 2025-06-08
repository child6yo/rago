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
// На вход принимает модель пользователя. Возвращает ошибку в случае неудачи.
func (ar *AuthorizationRepository) CreateUser(user internal.User) error {
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash) values ($1, $2)", userTable)
	_, err := ar.db.Exec(query, user.Login, user.Password)

	return err
}

// GetUser проверяет наличие пользователя с указанными атрибутами в базе данных.
// На вход принимает логин и пароль. Возвращает ошибку в случае неудачи.
func (ar *AuthorizationRepository) GetUser(login, password string) (internal.User, error) {
	var user internal.User

	query := fmt.Sprintf("SELECT id, login, password_hash AS password FROM %s WHERE login=$1 AND password_hash=$2", userTable)
	err := ar.db.Get(&user, query, login, password)

	return user, err
}
