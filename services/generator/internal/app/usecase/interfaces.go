package usecase

import "context"

// Generation определяет интерфейс сервиса генерации на основе контекста.
type Generation interface {
	// Generate генерирует ответ за запрос. Самостоятельно идет в сервис хранения данных.
	// Ответ генерирует порционно. Возвращает канал, в который будет стримить поток ответа.
	// Возвращает ошибку в случае неполадок.
	Generate(ctx context.Context, query string, collection string) (<-chan string, error)
}
