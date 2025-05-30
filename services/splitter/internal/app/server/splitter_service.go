package server

import (
	"context"

	"github.com/child6yo/rago/services/splitter/internal/app/usecase"
	pb "github.com/child6yo/rago/services/splitter/proto"
)

type SplitterServiceServer interface {
	HandleDocuments(ctx context.Context, docs *pb.DocumentArray) (*pb.Empty, error)
	mustEmbedUnimplementedSplitterServiceServer()
}

type splitterService struct {
	pb.SplitterServiceServer
	splitter usecase.Splitter
}

func (s *splitterService) HandleDocuments(ctx context.Context, docs *pb.DocumentArray) (*pb.Empty, error) {
	s.splitter.SplitDocuments(docs.Doc)
	return nil, nil
}
