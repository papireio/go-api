package home

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// swagger:model ExampleResponse
type ExampleResponse struct {
	Message string `json:"message"`
}

func Example(ctx context.Context) func(c *gin.Context) {
	// swagger:operation GET /home/example Home Example
	//
	// Example
	//
	// ---
	// produces:
	// - application/json
	//
	// responses:
	//   '200':
	//     description: successful
	//     schema:
	//       "$ref": "#/definitions/ExampleResponse"
	//   '500':
	//     description: Server Error
	return func(c *gin.Context) {
		response := &ExampleResponse{}

		c.JSON(http.StatusOK, response)
	}
}
