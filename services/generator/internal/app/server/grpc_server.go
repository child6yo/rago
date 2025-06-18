package server

import (
	"fmt"
	"net"

	"github.com/child6yo/rago/services/generator/internal/app/usecase"
	pb "github.com/child6yo/rago/proto/generator"
	"google.golang.org/grpc"
)

// GRPCServer - структура gRPC-сервера приложения.
type GRPCServer struct {
	server *grpc.Server
	usecase usecase.Generation

	host string
	port string
}

// NewGRPCServer создает новый экземпляр GRPCServer.
func NewGRPCServer(usecase usecase.Generation, host string, port string) *GRPCServer {
	return &GRPCServer{
		server: grpc.NewServer(),
		usecase: usecase,
		host:   host,
		port:   port,
	}
}

// StartGRPCServer запускает сервер в соответствии с указанными в нем параметрами.
func (g *GRPCServer) StartGRPCServer() error {
	addr := fmt.Sprintf("%s:%s", g.host, g.port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	pb.RegisterGeneratorServiceServer(g.server, &GenerationService{service: g.usecase})

	if err := g.server.Serve(lis); err != nil {
		return err
	}
	return nil
}

// ShutdownGRPCServer останавливает работу сервера.
func (g *GRPCServer) ShutdownGRPCServer() {
	g.server.GracefulStop()
}
