package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"go-api/pkg/handler"
	"go-api/pkg/utility"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const defaultPort = "3000"

func port() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":" + defaultPort
}

func dsn() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	protocol := os.Getenv("DB_PROTOCOL")
	dbName := os.Getenv("DB_DATABASE")

	if user == "" || password == "" || protocol == "" || dbName == "" {
		log.Fatal("dsn() failed: some empty go.Getenv variables.")
	}

	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		user,
		password,
		protocol,
		dbName,
	)

	return connStr
}

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
