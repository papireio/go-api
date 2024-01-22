package clients

import (
	"github.com/papireio/go-api/internal/env"
	session "github.com/papireio/go-session-service/pkg/api/grpc"
	users "github.com/papireio/go-users-service/pkg/api/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Session session.GoSessionClient
	Users   users.GoUsersClient
}

func SetupGrpcClients(cfg *env.Config) *GrpcClients {
	sessionConn, _ := grpc.Dial(cfg.SessionServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	usersConn, _ := grpc.Dial(cfg.UsersServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return &GrpcClients{
		Session: session.NewGoSessionClient(sessionConn),
		Users:   users.NewGoUsersClient(usersConn),
	}
}
