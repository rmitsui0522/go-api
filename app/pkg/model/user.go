package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Account   string    `json:"mailAddress" validate:"required,email"`
	Password  string    `json:"password"`
	CreateAt  time.Time `json:"createAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

func FindUsers() ([]User, error) {
	var users []User
	err := DB.Find(&users).Error
	return users, err
}

func FindUser(u *User) (User, error) {
	var user User
	err := DB.Where(u).First(&user).Error
	return user, err
}

func CreateUser(u *User) error {
	u.CreateAt = time.Now().Round(time.Second)
	u.UpdateAt = time.Now().Round(time.Second)
	return DB.Create(u).Error
}

func UpdateUser(u *User) (User, error) {
	var user User
	u.UpdateAt = time.Now().Round(time.Second)

	err := DB.Where(&User{ID: u.ID}).Updates(u).First(&user).Error

	return user, err
}

func DeleteUser(u *User) (User, error) {
	var user User
	err := DB.Where(u).Delete(&user).Error
	return user, err
}
