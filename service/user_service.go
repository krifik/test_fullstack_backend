package service

import (
	"github.com/krifik/test_fullstack_backend/model"
)

type UserService interface {
	Register(request model.CreateUserRequest) (response model.CreateUserResponse, err error)
	FindAll() ([]model.GetUserResponse, error)
	Login(request model.CreateUserRequest) (response model.CreateUserResponse, err error)
}
