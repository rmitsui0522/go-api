package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	conn "go-api/pkg/connection"

	"go-api/pkg/v1/handler"
	"go-api/pkg/v1/model"
)

func main() {
	h := handler.New()

	server := &http.Server{
		Addr:    conn.Port(),
		Handler: h,
	}

	defer model.DB.Close()

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)

		<-stop
		log.Println("Shutting down...")

		if err := server.Shutdown(context.Background()); err != nil {
			log.Println("server.Shutdown() failed: ", err)
		}
		log.Println("Server stopped.")
	}()

	fmt.Println("server listening on http://localhost" + conn.Port())
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("server.ListenAndServe() failed: ", err)
	}
}
