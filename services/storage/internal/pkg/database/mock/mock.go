package mock

import (
	"context"

	"github.com/tmc/langchaingo/schema"
)

// MockVectorDB мок интерфейса VectorDB.
type MockVectorDB struct {
	PutFunc   func(ctx context.Context, docs []schema.Document) error
	QueryFunc func(ctx context.Context, query string, numDocs int) ([]schema.Document, error)
}

// Put mock
func (m *MockVectorDB) Put(ctx context.Context, docs []schema.Document) error {
	return m.PutFunc(ctx, docs)
}

// Query mock
func (m *MockVectorDB) Query(ctx context.Context, query string, numDocs int) ([]schema.Document, error) {
	return m.QueryFunc(ctx, query, numDocs)
}
