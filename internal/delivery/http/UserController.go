package http

import (
	"nutrimama/internal/model"
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"

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

func (c *UserController) Login(ctx *fiber.Ctx) error {
	req := new(model.LoginRequest)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := c.UserUseCase.Login(req.Email, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return utils.SuccessResponse(ctx, 201, "Login successful", user)
}

func (c *UserController) Profile(ctx *fiber.Ctx) error {
	return utils.SuccessResponse(ctx, 200, "This is a protected route", nil)
}