package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	conn "go-api/pkg/connection"
	apiHandler "go-api/pkg/handler"
	"go-api/pkg/middleware/auth0"
	"go-api/pkg/middleware/logger"
	"go-api/pkg/model"

	"github.com/rs/cors"
)

func main() {
	api := apiHandler.New()

	log := logger.NewLogger()
	logMiddleware := logger.NewMiddleware(log)
	jwtMiddleware, err := auth0.NewMiddleware()
	if err != nil {
		log.Fatal().Err(err).Msg("Fatal: initialize jwt middleware")
	}

	api.Use(logMiddleware)
	api.Use(jwtMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization"},
	})

	handler := c.Handler(api)

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
