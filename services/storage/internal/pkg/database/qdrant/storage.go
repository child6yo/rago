package qdrant

import (
	"context"

	"github.com/tmc/langchaingo/schema"
)

// Put позволяет загрузить единицу данных в хранилище.
func (q *Qdrant) Put(ctx context.Context, docs []schema.Document) error {
	if _, err := q.Store.AddDocuments(ctx, docs); err != nil {
		return err
	}
	return nil
}
