package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/child6yo/rago/services/storage/internal"
	"github.com/child6yo/rago/services/storage/internal/app/repository"
	"github.com/child6yo/rago/services/storage/internal/pkg/embedding"
)

// Loader имплементирует интерфейс DocumentLoader.
type Loader struct {
	db       repository.VectorDB
	embedder embedding.Embedder
}

// NewLoader создает новый экземпляр Loader.
func NewLoader(db repository.VectorDB, embedder embedding.Embedder) *Loader {
	return &Loader{db, embedder}
}

// LoadDocument обрабатывает закодированные json-документы,
// декодирует их в структуры и передает далее в векторную базу данных.
func (l *Loader) LoadDocument(ctx context.Context, message []byte) error {
	doc, err := unmarshalDocs(message)
	if err != nil {
		return err
	}

	embs, err := l.embedder.GenerateEmbeddings(ctx, doc.Content)
	if err != nil {
		return err
	}

	doc.Embedding = embs

	if err := l.db.PutDocument(ctx, doc); err != nil {
		return err
	}

	return nil
}

func unmarshalDocs(message []byte) (internal.Document, error) {
	if len(message) == 0 {
		return internal.Document{}, errors.New("loader: empty message")
	}

	var document internal.Document
	if err := json.Unmarshal(message, &document); err != nil {
		return internal.Document{}, fmt.Errorf("loader: %w", err)
	}

	if len(document.Content) == 0 {
		return internal.Document{}, errors.New("loader: empty content")
	}

	return document, nil
}
