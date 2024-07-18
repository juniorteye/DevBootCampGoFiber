package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juniorteye/devCamp/model"
)

func Authorize(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(model.User)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "User not found in context",
			})
		}

		for _, role := range roles {
			if string(user.Role) == role {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "User role is not authorized to perform this action",
		})
	}
}
