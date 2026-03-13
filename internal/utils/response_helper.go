package utils

import (
	"nutrimama/internal/model"

	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(model.Response{
		Message: message,
		Data:    data,
	})
}

// Error response
func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(model.Response{
		Message: message,
		Data:    nil,
	})
}