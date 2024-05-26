package outputstructs

import "backend-core/internal/models"

type RegisterOutput struct {
	User *models.User `json:"user"`
}

type LoginOutput struct {
	User  *models.User `json:"user"`
}