package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/child6yo/rago/services/splitter/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	gRPChost := "localhost"
	gRPCport := "5000"
	addr := fmt.Sprintf("%s:%s", gRPChost, gRPCport)

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("could not connect to grpc server: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	grpcClient := pb.NewSplitterServiceClient(conn)

	grpcClient.HandleDocuments(context.Background(), &pb.DocumentArray{
		Doc: []*pb.Document{
			{
				Content: "TLS — это критически важный элемент современного интернета.",
				Metadata: &pb.Metadata{
					Url: "test.com",
				},
			},
			{
				Content: "В следующей главе мы рассмотрим, как работает HTTP/1.x и какие ограничения он имеет, особенно в контексте производительности и использования TCP.",
				Metadata: &pb.Metadata{
					Url: "test.com",
				},
			},
		},
	})
}
