package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/papireio/go-api/internal/clients"
	"github.com/papireio/go-api/internal/models"
	session "github.com/papireio/go-session-service/pkg/api/grpc"
	users "github.com/papireio/go-users-service/pkg/api/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func AuthMiddleware(ctx context.Context, grpcClients *clients.GrpcClients) func(c *gin.Context) {
	return func(c *gin.Context) {
		sessionToken := c.GetHeader("session_token")
		if sessionToken == "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		r, err := grpcClients.Session.ExtractSession(ctx, &session.ExtractSessionRequest{SessionToken: sessionToken})
		if e, ok := status.FromError(err); ok && err != nil {
			switch e.Code() {
			case codes.InvalidArgument:
				c.AbortWithStatus(http.StatusBadRequest)
			case codes.NotFound:
				c.AbortWithStatus(http.StatusUnauthorized)
			default:
				// TODO: log error
				c.AbortWithStatus(http.StatusInternalServerError)
			}

			return
		}

		u, err := grpcClients.Users.GetUser(ctx, &users.GetUserRequest{Uuid: r.Uuid})
		if e, ok := status.FromError(err); ok && err != nil {
			switch e.Code() {
			case codes.InvalidArgument:
				c.AbortWithStatus(http.StatusBadRequest)
			case codes.NotFound:
				c.AbortWithStatus(http.StatusUnauthorized)
			default:
				// TODO: log error
				c.AbortWithStatus(http.StatusInternalServerError)
			}

			return
		}

		c.Set("user", &models.User{
			Name:     u.Name,
			Email:    u.Email,
			Uuid:     u.Uuid,
			Verified: u.Verified,
		})

		c.Next()
	}
}
