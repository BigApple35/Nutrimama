package http

import (
	"nutrimama/internal/model"
	"nutrimama/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserController(userUseCase *usecase.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	req := new(model.RegisterRequest)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := c.UserUseCase.Create(*req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(user)
}
