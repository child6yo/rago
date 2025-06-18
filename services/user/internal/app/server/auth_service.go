package server

import (
	"context"
	"log"

	"github.com/child6yo/rago/services/user/internal"
	"github.com/child6yo/rago/services/user/internal/app/usecase"
	"github.com/child6yo/rago/services/user/pkg/pb"
)

// AuthServiceServer определяет интерфейс gRPC сервера сервиса авторизации.
type AuthServiceServer interface {
	// Register принимает на вход схему пользователя,
	// возвращает коллекцию, принадлежающую пользователю.
	Register(ctx context.Context, user *pb.User) (*pb.Collection, error)

	// Login принимает на вход схему пользователя,
	// при успехе возвращает авторизационный токен.
	Login(ctx context.Context, user *pb.User) (*pb.Token, error)

	// Auth принимает на вход токен, при успехе возвращает соответсвующий статус.
	Auth(ctx context.Context, token *pb.Token) (*pb.UserID, error)

	mustEmbedUnimplementedAuthServiceServer()
}

// AuthService имплементирует интерфейс AuthServiceServer.
type AuthService struct {
	pb.AuthServiceServer
	service usecase.Authorization
}

// Register принимает на вход схему пользователя,
// возвращает коллекцию, принадлежающую пользователю.
func (as *AuthService) Register(ctx context.Context, user *pb.User) (*pb.Collection, error) {
	collection, err := as.service.Register(internal.User{
		Login:    user.Login,
		Password: user.Password,
	})

	if err != nil {
		log.Printf("user service error: %v", err)
		return nil, err
	}

	return &pb.Collection{Collection: collection}, nil
}

// Login принимает на вход схему пользователя,
// возвращает авторизационный токен.
func (as *AuthService) Login(ctx context.Context, user *pb.User) (*pb.Token, error) {
	token, err := as.service.Login(user.Login, user.Password)
	if err != nil {
		log.Printf("user service error: %v", err)
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}

// Auth принимает на вход токен, возвращает айди пользователя.
func (as *AuthService) Auth(ctx context.Context, token *pb.Token) (*pb.UserID, error) {
	userID, err := as.service.Auth(token.Token)
	if err != nil {
		log.Printf("user service error: %v", err)
		return nil, err
	}

	return &pb.UserID{Id: int32(userID)}, nil
}
