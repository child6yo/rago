package server

import (
	"context"

	"github.com/child6yo/rago/services/storage/internal/app/usecase"
	"github.com/child6yo/rago/services/storage/pkg/pb"
)

// AuthServiceServer определяет интерфейс gRPC сервера сервиса векторного хранилища.
type StorageServiceServer interface {
	Search(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error)
	mustEmbedUnimplementedStorageServiceServer()
}

// StorageService имплементирует интерфейс StorageService.
type StorageService struct {
	pb.StorageServiceServer
	service usecase.Storage
}

// Search реализует удаленный метод поиска в векторном хранилище.
func (s *StorageService) Search(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	docs, err := s.service.Search(ctx, req.Query, int(req.Limit))

	documents := make([]*pb.Document, len(docs))
	for i, doc := range docs {
		documents[i] = &pb.Document{
			Content: doc.Content,
			Metadata: &pb.Metadata{Url: doc.Metadata.URL},
			Score: doc.Score,
		}
	}

	return &pb.QueryResponse{Document: documents}, err
}