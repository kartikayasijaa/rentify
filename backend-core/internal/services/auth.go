package services

import (
	"backend-core/internal/models"
	inputstructs "backend-core/internal/structs/inputStructs"
	outputstructs "backend-core/internal/structs/outputStructs"
	"backend-core/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
	"time"
)

func (h *Service) RegisterUserService(user *inputstructs.AuthRegisterUserInput) (*outputstructs.RegisterOutput, *fiber.Cookie, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot hash password: %v", err)
	}

	// Generate a new UUID for both User and UserAuth
	userID := uuid.New()

	newUser := &models.User{
		ID:        userID,
		Email:     user.Email,
		Mobile:    user.Mobile,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// Use the same ID for UserAuth
	newUserAuth := &models.UserAuth{
		ID:       userID,
		Password: hashedPassword,
		Email:    user.Email,
		Role:     models.RoleChoice(user.Role),
	}

	// Start a transaction
	tx := h.DB.Begin()

	// Create the user
	if err := tx.Create(newUser).Error; err != nil {
		tx.Rollback()
		if utils.IsUniqueConstraintViolation(err) {
			return nil, nil, fmt.Errorf("email already exists")
		}
		return nil, nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Create the user authentication record
	if err := tx.Create(newUserAuth).Error; err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("failed to create user authentication record: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	// Generate JWT token
	token, err := utils.GenerateJWTToken(newUser.ID, user.Role, 30*24*time.Hour)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate JWT token: %v", err)
	}

	// Generate and return cookie
	cookie := utils.GetCookie(token, os.Getenv("REFRESH_COOKIE_NAME"))

	res := &outputstructs.RegisterOutput{
		User: newUser,
	}

	return res, cookie, nil
}

func (h *Service) LoginUserService(user *inputstructs.AuthLoginUserInput) (*outputstructs.LoginOutput, *fiber.Cookie, error) {
	userModel := &models.UserAuth{}

	if err := h.DB.Preload("User").Where(&models.User{Email: user.Email}).First(userModel).Error; err != nil {
		return nil, nil, fmt.Errorf("user not found")
	}

	if err := utils.ComparePassword(userModel.Password, user.Password); err != nil {
		return nil, nil, fmt.Errorf("invalid password")
	}

	token, err := utils.GenerateJWTToken(userModel.ID, userModel.Role, 30*24*time.Hour)
	if err != nil {
		return nil, nil, err
	}

	cookie := utils.GetCookie(token, os.Getenv("REFRESH_COOKIE_NAME"))

	res := &outputstructs.LoginOutput{
		User: &userModel.User,
	}

	return res, cookie, nil
}

func (h *Service) RefreshTokenService(token string, userId uuid.UUID) (string, error) {

	userModel := &models.UserAuth{}
	if err := h.DB.Preload("User").Where(&models.User{ID: userId}).First(&userModel).Error; err != nil {
		return "", fmt.Errorf("user not found")
	}

	newToken, err := utils.GenerateJWTToken(userModel.ID, userModel.Role, 24*time.Hour)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
