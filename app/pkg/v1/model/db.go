package model

import (
	"log"

	conn "go-api/pkg/connection"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", conn.Dsn())
	if err != nil {
		log.Fatal("gorm.Open() failed: ", err)
	}
}
