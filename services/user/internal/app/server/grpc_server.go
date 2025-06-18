package server

import (
	"fmt"
	"net"

	"github.com/child6yo/rago/services/user/internal/app/usecase"
	pb "github.com/child6yo/rago/proto/user"
	"google.golang.org/grpc"
)

// GRPCServer - структура gRPC-сервера приложения.
type GRPCServer struct {
	server  *grpc.Server
	usecase *usecase.Usecase

	host string
	port string
}

// NewGRPCServer создает новый экземпляр GRPCServer.
func NewGRPCServer(usecase *usecase.Usecase, host string, port string) *GRPCServer {
	return &GRPCServer{
		server:  grpc.NewServer(),
		usecase: usecase,
		host:    host,
		port:    port,
	}
}

// StartGRPCServer запускает сервер в соответствии с указанными в нем параметрами.
func (g *GRPCServer) StartGRPCServer() error {
	addr := fmt.Sprintf("%s:%s", g.host, g.port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	pb.RegisterAuthServiceServer(g.server, &AuthService{service: g.usecase.Authorization})
	pb.RegisterAPIKeyServiceServer(g.server, &APIKeyService{service: g.usecase.APIKey})
	pb.RegisterCollectionServiceServer(g.server, &CollectionService{service: g.usecase.Collection})

	if err := g.server.Serve(lis); err != nil {
		return err
	}
	return nil
}

// ShutdownGRPCServer останавливает работу сервера.
func (g *GRPCServer) ShutdownGRPCServer() {
	g.server.GracefulStop()
}
