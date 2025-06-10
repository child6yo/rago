package usecase

import (
	"context"

	"github.com/child6yo/rago/services/storage/internal"
)

// DocHandler определяет интерфейс обработки пришедших документов.
type DocHandler interface {
	// HandleDocMessage обрабатывает закодированные json-документы,
	// декодирует их в структуры и передает далее в векторную базу данных.
	HandleDocMessage(message []byte) error
}

// StorageService определяет интерфейс взаимодействия с векторным хранилищем.
type Storage interface {
	// Search выполняет поиск ближайших векторов.
	// На вход принимает текст и количество документов, которое нужно вернуть.
	// Возвращает слайс ближайших (в векторном представлении) к нему документов.
	Search(ctx context.Context, query string, numDocs int) ([]internal.Document, error)
}
