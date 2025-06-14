package qdrantrepo

import (
	"context"

	"github.com/child6yo/rago/services/storage/internal"
)

// PutDocument загружает единицу данных в коллекцию.
func (c *Client) PutDocument(ctx context.Context, docs internal.Document) error {
	return nil
}

// DeleteDocument удаляет документ из коллекции по айди.
func (c *Client) DeleteDocument(ctx context.Context, id string, collection string) error {
	return nil
}

// GetDocument возвращает документ из коллекции по его айди.
func (c *Client) GetDocument(ctx context.Context, collection string, id string) (internal.Document, error) {
	return internal.Document{}, nil
}

// GetAllDocuments возвращает список всех документов коллекции.
func (c *Client) GetAllDocuments(ctx context.Context, collection string) ([]internal.Document, error) {
	return []internal.Document{}, nil
}
