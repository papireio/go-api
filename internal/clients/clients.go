package clients

import (
	"github.com/papireio/go-api/internal/env"
	users "github.com/papireio/go-users-service/pkg/api/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Users users.GoUsersClient
}

func SetupGrpcClients(cfg *env.Config) *GrpcClients {
	usersConn, _ := grpc.Dial(cfg.UsersServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return &GrpcClients{Users: users.NewGoUsersClient(usersConn)}
}
