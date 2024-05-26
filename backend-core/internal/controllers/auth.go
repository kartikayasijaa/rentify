package controllers

import (
	"backend-core/internal/constants"
	inputstructs "backend-core/internal/structs/inputStructs"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Controller) RegisterUserController(c *fiber.Ctx) error {
	user := new(inputstructs.AuthRegisterUserInput)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.InputValidationFailedMessage,
			"error":   err.Error(),
		})
	}

	if err := user.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.InputValidationFailedMessage,
			"error":   err.Error(),
		})
	}

	res, cookie, err := h.Service.RegisterUserService(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user!",
			"error":   err.Error(),
		})
	}

	c.Cookie(cookie)
	return c.JSON(res)
}

func (h *Controller) LoginUserController(c *fiber.Ctx) error {
	user := new(inputstructs.AuthLoginUserInput)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.InputValidationFailedMessage,
			"error":   err.Error(),
		})
	}

	if err := user.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.InputValidationFailedMessage,
			"error":   err.Error(),
		})
	}

	response, cookie, err := h.Service.LoginUserService(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to login user!",
			"error":   err.Error(),
		})
	}

	c.Cookie(cookie)
	return c.JSON(response)
}

func (h *Controller) RefreshTokenController(c *fiber.Ctx) error {
	token := c.Cookies(os.Getenv("REFRESH_COOKIE_NAME"))
	userId := uuid.MustParse(c.Locals("user_id").(string))
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Refresh token is required",
		})
	}

	newToken, err := h.Service.RefreshTokenService(token, userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to refresh token",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": newToken,
	})
}
