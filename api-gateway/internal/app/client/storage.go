package client

import (
	"context"
	"fmt"
	"log"

	"github.com/child6yo/rago/api-gateway/internal"
	"github.com/child6yo/rago/api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Storage struct {
	splitter pb.SplitterServiceClient
	conn     *grpc.ClientConn

	host string
	port string
}

func newStorageClient(host string, port string) *Storage {
	return &Storage{host: host, port: port}
}

func (sc *Storage) startStorageClient() {
	addr := fmt.Sprintf("%s:%s", sc.host, sc.port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print("failed to connect splitter grpc server")
	}

	sc.splitter = pb.NewSplitterServiceClient(conn)
	sc.conn = conn
}

func (sc *Storage) stopStorageClient() {
	if sc.conn == nil {
		return
	}
	sc.conn.Close()
}

func (sc *Storage) LoadDocuments(docs []internal.Document) error {
	pbDocs := make([]*pb.Document, 0, len(docs))

	for _, doc := range docs {
		pbDocs = append(pbDocs, &pb.Document{
			Content:  doc.Content,
			Metadata: &pb.Metadata{Url: doc.Metadata.URL},
		})
	}

	pbDocArray := &pb.DocumentArray{
		Doc: pbDocs,
	}

	_, err := sc.splitter.HandleDocuments(context.Background(), pbDocArray)
	log.Print(err)

	return err
}
