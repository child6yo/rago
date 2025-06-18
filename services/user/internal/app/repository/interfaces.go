package repository

import "github.com/child6yo/rago/services/user/internal"

// Authorization определяет интерфейс репозитория авторизации.
// Отвечает за логику добавления и поиска пользователей в базе данных.
type Authorization interface {
	// CreateUser создает нового пользователя в базе данных.
	// На вход принимает модель пользователя и коллекцию, которая будет ему принадлежать.
	CreateUser(collection string, user internal.User) error

	// GetUser проверяет наличие пользователя с указанными атрибутами в базе данных.
	// На вход принимает логин и пароль. Возвращает модель пользователя при успехе и ошибку в случае неудачи.
	GetUser(login, password string) (internal.User, error)
}

// APIKey определяет интерфейс репозитория ключей API.
// Отвечает за создание, удаление и вывод апи ключей пользователями.
type APIKey interface {
	// CreateAPIKey создает в базе данных новый апи ключ для конкретного пользователя.
	// На вход принимает айди пользователя и ключ.
	CreateAPIKey(id string, userID int, key string) error

	// DeleteAPIKey удаляет из базы данных апи ключ по айди для конкретного пользователя.
	// На вход принимает айди ключа и айди пользователя.
	DeleteAPIKey(keyID string, userID int) error

	// GetAPIKeys возвращает из базы данных все апи ключи пользователя.
	// На вход принимает айди пользователя.
	GetAPIKeys(userID int) ([]internal.APIKey, error)

	// CheckAPIKey проверяет существование ключа апи в базе данных.
	// На вход принимает ключ. Возвращает ошибку, если его не существует.
	CheckAPIKey(key string) error
}

// Collection определяет интерфейс репозитория коллекций.
type Collection interface {
	// GetCollection возвращает коллекцию, принадлежащую пользователю.
	GetCollection(userID int) (string, error)
}
