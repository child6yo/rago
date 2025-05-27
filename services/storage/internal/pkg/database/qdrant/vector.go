package qdrant

import (
	"context"

	"github.com/tmc/langchaingo/schema"
)

// Query позволяет выполнить векторный поиск по хранилищу.
func (q *Qdrant) Query(ctx context.Context, query string, numDocs int) ([]schema.Document, error) {
	ans, err := q.Store.SimilaritySearch(ctx, query, numDocs)
	if err != nil {
		return []schema.Document{}, err
	}

	return ans, nil
}
