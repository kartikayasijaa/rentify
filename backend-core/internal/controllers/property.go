package controllers

import (
	"backend-core/internal/constants"
	inputstructs "backend-core/internal/structs/inputStructs"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Controller) PropertyCreateController(c *fiber.Ctx) error {
	property := new(inputstructs.PropertyCreateInput)
	userID := c.Locals("user_id").(uuid.UUID)

	if err := c.BodyParser(property); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.InputValidationFailedMessage,
			"error":   err.Error(),
		})
	}

	if err := property.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.InputValidationFailedMessage,
			"error":   err.Error(),
		})
	}

	// Create property
	createdProperty, err := h.Service.PropertyCreateService(property, userID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot create property",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdProperty)
}

func (h *Controller) PropertyGetController(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid page number",
			"error":   err.Error(),
		})
	}

	pageSizeStr := c.Query("pageSize", "10")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid page size",
			"error":   err.Error(),
		})
	}

	filters := parseFilters(c)

	properties, err := h.Service.PropertyGetService(page, pageSize, filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot get property",
			"error":   err.Error(),
		})
	}

	return c.JSON(properties)
}


func parseFilters(c *fiber.Ctx) map[string]interface{} {
    filters := make(map[string]interface{})

    if ownerId := c.Query("owner_id"); ownerId != "" {
        if ownerUUID, err := uuid.Parse(ownerId); err == nil {
            filters["owner_id"] = ownerUUID
        }
    }

    if city := c.Query("city"); city != "" {
        filters["city"] = city
    }

    if bedrooms := c.Query("bedrooms"); bedrooms != "" {
        if bedroomsInt, err := strconv.Atoi(bedrooms); err == nil {
            filters["bedrooms"] = bedroomsInt
        }
    }

    if minPrice := c.Query("min_price"); minPrice != "" {
        if minPriceFloat, err := strconv.ParseFloat(minPrice, 64); err == nil {
            filters["min_price"] = minPriceFloat
        }
    }

    if maxPrice := c.Query("max_price"); maxPrice != "" {
        if maxPriceFloat, err := strconv.ParseFloat(maxPrice, 64); err == nil {
            filters["max_price"] = maxPriceFloat
        }
    }

    if minSqft := c.Query("min_sqft"); minSqft != "" {
        if minSqftInt, err := strconv.Atoi(minSqft); err == nil {
            filters["min_sqft"] = minSqftInt
        }
    }

    if maxSqft := c.Query("max_sqft"); maxSqft != "" {
        if maxSqftInt, err := strconv.Atoi(maxSqft); err == nil {
            filters["max_sqft"] = maxSqftInt
        }
    }

    return filters
}