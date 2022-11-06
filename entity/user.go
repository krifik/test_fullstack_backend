package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primaryKey,not null,autoIncrement;uniqueIndex;"`
	Email     string `gorm:"size:256"`
	Password  string `gorm:"size:256"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
