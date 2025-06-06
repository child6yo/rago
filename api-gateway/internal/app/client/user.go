package client

import (
	"context"
	"fmt"
	"log"

	"github.com/child6yo/rago/api-gateway/internal"
	pb "github.com/child6yo/rago/api-gateway/pkg/pb/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// User определяет клиент пользовательского сервиса, доступного по gRPC.
type User struct {
	usrClient pb.AuthServiceClient
	usrConn   *grpc.ClientConn

	usrHost string
	usrPort string
}

func newUserClient(host string, port string) *User {
	return &User{usrHost: host, usrPort: port}
}

func (uc *User) startUserClient() {
	addr := fmt.Sprintf("%s:%s", uc.usrHost, uc.usrPort)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print("failed to connect user grpc server")
	}

	uc.usrClient = pb.NewAuthServiceClient(conn)
	uc.usrConn = conn
}

func (uc *User) stopUserClient() {
	if uc.usrConn == nil {
		return
	}
	uc.usrConn.Close()
}

// Register вызывает удалённый метод регистрации пользователя через gRPC.
func (uc *User) Register(input internal.User) error {
	_, err := uc.usrClient.Register(context.Background(), &pb.User{
		Login:    input.Login,
		Password: input.Password,
	})

	return err
}
