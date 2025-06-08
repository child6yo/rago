package client

import (
	"context"
	"fmt"
	"log"

	"github.com/child6yo/rago/api-gateway/internal"
	pb "github.com/child6yo/rago/api-gateway/pkg/pb/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// User определяет клиент пользовательского сервиса, доступного по gRPC.
type User struct {
	auth   pb.AuthServiceClient
	apiKey pb.APIKeyServiceClient
	conn   *grpc.ClientConn

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
func (uc *User) Register(input internal.User) error {
	_, err := uc.auth.Register(context.Background(), &pb.User{
		Login:    input.Login,
		Password: input.Password,
	})

	return err
}

// Register вызывает удалённый метод логина пользователя через gRPC.
func (uc *User) Login(input internal.User) (string, error) {
	token, err := uc.auth.Login(context.Background(), &pb.User{
		Login:    input.Login,
		Password: input.Password,
	})
	if err != nil {
		return "", err
	}

	return token.Token, nil
}

// Register вызывает удалённый метод авторизации пользователя через gRPC.
// Возвращает айди пользователя и ошибку.
func (uc *User) Auth(token string) (int, error) {
	id, err := uc.auth.Auth(context.Background(), &pb.Token{
		Token: token,
	})

	return int(id.Id), err
}

func (uc *User) CreateAPIKey(userID int) (string, error) {
	key, err := uc.apiKey.CreateAPIKey(context.Background(), &pb.UserID{Id: int32(userID)})
	if err != nil {
		return "", err
	}

	return key.Key, nil
}

func (uc *User) DeleteAPIKey(keyID, userID int) error {
	_, err := uc.apiKey.DeleteAPIKey(context.Background(), &pb.DeleteAPIKeyRequest{
		ApiKeyId: int32(keyID),
		UserId:   &pb.UserID{Id: int32(userID)},
	})

	return err
}

func (uc *User) GetAPIKeys(userID int) ([]internal.ApiKey, error) {
	keys, err := uc.apiKey.GetAPIKeys(context.Background(), &pb.UserID{Id: int32(userID)})
	if err != nil {
		return []internal.ApiKey{}, err
	}

	internalKeys := make([]internal.ApiKey, len(keys.Keys))
	for i, k := range keys.Keys {
		internalKeys[i] = internal.ApiKey{
			Key: k.Key,
		}
	}

	return internalKeys, nil
}

func (uc *User) CheckAPIKey(key string) error {
	_, err := uc.apiKey.CheckAPIKey(context.Background(), &pb.APIKey{Key: key})

	return err
}
