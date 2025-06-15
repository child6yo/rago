package server

import (
	"context"
	"log"

	"github.com/child6yo/rago/services/storage/internal/app/usecase"
	"github.com/child6yo/rago/services/storage/pkg/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// StorageServiceServer определяет интерфейс gRPC сервера сервиса векторного хранилища.
type StorageServiceServer interface {
	// CreateCollection реализует удаленный метод создания коллекции.
	CreateCollection(ctx context.Context, req *pb.CollectionRequest) (*emptypb.Empty, error)

	// DeleteCollection реализует удаленный метод удаления коллекции.
	DeleteCollection(ctx context.Context, req *pb.CollectionRequest) (*emptypb.Empty, error)

	// DeleteDocument реализует удаленный метод удаления документа из коллекции.
	DeleteDocument(ctx context.Context, req *pb.DocumentRequest) (*emptypb.Empty, error)

	// GetDocument реализует удаленный метод выдачи документа по айди.
	GetDocument(ctx context.Context, req *pb.DocumentRequest) (*pb.Document, error)

	// GetAllDocuments реализует удаленный метод выдачи всех документов коллекции.
	GetAllDocuments(ctx context.Context, req *pb.CollectionRequest) (*pb.DocumentArray, error)

	// Search реализует удаленный метод векторного поиска по коллекции.
	Search(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error)
	mustEmbedUnimplementedStorageServiceServer()
}

// StorageService имплементирует интерфейс StorageService.
type StorageService struct {
	pb.StorageServiceServer
	service usecase.Storage
}

// CreateCollection реализует удаленный метод создания коллекции.
func (s *StorageService) CreateCollection(ctx context.Context, req *pb.CollectionRequest) (*emptypb.Empty, error) {
	err := s.service.CreateCollection(ctx, req.CollectionName)
	if err != nil {
		log.Printf("storage service error (create collection): %v", err)
	} else {
		log.Printf("storage: new collection created: %s", req.CollectionName)
	}

	return nil, err
}

// DeleteCollection реализует удаленный метод удаления коллекции.
func (s *StorageService) DeleteCollection(ctx context.Context, req *pb.CollectionRequest) (*emptypb.Empty, error) {
	err := s.service.DeleteCollection(ctx, req.CollectionName)
	if err != nil {
		log.Printf("storage service error (delete collection): %v", err)
	} else {
		log.Printf("storage: collection deleted: %s", req.CollectionName)
	}

	return nil, err
}

// DeleteDocument реализует удаленный метод удаления документа из коллекции.
func (s *StorageService) DeleteDocument(ctx context.Context, req *pb.DocumentRequest) (*emptypb.Empty, error) {
	err := s.service.DeleteDocument(ctx, req.Id, req.CollectionName)
	if err != nil {
		log.Printf("storage service error (delete document): %v", err)
	}

	return nil, err
}

// GetDocument реализует удаленный метод выдачи документа по айди.
func (s *StorageService) GetDocument(ctx context.Context, req *pb.DocumentRequest) (*pb.Document, error) {
	doc, err := s.service.GetDocument(ctx, req.CollectionName, req.Id)
	if err != nil {
		log.Printf("storage service error (get document): %v", err)
		return nil, err
	}

	return &pb.Document{
		Content:  doc.Content,
		Id:       doc.ID,
		Metadata: &pb.Metadata{Url: doc.Metadata.URL},
	}, nil
}

// GetAllDocuments реализует удаленный метод выдачи всех документов коллекции.
func (s *StorageService) GetAllDocuments(ctx context.Context, req *pb.CollectionRequest) (*pb.DocumentArray, error) {
	docs, err := s.service.GetAllDocuments(ctx, req.CollectionName)
	if err != nil {
		log.Printf("storage service error (get all documents): %v", err)
		return nil, err
	}

	response := make([]*pb.Document, len(docs))
	for i, doc := range docs {
		response[i] = &pb.Document{
			Content:  doc.Content,
			Id:       doc.ID,
			Metadata: &pb.Metadata{Url: doc.Metadata.URL},
		}
	}

	return &pb.DocumentArray{Document: response}, nil
}

// Search реализует удаленный метод векторного поиска по коллекции.
func (s *StorageService) Search(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	docs, err := s.service.Search(ctx, req.CollectionName, req.Query, int(req.Limit))

	documents := make([]*pb.Document, len(docs))
	for i, doc := range docs {
		documents[i] = &pb.Document{
			Content:  doc.Content,
			Metadata: &pb.Metadata{Url: doc.Metadata.URL},
			Score:    doc.Score,
		}
	}

	return &pb.QueryResponse{Document: documents}, err
}
