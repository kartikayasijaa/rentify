package routes

import (
	"backend-core/internal/controllers"
	middleware "backend-core/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	App        fiber.Router
	Controller *controllers.Controller
	Middleware *middleware.Middleware
}

func New(app fiber.Router, controller *controllers.Controller, middleware *middleware.Middleware) *Routes {
	return &Routes{
		App:        app,
		Controller: controller,
		Middleware: middleware,
	}
}
