package usecase

import (
	"github.com/child6yo/rago/services/auth/internal"
)

// Authorization определяет интерфейс сервиса авторизации.
type Authorization interface {
	// Register регистрирует нового пользователя.
	// Возвращает ошибку в случае неудачи.
	Register(user internal.User) error

	// Login проверяет наличие пользователя и корректность введенных данных.
	// Возвращает авторизационный токен при успехе и ошибку в случае неудачи.
	Login(user internal.User) (string, error)

	// Auth валидирует авторизационный токен пользователя.
	// Возвращает ошибку в случае неудачи.
	Auth(accessToken string) error
}
