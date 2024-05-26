package inputstructs

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PropertyCreateInput struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"omitempty"`
	Price       float64   `json:"price" validate:"required"`
	Bedrooms    int       `json:"bedrooms" validate:"omitempty"`
	Bathrooms   int       `json:"bathrooms" validate:"omitempty"`
	Sqft        int       `json:"sqft" validate:"required"`
	Location    string    `json:"location" validate:"required"`
	City        string    `json:"city" validate:"required"`
	Zipcode     string    `json:"zipcode" validate:"required"`
	OwnerID     uuid.UUID `json:"owner_id" gorm:"foreignKey:Owner"`
}

func (input *PropertyCreateInput) Validate() error {
	validate := validator.New()
	return validate.Struct(input)
}
