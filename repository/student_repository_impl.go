package repository

import (
	"math"

	"github.com/krifik/test_fullstack_backend/config"
	"github.com/krifik/test_fullstack_backend/entity"
	"github.com/krifik/test_fullstack_backend/exception"
	"github.com/krifik/test_fullstack_backend/helper"
	"gorm.io/gorm"
)

type StudentRepositoryImpl struct {
	DB *gorm.DB
}

func NewStudentRepositoryImpl(db *gorm.DB) StudentRepository {
	return &StudentRepositoryImpl{DB: db}
}
func (repository *StudentRepositoryImpl) Store(student *entity.Student) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	err := repository.DB.WithContext(ctx).Create(&student).Error
	exception.PanicIfNeeded(err)
}

func (repository *StudentRepositoryImpl) FindAll(pagination helper.Pagination) *helper.Pagination {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	var students []entity.Student
	err := repository.DB.WithContext(ctx).Scopes(paginate(students, &pagination, repository.DB), pagination.SearchStudent).Find(&students).Error
	pagination.Data = students
	exception.PanicIfNeeded(err)
	return &pagination
}
func (repository *StudentRepositoryImpl) Search(key string) []entity.Student {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	var students []entity.Student
	err := repository.DB.WithContext(ctx).Where("name iLIKE ?", "%"+key+"%").Find(&students).Error
	exception.PanicIfNeeded(err)
	return students
}
func (repository *StudentRepositoryImpl) Find(id int) entity.Student {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	var student entity.Student
	err := repository.DB.WithContext(ctx).First(&student, "id = ?", id).Error
	exception.PanicIfNeeded(err)
	return student
}

func (repository *StudentRepositoryImpl) Delete(id int) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	var student entity.Student
	err := repository.DB.WithContext(ctx).Delete(&student, id).Error
	exception.PanicIfNeeded(err)
}
func (repository *StudentRepositoryImpl) Update(id int, student entity.Student) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	err := repository.DB.WithContext(ctx).Where("id", id).Updates(&student).Error
	exception.PanicIfNeeded(err)
}

func paginate(value interface{}, pagination *helper.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	// pagination.Limit = 10
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort() + " " + pagination.GetOrder())
	}
}
