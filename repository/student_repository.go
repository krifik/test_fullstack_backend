package repository

import (
	"github.com/krifik/test_fullstack_backend/entity"
	"github.com/krifik/test_fullstack_backend/helper"
)

type StudentRepository interface {
	FindAll(pagination helper.Pagination) *helper.Pagination
	Find(id int) entity.Student
	Delete(id int)
	Store(student *entity.Student)
	Update(id int, student entity.Student)
	Search(key string) []entity.Student
}
