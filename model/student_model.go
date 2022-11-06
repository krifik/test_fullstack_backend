package model

import (
	"time"

	"gorm.io/gorm"
)

type StudentRequest struct {
	ID        int            `json:"id"`
	AvatarUrl string         `json:"avatar_url"`
	Name      string         `json:"name"`
	Gender    string         `json:"gender"`
	Dob       time.Time      `json:"date_of_birth"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `json:"DeletedAt"`
}

type StudentResponse struct {
	ID        int            `json:"id"`
	AvatarUrl string         `json:"avatar_url"`
	Name      string         `json:"name"`
	Gender    string         `json:"gender"`
	Dob       time.Time      `json:"date_of_birth"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
type StudentResponseWithPagination struct {
	ID         int            `json:"id"`
	AvatarUrl  string         `json:"avatar_url"`
	Name       string         `json:"name"`
	Gender     string         `json:"gender"`
	Dob        time.Time      `json:"date_of_birth"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	Limit      int            `json:"limit,omitempty;query:limit"`
	Page       int            `json:"page,omitempty;query:page"`
	Sort       string         `json:"sort,omitempty;query:sort"`
	TotalRows  int64          `json:"total"`
	TotalPages int            `json:"total_pages"`
}

type StudentFormRequest struct {
	AvatarUrl string    `json:"avatar_url" default:"app/storage/doge.png"`
	Name      string    `json:"name" default:"Fikri Ilhamsyah"`
	Gender    string    `json:"gender" default:"M"`
	Dob       time.Time `json:"date_of_birth" default:"2000-11-09T11:45:26.371Z"`
}
