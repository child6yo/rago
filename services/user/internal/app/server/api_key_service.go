package server

import (
	"context"

	"github.com/child6yo/rago/services/user/internal/app/usecase"
	"github.com/child6yo/rago/services/user/pkg/pb"
)

type APIKeyServiceServer interface {
	CreateAPIKey(ctx context.Context, userID *pb.UserID) (*pb.APIKey, error)
	DeleteAPIKey(ctx context.Context, request *pb.DeleteAPIKeyRequest) (*pb.Empty, error)
	GetAPIKeys(ctx context.Context, userID *pb.UserID) (*pb.APIKeyArray, error)
	CheckAPIKey(ctx context.Context, key *pb.APIKey) (*pb.Empty, error)
	mustEmbedUnimplementedApiKeyServiceServer()
}

// APIKeyService имплементирует интерфейс ApiKeyServiceServer.
type APIKeyService struct {
	pb.APIKeyServiceServer
	service usecase.ApiKey
}

func (aks *APIKeyService) CreateAPIKey(ctx context.Context, userID *pb.UserID) (*pb.APIKey, error) {
	key, err := aks.service.CreateApiKey(int(userID.Id))
	if err != nil {
		return nil, err
	}

	return &pb.APIKey{Key: key}, nil
}

func (aks *APIKeyService) DeleteAPIKey(ctx context.Context, request *pb.DeleteAPIKeyRequest) (*pb.Empty, error) {
	err := aks.service.DeleteApiKey(int(request.ApiKeyId), int(request.UserId.Id))

	return nil, err
}

func (aks *APIKeyService) GetAPIKeys(ctx context.Context, userID *pb.UserID) (*pb.APIKeyArray, error) {
	internalKeys, err := aks.service.GetApiKeys(int(userID.Id))
	if err != nil {
		return nil, err
	}

	// преобразование []internal.ApiKey -> []*pb.APIKey
	keys := make([]*pb.APIKey, len(internalKeys))
	for i, k := range internalKeys {
		keys[i] = &pb.APIKey{
			Key: k.Key,
		}
	}

	return &pb.APIKeyArray{Keys: keys}, nil
}

func (aks *APIKeyService) CheckAPIKey(ctx context.Context, key *pb.APIKey) (*pb.Empty, error) {
	err := aks.service.CheckAPIKey(key.Key)

	return nil, err
}
