package repository

import "github.com/child6yo/rago/services/auth/internal"

// Authorization определяет интерфейс репозитория авторизации.
type Authorization interface {
	// CreateUser создает нового пользователя в базе данных.
	// На вход принимает модель пользователя. Возвращает ошибку в случае неудачи.
	CreateUser(internal.User) error

	// GetUser проверяет наличие пользователя с указанными атрибутами в базе данных.
	// На вход принимает логин и пароль. Возвращает модель пользователя при успехе и ошибку в случае неудачи.
	GetUser(login, password string) (internal.User, error)
}
