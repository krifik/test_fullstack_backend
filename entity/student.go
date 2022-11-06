package entity

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	AvatarUrl string `gorm:"size: 256"`
	Name      string `gorm:"size: 256"`
	Gender    string `gorm:"size: 1"`
	Dob       time.Time
}
