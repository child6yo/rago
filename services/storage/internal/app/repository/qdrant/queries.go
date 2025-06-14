package qdrantrepo

import (
	"context"

	"github.com/child6yo/rago/services/storage/internal"
)

// Query выполняет векторный поиск по хранилищу.
func (c *Client) Query(ctx context.Context, query string, numDocs int) ([]internal.Document, error) {
	return []internal.Document{}, nil
}
