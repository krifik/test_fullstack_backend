package service

import (
	"github.com/krifik/test_fullstack_backend/helper"
	"github.com/krifik/test_fullstack_backend/model"
)

type StudentService interface {
	Find(id int) model.StudentResponse
	FindAll(pagination helper.Pagination) ([]model.StudentResponse, *helper.Pagination)
	Update(id int, request model.StudentRequest)
	Store(model.StudentRequest) model.StudentResponse
	Delete(id int)
	Search(key string) []model.StudentResponse
}
