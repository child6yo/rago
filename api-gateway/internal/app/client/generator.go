package client

import (
	"context"
	"fmt"
	"log"

	"github.com/child6yo/rago/api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Generator struct {
	generator pb.GeneratorServiceClient
	conn      *grpc.ClientConn

	host string
	port string
}

func newGeneratorClient(host string, port string) *Generator {
	return &Generator{host: host, port: port}
}

func (g *Generator) startGeneratoClient() {
	addr := fmt.Sprintf("%s:%s", g.host, g.port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print("failed to connect generator grpc server")
	}

	g.generator = pb.NewGeneratorServiceClient(conn)
	g.conn = conn
}

func (g *Generator) stopGeneratoClient() {
	if g.conn == nil {
		return
	}
	g.conn.Close()
}

func (g *Generator) Generate(ctx context.Context, query string) (<-chan string, error) {
	out := make(chan string)
	go func() {
		stream, err := g.generator.Generate(ctx, &pb.Query{Query: query})
		if err != nil {
			log.Print(err)
		}

		defer close(out)
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Print("stream err: ", err)
				break
			}
			out <- msg.Chunk
		}
	}()

	return out, nil
}