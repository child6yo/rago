package server

import (
	"context"

	"github.com/child6yo/rago/services/auth/pkg/pb"
)

// AuthServiceServer определяет интерфейс gRPC сервера сервиса авторизации.
type AuthServiceServer interface {
	// Register принимает на вход схему пользователя,
	// при успехе возвращает соответствующий статус.
	// В обратном случае возвращает ошибку.
	Register(context.Context, *pb.User) (*pb.Empty, error)

	// Login принимает на вход схему пользователя,
	// при успехе возвращает авторизационный токен.
	// В обратном случае возвращает ошибку.
	Login(context.Context, *pb.User) (*pb.Token, error)

	// Auth принимает на вход токен, при успехе возвращает соответсвующий статус.
	// В обратном случае возвращает ошибку.
	Auth(context.Context, *pb.Token) (*pb.Empty, error)
	mustEmbedUnimplementedAuthServiceServer()
}
