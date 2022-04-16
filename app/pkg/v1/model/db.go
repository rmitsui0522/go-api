package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", dsn())
	if err != nil {
		log.Fatal("gorm.Open() failed: ", err)
	}
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
