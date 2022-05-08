package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	conn "go-api/pkg/connection"
	apiHandler "go-api/pkg/handler"
	"go-api/pkg/middleware/auth0"
	"go-api/pkg/middleware/cors"
	"go-api/pkg/middleware/logger"
	"go-api/pkg/model"

	"github.com/justinas/alice"
)

func main() {
	api := apiHandler.New()
	log := logger.NewLogger()

	// initializing middlewares
	logMiddleware := logger.NewMiddleware(log)
	corsMiddleware := cors.NewMiddleware()
	jwtMiddleware := auth0.NewMiddleware()

	handler := alice.New(corsMiddleware, logMiddleware, jwtMiddleware).Then(api)

	server := &http.Server{
		Addr:    conn.Port(),
		Handler: handler,
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
