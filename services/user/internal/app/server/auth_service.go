package server

import (
	"context"

	"github.com/child6yo/rago/services/auth/internal"
	"github.com/child6yo/rago/services/auth/internal/app/usecase"
	"github.com/child6yo/rago/services/auth/pkg/pb"
)

// AuthServiceServer определяет интерфейс gRPC сервера сервиса авторизации.
type AuthServiceServer interface {
	// Register принимает на вход схему пользователя,
	// при успехе возвращает соответствующий статус.
	// В обратном случае возвращает ошибку.
	Register(ctx context.Context, user *pb.User) (*pb.Empty, error)

	// Login принимает на вход схему пользователя,
	// при успехе возвращает авторизационный токен.
	// В обратном случае возвращает ошибку.
	Login(ctx context.Context, user *pb.User) (*pb.Token, error)

	// Auth принимает на вход токен, при успехе возвращает соответсвующий статус.
	// В обратном случае возвращает ошибку.
	Auth(ctx context.Context, token *pb.Token) (*pb.Empty, error)

	mustEmbedUnimplementedAuthServiceServer()
}

// AuthService имплементирует интерфейс AuthServiceServer.
type AuthService struct {
	pb.AuthServiceServer
	service usecase.Authorization
}

// Register принимает на вход схему пользователя,
// при успехе возвращает соответствующий статус.
// В обратном случае возвращает ошибку.
func (a *AuthService) Register(ctx context.Context, user *pb.User) (*pb.Empty, error) {
	err := a.service.Register(internal.User{
		Login:    user.Login,
		Password: user.Password,
	})

	return nil, err
}

// Login принимает на вход схему пользователя,
// при успехе возвращает авторизационный токен.
// В обратном случае возвращает ошибку.
func (a *AuthService) Login(ctx context.Context, user *pb.User) (*pb.Token, error) {
	token, err := a.service.Login(user.Login, user.Password)

	return &pb.Token{Token: token}, err
}

// Auth принимает на вход токен, при успехе возвращает соответсвующий статус.
// В обратном случае возвращает ошибку.
func (a *AuthService) Auth(ctx context.Context, token *pb.Token) (*pb.Empty, error) {
	err := a.service.Auth(token.Token)

	return nil, err
}
