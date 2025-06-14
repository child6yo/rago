package mock

import (
	"context"

	"github.com/child6yo/rago/services/storage/internal"
)

// VectorDBMock мок интерфейса VectorDB.
type VectorDBMock struct {
    CreateCollectionFunc func(ctx context.Context, name string) error
    DeleteCollectionFunc func(ctx context.Context, collection string) error
    PutDocumentFunc      func(ctx context.Context, docs internal.Document) error
    DeleteDocumentFunc   func(ctx context.Context, id string, collection string) error
    GetDocumentFunc      func(ctx context.Context, collection string, id string) (internal.Document, error)
    GetAllDocumentsFunc  func(ctx context.Context, collection string) ([]internal.Document, error)
    QueryFunc           func(ctx context.Context, query string, numDocs int) ([]internal.Document, error)
}

// CreateCollection mock
func (m *VectorDBMock) CreateCollection(ctx context.Context, name string) error {
    if m.CreateCollectionFunc != nil {
        return m.CreateCollectionFunc(ctx, name)
    }
    return nil
}

// DeleteCollection mock
func (m *VectorDBMock) DeleteCollection(ctx context.Context, collection string) error {
    if m.DeleteCollectionFunc != nil {
        return m.DeleteCollectionFunc(ctx, collection)
    }
    return nil
}

// PutDocument mock
func (m *VectorDBMock) PutDocument(ctx context.Context, docs internal.Document) error {
    if m.PutDocumentFunc != nil {
        return m.PutDocumentFunc(ctx, docs)
    }
    return nil
}

// DeleteDocument mock
func (m *VectorDBMock) DeleteDocument(ctx context.Context, id string, collection string) error {
    if m.DeleteDocumentFunc != nil {
        return m.DeleteDocumentFunc(ctx, id, collection)
    }
    return nil
}

// GetDocument mock
func (m *VectorDBMock) GetDocument(ctx context.Context, collection string, id string) (internal.Document, error) {
    if m.GetDocumentFunc != nil {
        return m.GetDocumentFunc(ctx, collection, id)
    }
    return internal.Document{}, nil
}

// GetAllDocuments mock
func (m *VectorDBMock) GetAllDocuments(ctx context.Context, collection string) ([]internal.Document, error) {
    if m.GetAllDocumentsFunc != nil {
        return m.GetAllDocumentsFunc(ctx, collection)
    }
    return []internal.Document{}, nil
}

// Query mock
func (m *VectorDBMock) Query(ctx context.Context, query string, numDocs int) ([]internal.Document, error) {
    if m.QueryFunc != nil {
        return m.QueryFunc(ctx, query, numDocs)
    }
    return []internal.Document{}, nil
}