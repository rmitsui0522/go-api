package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"go-api/pkg/utility"
	"go-api/pkg/v1/handler"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	utility.LoadEnv()

	db, err := gorm.Open("mysql", dsn())
	if err != nil {
		log.Fatal("gorm.Open() failed: ", err)
	}
	defer db.Close()

	h := handler.New(db)

	server := &http.Server{
		Addr:    port(),
		Handler: h,
	}

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

	fmt.Println("server listening on http://localhost" + port())
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("server.ListenAndServe() failed: ", err)
	}
}
