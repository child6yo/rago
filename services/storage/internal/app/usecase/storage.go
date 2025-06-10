package usecase

import (
	"context"

	"github.com/child6yo/rago/services/storage/internal"
	"github.com/child6yo/rago/services/storage/internal/pkg/database"
)

type StorageService struct {
	db database.VectorDB
}

// NewStorageService создает новый экземпляр StorageService.
func NewStorageService(db database.VectorDB) *StorageService {
	return &StorageService{db}
}

// Search выполняет поиск ближайших векторов.
// На вход принимает текст и количество документов, которое нужно вернуть.
// Возвращает слайс ближайших (в векторном представлении) к нему документов.
func (ss *StorageService) Search(ctx context.Context, query string, numDocs int) ([]internal.Document, error) {
	dbDocs, err := ss.db.Query(ctx, query, numDocs)
	if err != nil {
		return []internal.Document{}, err
	}

	documents := make([]internal.Document, len(dbDocs))
	for i, doc := range dbDocs {
		url, ok := doc.Metadata["URL"].(string)
		if !ok {
			url = "document without source"
		}
		documents[i] = internal.Document{
			Content:  doc.PageContent,
			Metadata: internal.Metadata{URL: url},
			Score:    doc.Score,
		}
	}

	return documents, nil
}
