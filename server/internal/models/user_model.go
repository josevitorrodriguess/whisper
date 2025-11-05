package models

import (
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey;size:128"`           
	Email     string    `gorm:"uniqueIndex;not null"`          
	Username  string    `gorm:"uniqueIndex;not null"`          
	PhotoURL  string    `gorm:"type:text"`                     
	CreatedAt time.Time `gorm:"autoCreateTime"`                
	UpdatedAt time.Time `gorm:"autoUpdateTime"`                
}