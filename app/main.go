package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	conn "go-api/pkg/connection"
	"go-api/pkg/handler"
	"go-api/pkg/logger"
	"go-api/pkg/model"
)

func main() {

	server := &http.Server{
		Addr:    conn.Port(),
		Handler: handler.New(),
	}

	defer model.DB.Close()

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

	logger.Info("server listening on http://localhost" + conn.Port())
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal("server.ListenAndServe() failed: " + err.Error())
	}
}
