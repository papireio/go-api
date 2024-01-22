package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/papireio/go-api/internal/clients"
	"github.com/papireio/go-api/internal/env"
	"github.com/papireio/go-api/internal/routes/authentication"
	"github.com/sethvargo/go-envconfig"
	"log"
)

var ctx = context.Background()

func main() {
	config := &env.Config{}

	if err := envconfig.Process(context.Background(), config); err != nil {
		log.Fatalln(err, "Fatal Error: Parsing OS ENV")
	}

	grpcClients := clients.SetupGrpcClients(config)

	r := gin.Default()
	r.Static("/docs", "docs")

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	r.POST("/sign/in", authentication.SignIn(ctx, grpcClients))

	if err := r.Run(fmt.Sprintf("0.0.0.0:%v", config.Port)); err != nil {
		log.Fatalln(err, "Fatal Error: Running Server")
	}
}
