package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/child6yo/rago/services/storage/internal"
	"github.com/child6yo/rago/services/storage/internal/app/repository"
)

// DocumentLoader имплементирует интерфейс DocHandler.
type Loader struct {
	db repository.VectorDB
}

// NewLoader создает новый экземпляр Loader.
func NewLoader(db repository.VectorDB) *Loader {
	return &Loader{db}
}

// LoadDocument обрабатывает закодированные json-документы,
// декодирует их в структуры и передает далее в векторную базу данных.
func (l *Loader) LoadDocument(message []byte) error {
	doc, err := unmarshalDocs(message)
	if err != nil {
		return err
	}

	if err := l.db.PutDocument(context.Background(), doc); err != nil {
		return err
	}

	return nil
}

func unmarshalDocs(message []byte) (internal.Document, error) {
	if len(message) == 0 {
		return internal.Document{}, errors.New("empty message")
	}

	var document internal.Document
	if err := json.Unmarshal(message, &document); err != nil {
		return internal.Document{}, err
	}

	if len(document.Content) == 0 {
		return internal.Document{}, errors.New("empty content")
	}

	return document, nil
}
