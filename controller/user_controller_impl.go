package controller

import (
	"time"

	"github.com/krifik/test_fullstack_backend/exception"
	"github.com/krifik/test_fullstack_backend/helper"
	"github.com/krifik/test_fullstack_backend/model"
	"github.com/krifik/test_fullstack_backend/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

// Register new user godoc
// @Summary      Show an registered account
// @Description  Create a new user account
// @Tags         user
// @Accept       json
// @Produce      json
// @Param request body model.UserRequest true "Form User"
// @Success 200 {object} map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/register [post]
func (controller *UserControllerImpl) Register(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response, err := controller.UserService.Register(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
func (controller *UserControllerImpl) FindAll(c *fiber.Ctx) error {
	responses, err := controller.UserService.FindAll()
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}

// Login godoc
// @Summary      Show a logged in account
// @Description  Get a token to be logged in
// @Tags         user
// @Accept       json
// @Produce      json
// @Param request body model.UserRequest true "Form User"
// @Success 200 {object} map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/login [post]
func (controller *UserControllerImpl) Login(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response, err := controller.UserService.Login(request)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Wrong credential",
		})
	}

	claims := jwt.MapClaims{}
	claims["email"] = response.Email
	claims["expired_at"] = time.Now().Add(60 * time.Minute).Unix()
	token, err := helper.GenerateJWT(&claims)
	exception.PanicIfNeeded(err)

	return c.Status(200).JSON(fiber.Map{
		"token": token,
	})
}
