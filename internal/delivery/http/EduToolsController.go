package http

import (
	"nutrimama/internal/model"
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EduToolsController struct {
	EduToolsUseCase *usecase.EduToolsUseCase
}

func NewEduToolsController(eduToolsUseCase *usecase.EduToolsUseCase) *EduToolsController {
	return &EduToolsController{
		EduToolsUseCase: eduToolsUseCase,
	}
}

func (c *EduToolsController) Create(ctx *fiber.Ctx) error {
	req := new(model.CreateEduToolsRequest)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := c.EduToolsUseCase.Create(*req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return utils.SuccessResponse(ctx, fiber.StatusCreated, "Request Successfull", res)
}

func (c *EduToolsController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id parameter"})
	}

	req := new(model.UpdateEduToolsRequest)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.ID = id

	res, err := c.EduToolsUseCase.Update(*req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return utils.SuccessResponse(ctx, fiber.StatusOK, "Request Successfull", res)
}

func (c *EduToolsController) Get(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id parameter"})
	}

	res, err := c.EduToolsUseCase.Get(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return utils.SuccessResponse(ctx, fiber.StatusOK, "Request Successfull", res)
}

func (c *EduToolsController) List(ctx *fiber.Ctx) error {
	res, err := c.EduToolsUseCase.List()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return utils.SuccessResponse(ctx, fiber.StatusOK, "Request Successfull", res)
}

func (c *EduToolsController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id parameter"})
	}

	err = c.EduToolsUseCase.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return utils.SuccessResponse(ctx, fiber.StatusOK, "Request Successfull", nil)
}
