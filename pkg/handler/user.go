package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type User struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (h *handler) getUsers(c echo.Context) error {
	var users []User

	err := h.DB.Find(&users).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

// SQL
// CREATE TABLE `users` (
// 	`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "ユーザーID",
// 	`username` VARCHAR(100) NOT NULL COMMENT "ユーザー名",
// 	`createdAt` timestamp DEFAULT NULL COMMENT "登録日",
// 	`updatedAt` datetime DEFAULT NULL COMMENT "更新日"
// );

// INSERT INTO users (username, createdAt, updatedAt) VALUES ("test user", now(), now());
