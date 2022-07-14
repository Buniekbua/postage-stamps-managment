package models

import "time"

type Stamp struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Quantity    uint      `json:"quantity"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	UserRefer   int       `json:"user_id"`
	User        User      `gorm:"foreignKey:UserRefer"`
}
