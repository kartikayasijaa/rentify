package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Property struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"omitempty"`
	Price       float64   `json:"price" validate:"required"`
	Bedrooms    int       `json:"bedrooms" validate:"required"`
	Bathrooms   int       `json:"bathrooms" validate:"required"`
	Sqft        int       `json:"sqft" validate:"required"`
	Location    string    `json:"location" validate:"required"`
	City        string    `json:"city" validate:"required"`
	Zipcode     string    `json:"zipcode" validate:"required"`
	OwnerID     uuid.UUID `json:"owner_id" gorm:"foreignKey:Owner"` // Field to store the owner ID and define foreign key constraint
    Owner       User      `json:"owner,omitempty" gorm:"foreignKey:OwnerID"` // Field to preload the full owner object
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// auto migration
func MigrateProperty(db *gorm.DB) error {
	err := db.AutoMigrate(&Property{})
	return err
}
