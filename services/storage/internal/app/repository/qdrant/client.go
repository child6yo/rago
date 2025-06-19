package qdrantrepo

import (
	"context"
	"log"
	"time"

	"github.com/qdrant/go-client/qdrant"
)

// Client имплементирует интерфейс VectorDB.
type Client struct {
	client *qdrant.Client // qdrant клиент
}

// NewQdrantClient создает новый экземпляр Client.
// На вход принимает хост и порт qdrant сервиса.
func NewQdrantClient(host string, port int) (*Client, error) {
	for {
		client, err := qdrant.NewClient(&qdrant.Config{
			Host: host,
			Port: port,
		})
		if err != nil {
			log.Printf("repository: failed to connect to qdrant: %v", err)
		}

		_, err = client.HealthCheck(context.Background())

		if err == nil {
			client.CreateCollection(context.Background(), &qdrant.CreateCollection{
				CollectionName: "dev_coll",
				VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
					Size:     768,
					Distance: qdrant.Distance_Cosine,
				}),
			})
			return &Client{client: client}, nil
		}

		log.Print("repository: waiting for qdrant")
		time.Sleep(4 * time.Second)
	}
}
