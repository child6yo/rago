package server

import (
	"fmt"
	"net"

	pb "github.com/child6yo/rago/services/splitter/proto"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	server *grpc.Server
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{server: grpc.NewServer()}
}

func (g *GRPCServer) StartGRPCServer(host, port string) error {
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	pb.RegisterSplitterServiceServer(g.server, new(splitterService))

	if err := g.server.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (g *GRPCServer) ShutdownGRPCServer() {
	g.server.GracefulStop()
}
