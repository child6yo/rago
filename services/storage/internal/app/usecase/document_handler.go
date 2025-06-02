package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/child6yo/rago/services/storage/internal"
	"github.com/child6yo/rago/services/storage/internal/pkg/database"
	"github.com/tmc/langchaingo/schema"
)

// DocHandlerService имплементирует интерфейс DocHandler.
type DocHandlerService struct {
	db database.VectorDB
}

// NewDocHandlerService создает новый экземпляр DocHandlerService.
func NewDocHandlerService(db database.VectorDB) *DocHandlerService {
	return &DocHandlerService{db}
}

// HandleDocMessage обрабатывает закодированные json-документы,
// декодирует их в структуры и передает далее в векторную базу данных.
func (dh *DocHandlerService) HandleDocMessage(message []byte) error {
	rawDoc, err := unmarshalDocs(message)
	if err != nil {
		return err
	}
	doc := schema.Document{
		PageContent: rawDoc.Content,
		Metadata: map[string]any{
			"URL": rawDoc.Metadata.URL,
		}}

	if err := dh.db.Put(context.Background(), []schema.Document{doc}); err != nil {
		return err
	}

	return nil
}

func unmarshalDocs(message []byte) (*internal.Document, error) {
	if len(message) == 0 {
		return nil, errors.New("empty message")
	}

	var document internal.Document
	if err := json.Unmarshal(message, &document); err != nil {
		return nil, err
	}
	
	if len(document.Content) == 0 {
		return nil, errors.New("empty content")
	}

	return &document, nil
}
