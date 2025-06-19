package client

import (
	"context"
	"fmt"
	"log"

	"github.com/child6yo/rago/services/generator/internal"
	pb "github.com/child6yo/rago/proto/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Storage определяет структуру клиента для gRPC сервера вектоного хранилища.
type Storage struct {
	client pb.StorageServiceClient
	conn   *grpc.ClientConn

	host string
	port string
}

func newStorageClient(host string, port string) *Storage {
	return &Storage{host: host, port: port}
}

func (s *Storage) startStorageClient() {
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print("failed to connect splitter grpc server")
	}

	s.client = pb.NewStorageServiceClient(conn)
	s.conn = conn
}

func (s *Storage) stopStorageClient() {
	if s.conn == nil {
		return
	}
	s.conn.Close()
}

// Search вызывает удаленную функцию векторного поиска по хранилищу.
func (s *Storage) Search(ctx context.Context, query string, limit int, collection string) ([]internal.Document, error) {
	resp, err := s.client.Search(ctx, &pb.QueryRequest{
		Query:          query,
		Limit:          int32(limit),
		CollectionName: collection,
	})
	if err != nil {
		return []internal.Document{}, fmt.Errorf("storage client (generator): failed to search: %v", err)
	}

	docsPb := resp.Document
	documents := make([]internal.Document, len(docsPb))
	for i, doc := range docsPb {
		documents[i] = internal.Document{
			Content:  doc.Content,
			Metadata: internal.Metadata{URL: doc.Metadata.Url},
			Score:    doc.Score,
		}
	}

	return documents, nil
}
