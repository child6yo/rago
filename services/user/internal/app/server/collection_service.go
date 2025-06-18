package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/child6yo/rago/proto/user"
	"github.com/child6yo/rago/services/user/internal/app/usecase"
)

// CollectionServiceServer определяет интерфейс gRPC сервера сервиса коллекций.
type CollectionServiceServer interface {
	// GetCollection принимает айди пользователя и возвращает коллецию,
	// которая ему принадлежит.
	GetCollection(ctx context.Context, request *pb.UserID) (*pb.Collection, error)
	mustEmbedUnimplementedCollectionServiceServer()
}

// CollectionService имплементирует интерфейс GetCollection.
type CollectionService struct {
	pb.CollectionServiceServer
	service usecase.Collection
}

// GetCollection принимает айди пользователя и возвращает коллецию,
// которая ему принадлежит.
func (cs *CollectionService) GetCollection(ctx context.Context, request *pb.UserID) (*pb.Collection, error) {
	if request == nil {
		log.Printf("user service error: empty request")
		return nil, fmt.Errorf("get collection: empty request")
	}

	collection, err := cs.service.GetCollection(int(request.Id))
	if err != nil {
		log.Printf("user service error: %v", err)
		return nil, err
	}

	return &pb.Collection{Collection: collection}, nil
}
