package qdrantrepo

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

// CreateCollection создает новую коллекцию.
func (c *Client) CreateCollection(ctx context.Context, name string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("repository; failed to create collection: %w", err)
		}
	}()
	return c.client.CreateCollection(ctx, &qdrant.CreateCollection{
		CollectionName: name,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     768,
			Distance: qdrant.Distance_Cosine,
		}),
	})
}

// DeleteCollection удаляет коллекцию вместе с ее содержимым.
func (c *Client) DeleteCollection(ctx context.Context, collection string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("repository; failed to delete collection: %w", err)
		}
	}()
	return c.client.DeleteCollection(ctx, collection)
}
