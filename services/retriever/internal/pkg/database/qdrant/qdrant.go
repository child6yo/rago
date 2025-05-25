package qdrant

import (
	"net/url"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

type Qdrant struct {
	Store qdrant.Store
}

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
