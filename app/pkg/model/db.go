package model

import (
	conn "go-api/pkg/connection"
	"go-api/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	var dsn *conn.DSN

	DB, err = gorm.Open(mysql.Open(dsn.MySQL()), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect database: " + err.Error())
	}
}
