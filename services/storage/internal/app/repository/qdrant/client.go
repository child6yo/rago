package qdrantrepo

import (
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

// Client имплементирует интерфейс VectorDB.
type Client struct {
	client *qdrant.Client // qdrant клиент
}

// NewQdrantClient создает новый экземпляр Client.
// На вход принимает хост и порт qdrant сервиса.
func NewQdrantClient(host string, port int) (*Client, error) {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: host,
		Port: port,
	})
	if err != nil {
		return nil, fmt.Errorf("qdrant repo: %w", err)
	}

	return &Client{client: client}, nil
}