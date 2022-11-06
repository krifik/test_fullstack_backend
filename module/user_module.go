package module

import (
	"github.com/krifik/test_fullstack_backend/controller"
	"github.com/krifik/test_fullstack_backend/repository"
	"github.com/krifik/test_fullstack_backend/service"

	"gorm.io/gorm"
)

func NewUserModule(database *gorm.DB) controller.UserController {
	// Setup Repository
	userRepository := repository.NewUserRepositoryImpl(database)

	// Setup Service
	userService := service.NewUserServiceImpl(userRepository)

	// Setup Controller
	userController := controller.NewUserControllerImpl(userService)
	return userController

}
