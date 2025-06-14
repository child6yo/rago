package qdrantrepo

import "context"

// CreateCollection создает новую коллекцию.
func (c *Client) CreateCollection(ctx context.Context, name string) error {
	return nil
}

// DeleteCollection удаляет коллекцию вместе с ее содержимым.
func (c *Client) DeleteCollection(ctx context.Context, collection string) error {
	return nil
}
