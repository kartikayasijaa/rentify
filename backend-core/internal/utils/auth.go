package utils

import (
	"backend-core/internal/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = "secret"

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateJWTToken(userID uuid.UUID, role models.RoleChoice, expire time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   userID,
		"role": role,
		"exp":  time.Now().Add(expire).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type CustomClaims struct {
	jwt.StandardClaims
	ID   string `json:"id"`
	Role string `json:"role"`
}

func VerifyJWTToken(tokenString string) (*jwt.Token, *CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	return token, claims, nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetCookie(data string, name string) *fiber.Cookie {
	isProd := true
	if os.Getenv("APP_ENV") == "local" {
		isProd = false
	}
	cookie := new(fiber.Cookie)
	cookie.HTTPOnly = true
	cookie.Name = name
	cookie.Value = data
	cookie.Secure = isProd
	cookie.Expires = time.Now().Add(time.Hour * 24 * 30)
	cookie.SameSite = "None"
	return cookie
}
