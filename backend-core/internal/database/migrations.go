package database

import (
	"backend-core/internal/models"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	if err := models.MigrateUser(db); err != nil {
		log.Fatalf("cannot run migrations: %v", err.Error())
	}

	if err := models.MigrateProperty(db); err != nil {
		log.Fatalf("cannot run migrations: %v", err.Error())
	}

	if err := models.MigrateUserAuth(db); err != nil {
		log.Fatalf("cannot run migrations: %v", err.Error())
	}

	log.Println("Migrations run successfully")
}
