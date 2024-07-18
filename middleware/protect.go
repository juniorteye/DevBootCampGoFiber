package middleware

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/juniorteye/devCamp/database"
	"github.com/juniorteye/devCamp/model"
)

// Protect middleware to protect routes with JWT
func Protect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("Protect middleware called for:", c.OriginalURL())
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Not authorized to access this route",
			})
		}

		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(os.Getenv("JWT_SECRET")), nil // Use the secret key from environment variable
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Not authorized to access this route",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Not authorized to access this route",
			})
		}

		userID, ok := claims["userID"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Not authorized to access this route",
			})
		}

		var user model.User
		db := database.DB.Db
		if err := db.WithContext(context.Background()).Where("id = ?", userID).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Not authorized to access this route",
			})
		}
		c.Locals("user", user)
		return c.Next()
	}
}
