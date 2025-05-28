package app

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/child6yo/rago/services/storage/internal"
	"github.com/tmc/langchaingo/schema"
)

func (a *Application) handleDocMessage(message *sarama.ConsumerMessage) error {
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

	if err := a.Db.Put(ctx, []schema.Document{doc}); err != nil {
		log.Print(err)
	}
	
	return nil
}

func unmarshalDocs(message *sarama.ConsumerMessage) (*internal.Document, error) {
	var document internal.Document
	if err := json.Unmarshal(message.Value, &document); err != nil {
		log.Printf("Failed to unmarshal: %v. Message: %s", err, string(message.Value))
		return nil, err
	}
	return &document, nil
}
