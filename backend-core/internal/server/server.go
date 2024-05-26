package server

import (
	"backend-core/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

type FiberServer struct {
	*fiber.App
	DB *gorm.DB
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "backend-core",
			AppName:      "backend-core",
		}),

		DB: database.New(),
	}

	server.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "*",
		AllowHeaders:     "Content-Type,Authorization",
		AllowCredentials: true,
	}))

	return server
}
