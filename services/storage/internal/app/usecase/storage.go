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
