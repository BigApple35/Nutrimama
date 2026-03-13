package route

import (
	"nutrimama/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App           *fiber.App
	UserController *http.UserController
}

func (c *RouteConfig) Setup() {
	SetupGuestRoutes(c)
}

func SetupGuestRoutes(c *RouteConfig) {
	c.App.Post("/api/register", c.UserController.Register)
}