package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
}
