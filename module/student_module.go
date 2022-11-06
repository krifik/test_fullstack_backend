package module

import (
	"github.com/krifik/test_fullstack_backend/controller"
	"github.com/krifik/test_fullstack_backend/repository"
	"github.com/krifik/test_fullstack_backend/service"

	"gorm.io/gorm"
)

func NewStudentModule(database *gorm.DB) controller.StudentController {
	// Setup Repository
	studentRepository := repository.NewStudentRepositoryImpl(database)

	// Setup Service
	studentService := service.NewStudentServiceImpl(studentRepository)

	// Setup Controller
	studentController := controller.NewStudentControllerImpl(studentService)
	return studentController

}
