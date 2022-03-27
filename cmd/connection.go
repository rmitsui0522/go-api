package main

import (
	"fmt"
	"log"
	"os"
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
