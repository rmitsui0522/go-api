package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"go-api/pkg/connection"
	"go-api/pkg/handler"
	"go-api/pkg/logger"
)

func main() {

	server := &http.Server{
		Addr:    connection.Port(),
		Handler: handler.New(),
	}

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)

		<-stop
		logger.Info("Shutting down...")

		if err := server.Shutdown(context.Background()); err != nil {
			logger.Error("server.Shutdown() failed: " + err.Error())
		}

		logger.Info("Server stopped.")
	}()

	logger.Info("server listening on http://localhost" + connection.Port())

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal("server.ListenAndServe() failed: " + err.Error())
	}
}
