package models

import (
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Email     string `gorm:"unique;not null"`
	Username  string `gorm:"unique;not null"`
	PhotoURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
