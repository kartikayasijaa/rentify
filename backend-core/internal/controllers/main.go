package controllers

import "backend-core/internal/services"

type Controller struct {
	Service *services.Service
}

func New(service *services.Service) *Controller {
	return &Controller{
		Service: service,
	}
}