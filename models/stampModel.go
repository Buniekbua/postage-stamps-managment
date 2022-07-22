package models

import (
	"gorm.io/gorm"
)

type Stamp struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Quantity    uint    `json:"quantity"`
	UserRefer   int     `json:"user_id"`
	User        User    `gorm:"foreignKey:UserRefer"`
}
