package usecase

import (
	"context"

	"github.com/child6yo/rago/services/storage/internal"
	"github.com/child6yo/rago/services/storage/internal/app/repository"
)

// StorageService имплементирует интерфейс Storage.
type StorageService struct {
	db repository.VectorDB
}

// NewStorageService создает новый экземпляр StorageService.
func NewStorageService(db repository.VectorDB) *StorageService {
	return &StorageService{db}
}

// Search выполняет поиск ближайших векторов.
// На вход принимает текст и количество документов, которое нужно вернуть.
// Возвращает слайс ближайших (в векторном представлении) к нему документов.
func (ss *StorageService) Search(ctx context.Context, query string, numDocs int) ([]internal.Document, error) {
	docs, err := ss.db.Query(ctx, query, numDocs)
	if err != nil {
		return []internal.Document{}, err
	}

	return docs, nil
}
