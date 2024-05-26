package inputstructs

import (
	"backend-core/internal/models"

	"github.com/go-playground/validator/v10"
)

type AuthRegisterUserInput struct {
	Email     string            `json:"email" validate:"required,email"`
	Password  string            `json:"password" validate:"required"`
	Role      models.RoleChoice `json:"role" validate:"required,eq=seller|eq=buyer"`
	FirstName string            `json:"first_name" validate:"required"`
	LastName  string            `json:"last_name" validate:"omitempty"`
	Mobile    string            `json:"mobile" validate:"omitempty"`
}
func (input *AuthRegisterUserInput) Validate() error {
	validate := validator.New()
	return validate.Struct(input)
}

type AuthLoginUserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
func (input *AuthLoginUserInput) Validate() error {
	validate := validator.New()
	return validate.Struct(input)
}


