package usecase

import (
	"context"

	"github.com/child6yo/rago/services/storage/internal"
	"github.com/child6yo/rago/services/storage/internal/app/repository"
	"github.com/child6yo/rago/services/storage/internal/pkg/embedding"
)

// StorageService имплементирует интерфейс Storage.
type StorageService struct {
	db       repository.VectorDB
	embedder embedding.Embedder
}

// NewStorageService создает новый экземпляр StorageService.
func NewStorageService(db repository.VectorDB, embedder embedding.Embedder) *StorageService {
	return &StorageService{db, embedder}
}

// CreateCollection создает новую коллекцию.
func (ss *StorageService) CreateCollection(ctx context.Context, name string) error {
	return ss.db.CreateCollection(ctx, name)
}

// DeleteCollection удаляет коллекцию вместе с ее содержимым.
func (ss *StorageService) DeleteCollection(ctx context.Context, collection string) error {
	return ss.db.DeleteCollection(ctx, collection)
}

// DeleteDocument удаляет документ из коллекции по айди.
func (ss *StorageService) DeleteDocument(ctx context.Context, id string, collection string) error {
	return ss.db.DeleteDocument(ctx, id, collection)
}

// GetDocument возвращает документ из коллекции по его айди.
func (ss *StorageService) GetDocument(ctx context.Context, collection string, id string) (internal.Document, error) {
	return ss.db.GetDocument(ctx, collection, id)
}

// GetAllDocuments возвращает все документ из коллекции.
func (ss *StorageService) GetAllDocuments(ctx context.Context, collection string) ([]internal.Document, error) {
	return ss.db.GetAllDocuments(ctx, collection)
}

// Search выполняет поиск ближайших векторов.
// На вход принимает текст и количество документов, которое нужно вернуть.
// Возвращает слайс ближайших (в векторном представлении) к нему документов.
func (ss *StorageService) Search(ctx context.Context, collection string, query string, numDocs int) ([]internal.Document, error) {
	embs, err := ss.embedder.GenerateEmbeddings(ctx, query)
	if err != nil {
		return []internal.Document{}, err
	}

	docs, err := ss.db.Query(ctx, collection, embs, numDocs)
	if err != nil {
		return []internal.Document{}, err
	}

	return docs, nil
}
