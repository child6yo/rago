package qdrantrepo

import (
	"context"
	"fmt"

	"github.com/child6yo/rago/services/storage/internal"
	"github.com/google/uuid"
	"github.com/qdrant/go-client/qdrant"
)

// PutDocument загружает единицу данных в коллекцию.
func (c *Client) PutDocument(ctx context.Context, docs internal.Document, collection string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("repository; failed to put document: %w", err)
		}
	}()

	// создает новый uuid
	id := uuid.NewString()

	_, err = c.client.Upsert(ctx, &qdrant.UpsertPoints{
		CollectionName: collection,
		Points: []*qdrant.PointStruct{
			{
				Id:      qdrant.NewIDUUID(id),
				Vectors: qdrant.NewVectors(docs.Embedding...),
				Payload: qdrant.NewValueMap(map[string]any{"document": docs.Content, "url": docs.Metadata.URL}),
			},
		},
	})

	return err
}

// DeleteDocument удаляет документ из коллекции по айди.
func (c *Client) DeleteDocument(ctx context.Context, id string, collection string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("repository; failed to delete document: %w", err)
		}
	}()

	_, err = c.client.Delete(ctx, &qdrant.DeletePoints{
		CollectionName: collection,
		Points: qdrant.NewPointsSelectorIDs([]*qdrant.PointId{
			qdrant.NewIDUUID(id),
		}),
	})
	return err
}

// GetDocument возвращает документ из коллекции по его айди.
func (c *Client) GetDocument(ctx context.Context, collection string, id string) (internal.Document, error) {
	limit := uint32(1)
	points, err := c.client.Scroll(ctx, &qdrant.ScrollPoints{
		CollectionName: collection,
		Offset:         qdrant.NewIDUUID(id),
		Limit:          &limit,
	},
	)
	if err != nil {
		return internal.Document{}, fmt.Errorf("repository; failed to scroll collection: %w", err)
	}

	point := points[0]

	doc := internal.Document{
		ID:      point.Id.GetUuid(),
		Content: point.GetPayload()["document"].GetStringValue(),
		Metadata: internal.Metadata{
			URL: point.GetPayload()["url"].GetStringValue(),
		},
	}

	return doc, nil
}

// GetAllDocuments возвращает список всех документов коллекции.
func (c *Client) GetAllDocuments(ctx context.Context, collection string) ([]internal.Document, error) {
	points, err := c.client.Scroll(ctx, &qdrant.ScrollPoints{
		CollectionName: collection,
	},
	)
	if err != nil {
		return []internal.Document{}, fmt.Errorf("repository; failed to scroll collection: %w", err)
	}

	docs := make([]internal.Document, len(points))
	for i, point := range points {
		docs[i] = internal.Document{
			ID:      point.Id.GetUuid(),
			Content: point.GetPayload()["document"].GetStringValue(),
			Metadata: internal.Metadata{
				URL: point.GetPayload()["url"].GetStringValue(),
			},
		}
	}

	return docs, nil
}
