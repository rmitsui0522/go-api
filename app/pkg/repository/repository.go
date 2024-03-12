package repository

import (
	conn "go-api/pkg/connection"
	"go-api/pkg/logger"
	"go-api/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var dsn *conn.DSN
	var err error

	db, err = gorm.Open(mysql.Open(dsn.MySQL()), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect database: " + err.Error())
	}

	db.AutoMigrate(&model.Employee{}, &model.Employee{})
}
