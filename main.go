package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pascallin/gin-template/pkg"
	"github.com/pascallin/gin-template/pubsub"
	app "github.com/pascallin/gin-template/server"
	"github.com/sirupsen/logrus"

	// NOTE: import swagger docs
	_ "github.com/pascallin/gin-template/docs"
)

func init() {
	pkg.SetupLogger()
}

func main() {
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
