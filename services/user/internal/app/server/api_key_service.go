package server

import (
	"context"
	"log"

	pb "github.com/child6yo/rago/proto/user"
	"github.com/child6yo/rago/services/user/internal/app/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
)

// APIKeyServiceServer определяет интерфейс gRPC сервера сервиса API ключей.
type APIKeyServiceServer interface {
	// CreateAPIKey принимает на вход айди пользователя,
	// возвращает его новый ключ API.
	CreateAPIKey(ctx context.Context, request *pb.UserID) (*pb.APIKey, error)

	// DeleteAPIKey принимает запрос в виде айди пользователя и айди API ключа,
	// отправляет запрос на удаление API ключа.
	DeleteAPIKey(ctx context.Context, request *pb.DeleteAPIKeyRequest) (*emptypb.Empty, error)

	// GetAPIKeys принимает айди пользователя и возвращает все его API ключи.
	GetAPIKeys(ctx context.Context, request *pb.UserID) (*pb.APIKeyArray, error)

	// CheckAPIKey принимает ключ API и проверяет его существование в базе данных.
	CheckAPIKey(ctx context.Context, request *pb.APIKey) (*emptypb.Empty, error)
	mustEmbedUnimplementedAPIKeyServiceServer()
}

// APIKeyService имплементирует интерфейс APIKeyServiceServer.
type APIKeyService struct {
	pb.APIKeyServiceServer
	service usecase.APIKey
}

// CreateAPIKey принимает на вход айди пользователя,
// возвращает его новый ключ API.
func (aks *APIKeyService) CreateAPIKey(ctx context.Context, userID *pb.UserID) (*pb.APIKey, error) {
	key, err := aks.service.CreateAPIKey(int(userID.Id))
	if err != nil {
		log.Printf("user service error: %v", err)
		return nil, err
	}

	return &pb.APIKey{Id: key.ID, Key: key.Key}, nil
}

// DeleteAPIKey принимает запрос в виде айди пользователя и айди API ключа,
// отправляет запрос на удаление API ключа.
func (aks *APIKeyService) DeleteAPIKey(ctx context.Context, request *pb.DeleteAPIKeyRequest) (*emptypb.Empty, error) {
	err := aks.service.DeleteAPIKey(request.ApiKeyId, int(request.UserId.Id))
	if err != nil {
		log.Printf("user service error: %v", err)
		return nil, err
	}

	return nil, nil
}

// GetAPIKeys принимает айди пользователя и возвращает все его API ключи.
func (aks *APIKeyService) GetAPIKeys(ctx context.Context, userID *pb.UserID) (*pb.APIKeyArray, error) {
	internalKeys, err := aks.service.GetAPIKeys(int(userID.Id))
	if err != nil {
		log.Printf("user service error: %v", err)
		return nil, err
	}

	keys := make([]*pb.APIKey, len(internalKeys))
	for i, k := range internalKeys {
		keys[i] = &pb.APIKey{
			Id:  k.ID,
			Key: k.Key,
		}
	}

	return &pb.APIKeyArray{Keys: keys}, nil
}

// CheckAPIKey принимает ключ API и проверяет его существование в базе данных.
func (aks *APIKeyService) CheckAPIKey(ctx context.Context, key *pb.APIKey) (*emptypb.Empty, error) {
	err := aks.service.CheckAPIKey(key.Key)
	if err != nil {
		log.Printf("user service error: %v", err)
		return nil, err
	}

	return nil, nil
}
