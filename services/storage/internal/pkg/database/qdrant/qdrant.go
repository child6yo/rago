package qdrant

import (
	"net/url"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

// Qdrant имплементирует интерфейс VectorDB для
// векторной базы данных qdrant.
type Qdrant struct {
	Store qdrant.Store
}

// NewQdrantConnection открывает соединение с qdrant.
//
// Параметры:
// 	- url - адрес базы данных
// 	- embedder - эмбеддинговая модель
// 	- collectionName - имя коллекции в базе данных
func NewQdrantConnection(url *url.URL, embedder embeddings.Embedder, collectionName string) (*Qdrant, error) {
	store, err := qdrant.New(
		qdrant.WithURL(*url),
		qdrant.WithCollectionName(collectionName),
		qdrant.WithEmbedder(embedder),
	)
	if err != nil {
		return nil, err
	}
	
	return &Qdrant{Store: store}, nil
}
