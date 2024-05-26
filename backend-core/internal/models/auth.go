package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAuth struct {
    ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
    Email    string    `json:"email" validate:"required,email" gorm:"unique"`
    Password string    `json:"password" validate:"required"`
    Role     RoleChoice `json:"role" validate:"required,eq=seller|eq=buyer"`
    User     User      `json:"user" gorm:"foreignKey:ID"`
}

type RoleChoice string

const (
	Seller RoleChoice = "seller"
	Buyer  RoleChoice = "buyer"
)

// auto migration
func MigrateUserAuth(db *gorm.DB) error {
	err := db.AutoMigrate(&UserAuth{})
	return err
}