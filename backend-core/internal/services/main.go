package services

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) *Service {
	return &Service{
		DB: DB,
	}
}
