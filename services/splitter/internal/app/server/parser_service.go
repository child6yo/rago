package server

import (
	"context"

	pb "github.com/child6yo/rago/services/splitter/proto"
)

type SplitterServiceServer interface {
	HandleDocuments(context.Context, *pb.Document) (*pb.Empty, error)
	mustEmbedUnimplementedParserServiceServer()
}

type splitterService struct{
	pb.SplitterServiceServer
}

func (p *splitterService) HandleDocuments(context.Context, *pb.Document) (*pb.Empty, error) {
	return nil, nil
}
