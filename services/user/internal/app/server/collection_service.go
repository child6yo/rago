package server

import (
	"context"
	"log"

	"github.com/child6yo/rago/services/user/internal/app/usecase"
	"github.com/child6yo/rago/services/user/pkg/pb"
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
	collection, err := cs.service.GetCollection(int(request.Id))
	if err != nil {
		log.Printf("user service error: %v", err)
		return nil, err
	}

	return &pb.Collection{Collection: collection}, err
}
