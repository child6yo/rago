package client

import (
	"context"
	"fmt"
	"log"

	"github.com/child6yo/rago/api-gateway/internal"
	"github.com/child6yo/rago/services/user/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// User определяет клиент пользовательского сервиса, доступного по gRPC.
type User struct {
	auth       pb.AuthServiceClient
	apiKey     pb.APIKeyServiceClient
	collection pb.CollectionServiceClient
	conn       *grpc.ClientConn

	host string
	port string
}

func newUserClient(host string, port string) *User {
	return &User{host: host, port: port}
}

func (uc *User) startUserClient() {
	addr := fmt.Sprintf("%s:%s", uc.host, uc.port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print("failed to connect user grpc server")
	}

	uc.auth = pb.NewAuthServiceClient(conn)
	uc.apiKey = pb.NewAPIKeyServiceClient(conn)
	uc.conn = conn
}

func (uc *User) stopUserClient() {
	if uc.conn == nil {
		return
	}
	uc.conn.Close()
}

// Register вызывает удалённый метод регистрации пользователя через gRPC.
func (uc *User) Register(ctx context.Context, input internal.User) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("user client error (Register): %v", err)
		}
	}()

	collection, err := uc.auth.Register(ctx, &pb.User{
		Login:    input.Login,
		Password: input.Password,
	})

	return collection.Collection, err
}

// Login вызывает удалённый метод логина пользователя через gRPC.
func (uc *User) Login(ctx context.Context, input internal.User) (string, error) {
	token, err := uc.auth.Login(ctx, &pb.User{
		Login:    input.Login,
		Password: input.Password,
	})
	if err != nil {
		return "", fmt.Errorf("user client error (Login): %v", err)
	}

	return token.Token, nil
}

// Auth вызывает удалённый метод авторизации пользователя через gRPC.
// Возвращает айди пользователя и ошибку.
func (uc *User) Auth(ctx context.Context, token string) (_ int, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("user client error (Auth): %v", err)
		}
	}()

	id, err := uc.auth.Auth(ctx, &pb.Token{
		Token: token,
	})

	return int(id.Id), err
}

// CreateAPIKey вызывает удалённый метод создания API ключа.
func (uc *User) CreateAPIKey(ctx context.Context, userID int) (string, error) {
	key, err := uc.apiKey.CreateAPIKey(ctx, &pb.UserID{Id: int32(userID)})
	if err != nil {
		return "", fmt.Errorf("user client error (CreateAPIKey): %v", err)
	}

	return key.Key, nil
}

// DeleteAPIKey вызывает удалённый метод удаления API ключа.
func (uc *User) DeleteAPIKey(ctx context.Context, keyID, userID int) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("user client error (DeleteAPIKey): %v", err)
		}
	}()

	_, err = uc.apiKey.DeleteAPIKey(ctx, &pb.DeleteAPIKeyRequest{
		ApiKeyId: int32(keyID),
		UserId:   &pb.UserID{Id: int32(userID)},
	})

	return err
}

// GetAPIKeys вызывает удалённый метод возврата всех API ключей пользователя.
func (uc *User) GetAPIKeys(ctx context.Context, userID int) ([]internal.APIKey, error) {
	keys, err := uc.apiKey.GetAPIKeys(ctx, &pb.UserID{Id: int32(userID)})
	if err != nil {
		return []internal.APIKey{}, fmt.Errorf("user client error (GetAPIKeys): %v", err)
	}

	internalKeys := make([]internal.APIKey, len(keys.Keys))
	for i, k := range keys.Keys {
		internalKeys[i] = internal.APIKey{
			Key: k.Key,
		}
	}

	return internalKeys, nil
}

// CheckAPIKey вызывает удаленный метод валидации API ключа.
func (uc *User) CheckAPIKey(ctx context.Context, key string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("user client error (CheckAPIKey): %v", err)
		}
	}()

	_, err = uc.apiKey.CheckAPIKey(ctx, &pb.APIKey{Key: key})

	return err
}

// GetCollection вызывает удаленный метод возврата коллекции пользователя.
func (uc *User) GetCollection(ctx context.Context, userID int) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("user client error (GetCollection): %v", err)
		}
	}()

	collection, err := uc.collection.GetCollection(ctx, &pb.UserID{Id: int32(userID)})
	return collection.Collection, err
}
