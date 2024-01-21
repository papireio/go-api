// Package goapi
//
// Goapi endpoints
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/isalikov/goapi/internal/env"
	home "github.com/isalikov/goapi/internal/routes"
	"github.com/sethvargo/go-envconfig"
	"log"
)

var ctx = context.Background()

func main() {
	config := &env.Config{}

	if err := envconfig.Process(ctx, config); err != nil {
		log.Fatalln(err, "Fatal Error: Parsing OS ENV")
	}

	r := gin.Default()
	r.Static("/docs", "swagger-ui")

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	r.GET("/home/example", home.Example(ctx))

	if err := r.Run(fmt.Sprintf("0.0.0.0:%v", config.Port)); err != nil {
		log.Fatalln(err, "Fatal Error: Running Server")
	}
}
