package repository

import "github.com/child6yo/rago/services/user/internal"

// Authorization определяет интерфейс репозитория авторизации.
// Отвечает за логику добавления и поиска пользователей в базе данных.
type Authorization interface {
	// CreateUser создает нового пользователя в базе данных.
	// На вход принимает модель пользователя. Возвращает ошибку в случае неудачи.
	CreateUser(user internal.User) error

	// GetUser проверяет наличие пользователя с указанными атрибутами в базе данных.
	// На вход принимает логин и пароль. Возвращает модель пользователя при успехе и ошибку в случае неудачи.
	GetUser(login, password string) (internal.User, error)
}

// ApiKey определяет интерфейс репозитория ключей API.
// Отвечает за создание, удаление и вывод апи ключей пользователями.
type ApiKey interface {
	// CreateApiKey создает в базе данных новый апи ключ для конкретного пользователя.
	// На вход принимает айди пользователя и ключ.
	CreateApiKey(userID int, key string) error

	// CreateApiKey удаляет из базы данных апи ключ по айди для конкретного пользователя.
	// На вход принимает айди ключа и айди пользователя.
	DeleteApiKey(keyID int, userID int) error

	// GetApiKeys возвращает из базы данных все апи ключи пользователя.
	// На вход принимает айди пользователя.
	GetApiKeys(userID int) ([]internal.ApiKey, error)
}
