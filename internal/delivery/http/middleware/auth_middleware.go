package middleware

import (
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(u *usecase.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "missing token",
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		
		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		c.Locals("userID", claims.ID)
		c.Locals("email", claims.Email)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}