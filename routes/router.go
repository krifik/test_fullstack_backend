package routes

import (
	"github.com/krifik/test_fullstack_backend/controller"
	"github.com/krifik/test_fullstack_backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Route(app *fiber.App, userController controller.UserController, studentController controller.StudentController) {
	app.Get("/api/users", middleware.AuthMiddleware, userController.FindAll)
	app.Post("/api/register", userController.Register)
	app.Post("/api/login", userController.Login)
	app.Get("/api/docs/*", swagger.HandlerDefault)
	app.Get("/api/students", studentController.FindAll)
	app.Get("/api/students/:id", studentController.Find)
	app.Post("/api/students", studentController.Store)
	app.Put("/api/students/:id", studentController.Update)
	app.Static("/public", "app/")
	app.Delete("/api/students/:id", studentController.Delete)
}
