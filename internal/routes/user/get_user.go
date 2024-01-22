package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/papireio/go-api/internal/clients"
	"net/http"
)

func GetUser(ctx context.Context, grpcClients *clients.GrpcClients) func(c *gin.Context) {
	return func(c *gin.Context) {
		u, ok := c.Get("user")
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.JSON(http.StatusOK, u)
	}
}
