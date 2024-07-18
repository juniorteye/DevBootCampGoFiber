package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juniorteye/devCamp/handler"
	"github.com/juniorteye/devCamp/middleware"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.Post("/", handler.CreateUser)
	users.Get("/", middleware.Protect(), handler.GetAllUsers)
	users.Get("/:id", middleware.Protect(), handler.GetUser)
	users.Delete("/:userId", middleware.Protect(), handler.DeleteUser)
	users.Put("/:userId", middleware.Protect(), handler.UpdateUser)

	auth := api.Group("/auth")
	auth.Post("/signup", handler.SignUp)
	auth.Post("/login", handler.Login)
	auth.Put("/changepassword", handler.ChangePassword)
	auth.Post("/forgetpassword", handler.ForgetPassword)

	// bootcamp routes
	bootcamp := api.Group("/bootcamp")
	bootcamp.Post("/", middleware.Protect(), handler.CreateBootCamp)
	bootcamp.Get("/", handler.GetAllBootCamp)
	bootcamp.Get("/:bootcampId", handler.GetBootCamp)
	bootcamp.Put("/:bootcampId", handler.UpDateBootCamp)
	bootcamp.Delete("/:bootcampId", handler.DeleteBootCamp)

	// Review Routes
	review := api.Group("/review")
	review.Post("/:bootcampId", middleware.Protect(), handler.CreateReview)
	review.Get("/bootcamps/:bootcampId", handler.GetReviews)
	review.Get("/:reviewId", handler.GetReview)
	review.Put("/:reviewId", handler.UpdateReview)
	review.Delete("/:reviewId", handler.DeleteReview)
}
