package models

import (
	"time"
)

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Email      string    `json:"email" gorm:"unique"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
