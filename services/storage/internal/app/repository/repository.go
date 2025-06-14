package repository

import (
	"context"

	"github.com/child6yo/rago/services/storage/internal"
)

// VectorDB определяет интерфейс взаимодействия с векторным хранилищем данных.
type VectorDB interface {
	// CreateCollection создает новую коллекцию.
	CreateCollection(ctx context.Context, name string) error

	// DeleteCollection удаляет коллекцию вместе с ее содержимым.
	DeleteCollection(ctx context.Context, collection string) error

	// PutDocument загружает единицу данных в коллекцию.
	PutDocument(ctx context.Context, docs internal.Document) error

	// DeleteDocument удаляет документ из коллекции по айди.
	DeleteDocument(ctx context.Context, id string, collection string) error

	// GetDocument возвращает документ из коллекции по его айди.
	GetDocument(ctx context.Context, collection string, id string) (internal.Document, error)

	// GetAllDocuments возвращает список всех документов коллекции.
	GetAllDocuments(ctx context.Context, collection string) ([]internal.Document, error)

	// Query выполняет векторный поиск по хранилищу.
	Query(ctx context.Context, collection string, vector []float32, numDocs int) ([]internal.Document, error)
}
