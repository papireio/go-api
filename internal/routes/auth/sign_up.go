package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/papireio/go-api/internal/clients"
	"github.com/papireio/go-api/internal/validation"
	session "github.com/papireio/go-session-service/pkg/api/grpc"
	users "github.com/papireio/go-users-service/pkg/api/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type SignUpRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=7"`
}

type SignUpResponse struct {
	SessionToken string `json:"session_token"`
}

func SignUp(ctx context.Context, grpcClients *clients.GrpcClients) func(c *gin.Context) {
	return func(c *gin.Context) {
		var b SignInRequestBody
		if err := c.ShouldBindJSON(&b); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": validation.GetErrors(err)})
			return
		}

		r, err := grpcClients.Users.CreateUser(ctx, &users.CreateUserRequest{
			Email:    b.Email,
			Password: b.Password,
		})

		if e, ok := status.FromError(err); ok && err != nil {
			switch e.Code() {
			case codes.InvalidArgument:
				c.AbortWithStatus(http.StatusBadRequest)
			case codes.AlreadyExists:
				c.AbortWithStatus(http.StatusConflict)
			default:
				// TODO: log error
				c.AbortWithStatus(http.StatusInternalServerError)
			}

			return
		}

		if _, err := grpcClients.Session.CreateSession(ctx, &session.CreateSessionRequest{
			SessionToken: r.SessionToken,
			Uuid:         r.Uuid,
		}); err != nil {
			//	TODO: remove recently created user
			c.AbortWithStatus(http.StatusInternalServerError)

			return
		}

		c.JSON(http.StatusOK, &SignInResponse{
			SessionToken: r.SessionToken,
		})
	}
}
