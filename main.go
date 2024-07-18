// main.go
package main

import (
	// "time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/juniorteye/devCamp/database"
	_ "github.com/juniorteye/devCamp/docs" // This is necessary for Swagger to work
	"github.com/juniorteye/devCamp/migrations"
	"github.com/juniorteye/devCamp/router"
	"github.com/juniorteye/devCamp/validation"
	_ "github.com/lib/pq"
	fiberswagger "github.com/swaggo/fiber-swagger" // fiber-swagger middleware
	// "gorm.io/gorm"
)

// @title DevCamp API
// @version 1.0
// @description This is a sample server for DevCamp.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	database.Connect()

	app := fiber.New()

	// Initialize the validator
	validation.InitValidator()

	// Run the migrations
	migrations.Migrate()
	app.Use(logger.New())
	app.Use(cors.New())

	// Swagger route
	app.Get("/swagger/*", fiberswagger.WrapHandler)

	// Public routes
	router.SetupRoutes(app)

	// JWT Middleware applied to protected routes only
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	// Protected routes
	// router.SetupProtectedRoutes(app)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":8080")
}
