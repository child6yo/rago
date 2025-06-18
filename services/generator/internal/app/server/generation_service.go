package server

import (
	"fmt"
	"log"

	"github.com/child6yo/rago/services/generator/internal/app/usecase"
	"github.com/child6yo/rago/services/generator/pkg/pb"
	"google.golang.org/grpc"
)

// GeneratorServiceServer определяет интерфейс gRPC-сервера генератора.
type GeneratorServiceServer interface {
	// Generate принимает запрос для генерации по контексту и открывает поток
	// токенов ответа.
	Generate(query *pb.Query, stream grpc.ServerStreamingServer[pb.ResponseChunk]) error
	mustEmbedUnimplementedGeneratorServiceServer()
}

// GenerationService имплементирует интерфейс GeneratorServiceServer.
type GenerationService struct {
	pb.GeneratorServiceServer
	service usecase.Generation
}

// Generate принимает запрос для генерации по контексту и открывает поток
// токенов ответа.
func (gs *GenerationService) Generate(query *pb.Query, stream grpc.ServerStreamingServer[pb.ResponseChunk]) error {
	ctx := stream.Context()
	chunks, err := gs.service.Generate(ctx, query.Query)
	if err != nil {
		return fmt.Errorf("generation service: failed to generate answer: %v", err)
	}

	log.Printf("INFO: generator service new query: %s, starting stream...", query)

	for c := range chunks {
		log.Print(c)
		err := stream.Send(&pb.ResponseChunk{Chunk: c})
		if err != nil {
			return fmt.Errorf("generation service: error sending message to stream: %v", err)
		}
	}

	return nil
}
