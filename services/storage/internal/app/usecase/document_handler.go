package usecase

import (
	"context"
	"encoding/json"
	"log"

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
// unmarshall их в структуры и передает далее в векторную базу данных.
func (d *DocHandlerService) HandleDocMessage(message []byte) error {
	rawDoc, err := unmarshalDocs(message)
	if err != nil {
		return err
	}
	doc := schema.Document{
		PageContent: rawDoc.Content,
		Metadata: map[string]any{
			"URL": rawDoc.Metadata.URL,
		}}

	ctx := context.Background()

	if err := d.db.Put(ctx, []schema.Document{doc}); err != nil {
		log.Print(err)
	}

	return nil
}

func unmarshalDocs(message []byte) (*internal.Document, error) {
	var document internal.Document
	if err := json.Unmarshal(message, &document); err != nil {
		log.Printf("Failed to unmarshal: %v. Message: %s", err, string(message))
		return nil, err
	}
	return &document, nil
}
