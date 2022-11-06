package repository

import (
	"github.com/krifik/test_fullstack_backend/entity"
	"github.com/krifik/test_fullstack_backend/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(request model.CreateUserRequest) (user entity.User, err error)
	FindAll() ([]entity.User, error)
	Login(model.CreateUserRequest) (user entity.User, err error)
	Delete(db *gorm.DB, userId int)
	CheckEmail(request model.CreateUserRequest) (result int64)
}
