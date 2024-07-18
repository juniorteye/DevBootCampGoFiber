package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	// "gorm.io/gorm"
)

// custom errors

func HandleDBError(err error) fiber.Map {
	if strings.Contains(err.Error(), "duplicate key value") {
		return fiber.Map{
			"status":  "error",
			"message": "could not create a bootcamp",
			"data":    "A bootcamp with the same name already exists.",
		}
	}
	if strings.Contains(err.Error(), "22P02") {
		return fiber.Map{
			"status":  "error",
			"message": "could not get a bootcamp",
			"data":    "invalid bootcamp Id",
		}
	}
	return fiber.Map{
		"status":  "error",
		"message": "could not create a bootcamp",
		"error":   err.Error(),
	}
}
