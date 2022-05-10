package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	conn "go-api/pkg/connection"
	"go-api/pkg/handler"
	"go-api/pkg/middleware/logger"
	"go-api/pkg/model"
)

func main() {
	log := logger.NewLogger()

	server := &http.Server{
		Addr:    conn.Port(),
		Handler: handler.New(),
	}

	defer model.DB.Close()

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)

		<-stop
		log.Info().Msg("Shutting down...")

		if err := server.Shutdown(context.Background()); err != nil {
			log.Error().Err(err).Msg("server.Shutdown() failed: ")
		}
		log.Info().Msg("Server stopped.")
	}()

	log.Info().Msg("server listening on http://localhost" + conn.Port())
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Fatal: server.ListenAndServe()")
	}
}
