package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/pascallin/gin-template/pubsub"
	app "github.com/pascallin/gin-template/server"
	"github.com/sirupsen/logrus"

	// NOTE: import swagger docs
	_ "github.com/pascallin/gin-template/docs"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := app.Start(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	go func() {
		pubsub.Listen()
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	logrus.Println("Server exiting")
}
