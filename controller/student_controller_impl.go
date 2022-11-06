package controller

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/krifik/test_fullstack_backend/exception"
	"github.com/krifik/test_fullstack_backend/helper"
	"github.com/krifik/test_fullstack_backend/model"
	"github.com/krifik/test_fullstack_backend/service"
)

type StudentControllerImpl struct {
	StudentService service.StudentService
}

func NewStudentControllerImpl(studentService service.StudentService) StudentController {
	return &StudentControllerImpl{StudentService: studentService}
}

// Get all data student godoc
// @Summary      Show all data student
// @Description  Get all data
// @Tags         student
// @Accept       json
// @Produce      json
// @Success 200 {object} map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/students [get]
func (controller *StudentControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var pagination helper.Pagination
	err := ctx.QueryParser(&pagination)
	responses, students := controller.StudentService.FindAll(pagination)
	exception.PanicIfNeeded(err)
	return ctx.Status(200).JSON(fiber.Map{
		"code":        200,
		"status":      "Get all data student",
		"limit":       students.Limit,
		"page":        students.Page,
		"sort":        students.Sort,
		"order":       students.Order,
		"total":       students.TotalRows,
		"search":      students.Search,
		"total_pages": students.TotalRows / int64(students.Limit),
		"data":        responses,
	})
}
func (controller *StudentControllerImpl) Search(ctx *fiber.Ctx) error {
	var pagination helper.Pagination
	err := ctx.QueryParser(&pagination)
	exception.PanicIfNeeded(err)
	key := ctx.Query("student")

	responses := controller.StudentService.Search(key)
	return ctx.Status(200).JSON(fiber.Map{
		"code":   200,
		"status": "Get all data student",
		"data":   responses,
	})
}

// Get data student by id godoc
// @Summary      Show specific data student
// @Description  Get specific data
// @Tags         student
// @Accept       json
// @Produce      json
// @Param id path int true "Student Id"
// @Success 200 {object} map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/students/{id} [get]
func (controller *StudentControllerImpl) Find(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	exception.PanicIfNeeded(err)
	responses := controller.StudentService.Find(id)
	return ctx.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "Get specific student",
		Data:   responses,
	})
}

// Store data student godoc
// @Summary      Store data student
// @Description  Store
// @Tags         student
// @Accept       json
// @Produce      json
// @Param request body model.StudentFormRequest true "Form Student"
// @Success 200 {object} map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/students [post]
func (controller *StudentControllerImpl) Store(ctx *fiber.Ctx) error {
	var request model.StudentRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)
	file, err := ctx.FormFile("avatar_url")
	exception.PanicIfNeeded(err)
	t := time.Now().Unix()
	randFilename := strconv.Itoa(int(t)) + file.Filename
	request.AvatarUrl = os.Getenv("BASE_URL") + "public/storage/" + randFilename
	responses := controller.StudentService.Store(request)
	ctx.SaveFile(file, "app/storage/"+randFilename)
	return ctx.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "Store data student",
		Data:   responses,
	})
}

// Delete data student godoc
// @Summary      Delete data student
// @Description  Delete
// @Tags         student
// @Accept       json
// @Produce      json
// @Param id path int true "Student Id"
// @Success 200 {object} map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/students/{id} [delete]
func (controller *StudentControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	exception.PanicIfNeeded(err)
	student := controller.StudentService.Find(id)
	url := strings.Split(student.AvatarUrl, os.Getenv("BASE_URL"))
	sanitazeStr := strings.Split(url[1], "public")
	os.Remove("app/" + sanitazeStr[1])
	controller.StudentService.Delete(id)
	return ctx.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Successfully delete student",
	})
}

// Update data student godoc
// @Summary      Update data student by id
// @Description  Update
// @Tags         student
// @Accept       json
// @Produce      json
// @Param id path int true "Student Id"
// @Param request body model.StudentFormRequest true "Form Update"
// @Success 200 {object} map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/students/{id} [put]
func (controller *StudentControllerImpl) Update(ctx *fiber.Ctx) error {
	var request model.StudentRequest
	id, err := ctx.ParamsInt("id")
	exception.PanicIfNeeded(err)
	err = ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)
	file, err := ctx.FormFile("avatar_url")
	if err != nil {
		controller.StudentService.Update(id, request)
	} else {
		t := time.Now().Unix()
		randFilename := strconv.Itoa(int(t)) + file.Filename
		request.AvatarUrl = os.Getenv("BASE_URL") + "public/storage/" + randFilename
		controller.StudentService.Update(id, request)
		url := strings.Split(file.Filename, os.Getenv("BASE_URL"))
		sanitazeStr := strings.Split(url[1], "public")
		os.Remove("app/" + sanitazeStr[1])
		ctx.SaveFile(file, "app/storage/"+randFilename)
		controller.StudentService.Update(id, request)
	}
	return ctx.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "successfully updated data student with id " + strconv.Itoa(id),
	})
}
