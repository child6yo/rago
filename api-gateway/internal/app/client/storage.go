package client

import (
	"context"
	"fmt"
	"log"

	"github.com/child6yo/rago/api-gateway/internal"
	"github.com/child6yo/rago/services/storage/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Storage определяет клиент сервиса хранилища, доступного по gRPC.
type Storage struct {
	storage pb.StorageServiceClient
	conn    *grpc.ClientConn

	host string
	port string
}

func newStorageClient(host string, port string) *Storage {
	return &Storage{host: host, port: port}
}

func (sc *Storage) startStorageClient() {
	addr := fmt.Sprintf("%s:%s", sc.host, sc.port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print("failed to connect splitter grpc server")
	}

	sc.storage = pb.NewStorageServiceClient(conn)
	sc.conn = conn
}

func (sc *Storage) stopStorageClient() {
	if sc.conn == nil {
		return
	}
	sc.conn.Close()
}

// CreateCollection вызывает удаленный метод создания новой коллекции.
func (sc *Storage) CreateCollection(ctx context.Context, collection string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("storage client error (CreateCollection): %v", err)
		}
	}()

	_, err = sc.storage.CreateCollection(ctx, &pb.CollectionRequest{CollectionName: collection})
	return err
}

// DeleteCollection вызывает удаленный метод удаления коллекции.
func (sc *Storage) DeleteCollection(ctx context.Context, collection string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("storage client error (DeleteCollection): %v", err)
		}
	}()

	_, err = sc.storage.DeleteCollection(ctx, &pb.CollectionRequest{CollectionName: collection})
	return err
}

// DeleteDocument вызывает удаленный метод удаления документа по айди.
func (sc *Storage) DeleteDocument(ctx context.Context, collection string, docID string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("storage client error (DeleteDocument): %v", err)
		}
	}()

	_, err = sc.storage.DeleteDocument(ctx, &pb.DocumentRequest{CollectionName: collection, Id: docID})
	return err
}

// GetDocument вызывает удаленный метод получения документа по айди
func (sc *Storage) GetDocument(ctx context.Context, collection string, docID string) (*internal.Document, error) {
	doc, err := sc.storage.GetDocument(ctx, &pb.DocumentRequest{CollectionName: collection, Id: docID})
	if err != nil {
		return nil, fmt.Errorf("storage client error (GetDocument): %v", err)
	}

	return &internal.Document{
		ID:       doc.Id,
		Content:  doc.Content,
		Metadata: internal.Metadata{URL: doc.Metadata.Url},
	}, nil
}

// GetAllDocuments вызывает удаленный метод получения всех документов коллекции.
func (sc *Storage) GetAllDocuments(ctx context.Context, collection string) (*internal.DocumentArray, error) {
	docs, err := sc.storage.GetAllDocuments(ctx, &pb.CollectionRequest{CollectionName: collection})
	if err != nil {
		return nil, fmt.Errorf("storage client error (GetAllDocuments): %v", err)
	}

	result := make([]internal.Document, len(docs.Document))
	for i, doc := range docs.Document {
		result[i] = internal.Document{
			ID:       doc.Id,
			Content:  doc.Content,
			Metadata: internal.Metadata{URL: doc.Metadata.Url},
		}
	}

	return &internal.DocumentArray{Documents: result}, nil
}
