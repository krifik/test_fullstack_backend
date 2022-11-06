package service

import (
	"github.com/krifik/test_fullstack_backend/entity"
	"github.com/krifik/test_fullstack_backend/helper"
	"github.com/krifik/test_fullstack_backend/model"
	"github.com/krifik/test_fullstack_backend/repository"
	"github.com/krifik/test_fullstack_backend/validation"
)

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepository
}

func NewStudentServiceImpl(studentRepository repository.StudentRepository) StudentService {
	return &StudentServiceImpl{StudentRepository: studentRepository}
}

func (service *StudentServiceImpl) FindAll(pagination helper.Pagination) ([]model.StudentResponse, *helper.Pagination) {
	students := service.StudentRepository.FindAll(pagination)
	var responses []model.StudentResponse
	for _, student := range students.Data {
		responses = append(responses, model.StudentResponse{
			ID:        int(student.ID),
			AvatarUrl: student.AvatarUrl,
			Name:      student.Name,
			Gender:    student.Gender,
			Dob:       student.Dob,
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
			DeletedAt: student.DeletedAt,
		})
	}
	return responses, students
}
func (service *StudentServiceImpl) Search(key string) []model.StudentResponse {
	students := service.StudentRepository.Search(key)
	var responses []model.StudentResponse
	for _, student := range students {
		responses = append(responses, model.StudentResponse{
			ID:        int(student.ID),
			AvatarUrl: student.AvatarUrl,
			Name:      student.Name,
			Gender:    student.Gender,
			Dob:       student.Dob,
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
			DeletedAt: student.DeletedAt,
		})
	}
	return responses
}

func (service *StudentServiceImpl) Find(id int) model.StudentResponse {
	student := service.StudentRepository.Find(id)
	response := model.StudentResponse{
		ID:        int(student.ID),
		AvatarUrl: student.AvatarUrl,
		Name:      student.Name,
		Gender:    student.Gender,
		Dob:       student.Dob,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
		DeletedAt: student.DeletedAt,
	}
	return response
}
func (service *StudentServiceImpl) Store(request model.StudentRequest) model.StudentResponse {
	validation.ValidateStudent(request)
	student := entity.Student{
		AvatarUrl: request.AvatarUrl,
		Name:      request.Name,
		Gender:    request.Gender,
		Dob:       request.Dob,
	}
	service.StudentRepository.Store(&student)
	response := model.StudentResponse{
		ID:        int(student.ID),
		AvatarUrl: student.AvatarUrl,
		Name:      student.Name,
		Gender:    student.Gender,
		Dob:       student.Dob,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
		DeletedAt: student.DeletedAt,
	}
	return response
}

func (service *StudentServiceImpl) Update(id int, request model.StudentRequest) {
	validation.ValidateStudent(request)

	student := entity.Student{
		AvatarUrl: request.AvatarUrl,
		Name:      request.Name,
		Gender:    request.Gender,
		Dob:       request.Dob,
	}
	service.StudentRepository.Update(id, student)
}

func (service *StudentServiceImpl) Delete(id int) {
	service.StudentRepository.Delete(id)
}
