package usecase

import (
	"github.com/child6yo/rago/services/user/internal"
)

// Authorization определяет интерфейс сервиса авторизации.
type Authorization interface {
	// Register регистрирует нового пользователя с его личной коллекцией (UUID).
	// Возвращает коллекцию пользователя.
	Register(user internal.User) (string, error)

	// Login проверяет наличие пользователя и корректность введенных данных.
	// Возвращает авторизационный токен.
	Login(login, password string) (string, error)

	// Auth валидирует авторизационный токен пользователя.
	// Возвращает айди пользователя.
	Auth(accessToken string) (int, error)
}

// APIKey определяет интерфейс сервиса ключей API. 
type APIKey interface {
	// CreateAPIKey создает новый API ключ для пользователя.
	CreateAPIKey(userID int) (string, error)

	// DeleteAPIKey удаляет API ключ.
	DeleteAPIKey(keyID int, userID int) error

	// GetAPIKeys возвращает все API ключи пользователя.
	GetAPIKeys(userID int) ([]internal.APIKey, error)

	// CheckAPIKey валидирует API ключ.
	CheckAPIKey(key string) error
}

// Collection определяет интерфейс сервиса коллекций.
type Collection interface {
	// GetCollection возвращает коллекцию пользователя.
	GetCollection(userID int) (string, error)
}