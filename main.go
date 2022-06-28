package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pascallin/gin-template/pubsub"
	app "github.com/pascallin/gin-template/server"
	"github.com/pascallin/gin-template/server/ws"

	// NOTE: import swagger docs
	_ "github.com/pascallin/gin-template/docs"
)

// @title Gin API
// @version 1.0
// @description A Gin server demo API

// @contact.name pascal_lin

// @host localhost:4000
// @BasePath /v1

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {

	if err := ws.Start(); err != nil {
		panic(err)
	}

	router := app.InitServer()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: router,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	go func() {
		pubsub.Listen()
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("Server exiting")
}
