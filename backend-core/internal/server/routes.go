package server

import (
	"backend-core/internal/controllers"
	middleware "backend-core/internal/middlewares"
	"backend-core/internal/routes"
	"backend-core/internal/services"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)
	api := s.App.Group("/api")
	middleware := middleware.New()
	services := services.New(s.DB)
	controllers := controllers.New(services)

	routes := routes.New(api, controllers, middleware)
	
	routes.AuthRoutes()
	routes.PropertyRoutes()
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Server is running!",
	}

	return c.JSON(resp)
}
