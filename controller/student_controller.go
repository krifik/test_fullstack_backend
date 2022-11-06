package controller

import (
	"github.com/gofiber/fiber/v2"
)

type StudentController interface {
	Find(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Search(ctx *fiber.Ctx) error
}
