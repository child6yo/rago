package server

import (
	"fmt"
	"net"

	"github.com/child6yo/rago/services/splitter/internal/app/usecase"
	pb "github.com/child6yo/rago/services/splitter/proto"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	server   *grpc.Server
	splitter usecase.Splitter
	
	host string
	port string
}

func NewGRPCServer(splitter usecase.Splitter, host string, port string) *GRPCServer {
	return &GRPCServer{
		server: grpc.NewServer(),
		splitter: splitter,
		host: host,
		port: port,
	}
}

func (g *GRPCServer) StartGRPCServer() error {
	addr := fmt.Sprintf("%s:%s", g.host, g.port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	pb.RegisterSplitterServiceServer(g.server, &splitterService{splitter: g.splitter})

	if err := g.server.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (g *GRPCServer) ShutdownGRPCServer() {
	g.server.GracefulStop()
}
