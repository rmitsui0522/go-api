package connection

import (
	"fmt"
	"log"
	"os"
)

type DSN struct{}

func (d *DSN) MySQL() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	protocol := os.Getenv("DB_PROTOCOL")
	dbName := os.Getenv("DB_NAME")

	if user == "" || password == "" || protocol == "" || dbName == "" {
		log.Fatal("dsn() failed: some empty variables from os.Getenv.")
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
