package qdrantrepo

import (
	"context"
	"fmt"

	"github.com/child6yo/rago/services/storage/internal"
	"github.com/qdrant/go-client/qdrant"
)

// Query выполняет векторный поиск по хранилищу.
func (c *Client) Query(ctx context.Context, collection string, vector []float32, numDocs int) ([]internal.Document, error) {
	limit := uint64(numDocs)
	result, err := c.client.Query(ctx, &qdrant.QueryPoints{
		CollectionName: collection,
		Query:          qdrant.NewQuery(vector...),
		Limit:          &limit,
		WithPayload:    qdrant.NewWithPayload(true),
	})
	if err != nil {
		return []internal.Document{}, fmt.Errorf("repository; failed to query: %w", err)
	}

	docs := make([]internal.Document, len(result))
	for i, res := range result {
		payload := res.GetPayload()
		docs[i] = internal.Document{
			Content: payload["document"].GetStringValue(),
			Metadata: internal.Metadata{
				URL: payload["url"].GetStringValue(),
			},
			Score: res.Score,
		}
	}

	return docs, nil
}
