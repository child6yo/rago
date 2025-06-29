package usecase

import (
	"context"

	"github.com/child6yo/rago/services/storage/internal"
)

// DocumentLoader определяет интерфейс загрузчика документов.
type DocumentLoader interface {
	// LoadDocument обрабатывает закодированные json-документы,
	// декодирует их в структуры и передает далее в векторную базу данных.
	LoadDocument(ctx context.Context, message []byte) error
}

// Storage определяет интерфейс взаимодействия с векторным хранилищем.
type Storage interface {
	// CreateCollection создает новую коллекцию.
	CreateCollection(ctx context.Context, name string) error

	// DeleteCollection удаляет коллекцию вместе с ее содержимым.
	DeleteCollection(ctx context.Context, collection string) error

	// DeleteDocument удаляет документ из коллекции по айди.
	DeleteDocument(ctx context.Context, id string, collection string) error

	// GetDocument возвращает документ из коллекции по его айди.
	GetDocument(ctx context.Context, collection string, id string) (internal.Document, error)

	// GetAllDocuments возвращает все документ из коллекции.
	GetAllDocuments(ctx context.Context, collection string) ([]internal.Document, error)

	// Search выполняет поиск ближайших векторов.
	// На вход принимает текст и количество документов, которое нужно вернуть.
	// Возвращает слайс ближайших (в векторном представлении) к нему документов.
	Search(ctx context.Context, collection string, query string, numDocs int) ([]internal.Document, error)
}
