package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string `json:"nickname" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	Phone    string `json:"phone"`
}
