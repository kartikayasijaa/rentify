package services

import (
	"backend-core/internal/models"
	inputstructs "backend-core/internal/structs/inputStructs"
	"time"

	"github.com/google/uuid"
)

func (s *Service) PropertyCreateService(property *inputstructs.PropertyCreateInput, ownerId uuid.UUID) (*models.Property, error) {
	// Create the property
	newProperty := &models.Property{
		ID:          uuid.New(),
		Name:        property.Name,
		Description: property.Description,
		Price:       property.Price,
		Bedrooms:    property.Bedrooms,
		Bathrooms:   property.Bathrooms,
		Sqft:        property.Sqft,
		Location:    property.Location,
		City:        property.City,
		Zipcode:     property.Zipcode,
		OwnerID:     ownerId,
		UpdatedAt:   time.Now(),
	}

	// Create the property
	if err := s.DB.Create(newProperty).Error; err != nil {
		return nil, err
	}

	// Preload the associated owner
	if err := s.DB.Preload("Owner").First(newProperty, newProperty.ID).Error; err != nil {
		return nil, err
	}

	return newProperty, nil
}

func (s *Service) PropertyUpdateService() {

}

func (s *Service) PropertyDeleteService() {

}

func (s *Service) PropertyGetService(page int, pageSize int, filters map[string]interface{}) ([]models.Property, error) {
	var properties []models.Property

	query := s.DB.Model(&models.Property{})

	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	// Apply filters
	for key, value := range filters {
		switch key {
		case "owner_id":
			query = query.Where("owner_id = ?", value)
		case "city":
			query = query.Where("city = ?", value)
		case "bedrooms":
			query = query.Where("bedrooms = ?", value)
		case "minPrice":
			query = query.Where("price >= ?", value)
		case "maxPrice":
			query = query.Where("price <= ?", value)
		case "minSqft":
			query = query.Where("sqft >= ?", value)
		case "maxSqft":
			query = query.Where("sqft <= ?", value)
		}
	}

	// Execute the query
	if err := query.Preload("Owner").Find(&properties).Error; err != nil {
		return nil, err
	}

	return properties, nil
}
