package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	MailAddress string `json:"mailAddress" validate:"email" gorm:"unique"`
}
