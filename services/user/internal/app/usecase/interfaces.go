package usecase

import (
	"github.com/child6yo/rago/services/user/internal"
)

// Authorization определяет интерфейс сервиса авторизации.
type Authorization interface {
	// Register регистрирует нового пользователя.
	// Возвращает ошибку в случае неудачи.
	Register(user internal.User) error

	// Login проверяет наличие пользователя и корректность введенных данных.
	// Возвращает авторизационный токен при успехе и ошибку в случае неудачи.
	Login(login, password string) (string, error)

	// Auth валидирует авторизационный токен пользователя.
	// Возвращает ошибку в случае неудачи.
	Auth(accessToken string) error
}

// ApiKey определяет интерфейс сервиса ключей API. 
type ApiKey interface {
	CreateApiKey(userID int, key string) error
	DeleteApiKey(keyID int, userID int) error
	GetApiKeys(userID int) ([]internal.ApiKey, error)
}