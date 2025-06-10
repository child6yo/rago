package server

import (
	"context"
	"fmt"
	"log"

	"github.com/child6yo/rago/services/generator/internal/app/usecase"
	"github.com/child6yo/rago/services/generator/pkg/pb"
	"google.golang.org/grpc"
)

type GeneratorServiceServer interface {
	Generate(query *pb.Query, stream grpc.ServerStreamingServer[pb.ResponseChunk]) error
	mustEmbedUnimplementedGeneratorServiceServer()
}

type GenerationService struct {
	pb.GeneratorServiceServer
	service usecase.Generation
}

func (gs *GenerationService) Generate(query *pb.Query, stream grpc.ServerStreamingServer[pb.ResponseChunk]) error {
	ctx := context.Background()
	chunks, err := gs.service.Generate(ctx, query.Query)
	if err != nil {
		log.Print(err)
		return err
	}
	
	for c := range chunks {
		err := stream.Send(&pb.ResponseChunk{Chunk: c})
		if err != nil {
			return fmt.Errorf("error sending message to stream: %v", err)
		}
	}

	return nil
}
