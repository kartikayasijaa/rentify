package middleware

import (
	"backend-core/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (m *Middleware) SellerMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	userID, role, err := VerifyToken(authHeader)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if role != "seller" {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Unauthorized",
		})
	}

	c.Locals("user_id", userID)
	c.Locals("role", role)

	return c.Next()
}

func (m *Middleware) BuyerMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	userID, role, err := VerifyToken(authHeader)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if role != "buyer" {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Unauthorized",
		})
	}

	c.Locals("userID", userID)
	c.Locals("role", role)

	return c.Next()
}

func (m *Middleware) UserMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	userID, role, err := VerifyToken(authHeader)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	c.Locals("userID", userID)
	c.Locals("role", role)

	return c.Next()

}

func VerifyToken(authHeader string) (userID uuid.UUID, role string, err error) {

	if authHeader == "" {
		return uuid.Nil, "", fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return uuid.Nil, "", fiber.NewError(fiber.StatusUnauthorized, "Invalid Authorization header format")
	}

	tokenString := parts[1]

	_, claims, err := utils.VerifyJWTToken(tokenString)
	if err != nil {
		return uuid.Nil, "", fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	return uuid.MustParse(claims.ID), claims.Role, nil
}
