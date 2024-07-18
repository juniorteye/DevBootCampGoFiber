package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/juniorteye/devCamp/database"
	"github.com/juniorteye/devCamp/middleware"
	"github.com/juniorteye/devCamp/model"
)

// CreateReview godoc
// @Summary Create a new Review
// @Description Create a new Review
// @Tags Reviews
// @Produce json
// @Param title string true "Title"
// @Param Text  string true "Text"
// @Param rating  int true "rating"
// @Success 201 {object} model.Review
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /review [post]
func CreateReview(c *fiber.Ctx) error {
	newReview := new(model.Review)
	reviewDB := database.DB.Db
	var bootcamp model.Bootcamp

	bootcampID := c.Params("bootcampId")
	user, ok := c.Locals("user").(model.User)

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "User not found in context"})
	}
	userID := user.ID

	if err := c.BodyParser(newReview); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Provide the correct input", "data": err.Error()})
	}

	bootcampUUID, err := uuid.Parse(bootcampID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid bootcamp ID"})
	}
	newReview.BootcampID = bootcampUUID
	newReview.UserID = userID

	// Check if the bootcamp exists
	if err := reviewDB.First(&bootcamp, "id = ?", bootcampUUID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No bootcamp found with the specified ID"})
	}

	// Check if the user has already submitted a review for this bootcamp
	var existingReview model.Review
	if err := reviewDB.Where("bootcamp_id = ? AND user_id = ?", bootcampUUID, userID).First(&existingReview).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "You have already submitted a review for this bootcamp"})
	}

	if err := reviewDB.Create(newReview).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Sorry, something went wrong", "data": err.Error()})
	}

	// Creating the user response
	userResponse := model.UserReview{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	// Create the Review Response
	reviewResponse := model.ReviewResponse{
		ID:         newReview.ID,
		BootcampID: newReview.BootcampID,
		Text:       newReview.Text,
		CreatedAt:  newReview.CreatedAt,
		UpdatedAt:  newReview.UpdatedAt,
		Title:      newReview.Title,
		Rating:     newReview.Rating,
		User:       userResponse,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Review Created!", "data": reviewResponse})
}

// ReviewResponse represents the review data to be returned in the response
type ReviewResponse struct {
	ID         uuid.UUID    `json:"id"`
	Title      string       `json:"title"`
	Text       string       `json:"text"`
	Rating     int          `json:"rating"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
	BootcampID uuid.UUID    `json:"bootcampID"`
	User       UserResponse `json:"user"`
}

// UserResponse represents the user data to be returned in the response
type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func GetReviews(c *fiber.Ctx) error {
	var reviews []model.Review
	db := database.DB.Db
	bootcampID := c.Params("bootcampId")

	bootcampUUID, err := uuid.Parse(bootcampID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid bootcamp ID"})
	}
	if err := db.Preload("User").Where("bootcamp_id = ?", bootcampUUID).Find(&reviews).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "There is no bootcamp with that id", "data": err.Error()})
	}

	// Map reviews to response struct
	var reviewResponses []ReviewResponse
	for _, review := range reviews {
		reviewResponse := ReviewResponse{
			ID:         review.ID,
			Title:      review.Title,
			Text:       review.Text,
			Rating:     review.Rating,
			CreatedAt:  review.CreatedAt,
			UpdatedAt:  review.UpdatedAt,
			BootcampID: review.BootcampID,
			User: UserResponse{
				ID:       review.User.ID,
				Username: review.User.Username,
				Email:    review.User.Email,
			},
		}
		reviewResponses = append(reviewResponses, reviewResponse)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "successful",
		"success": true,
		"data":    reviewResponses,
		"count":   len(reviewResponses),
	})
}

func GetReview(c *fiber.Ctx) error {
	reviewId := c.Params("reviewId")
	db := database.DB.Db
	var review model.Review
	if err := db.Preload("User").Where("id = ?", reviewId).First(&review).Error; err != nil {
		customError := middleware.HandleDBError(err)
		return c.Status(fiber.StatusNotFound).JSON(customError)
	}

	var reponse ReviewResponse
	reponse.ID = review.ID
	reponse.Title = review.Title
	reponse.Text = review.Text
	reponse.Rating = review.Rating
	reponse.CreatedAt = review.CreatedAt
	reponse.UpdatedAt = review.UpdatedAt
	reponse.BootcampID = review.BootcampID
	reponse.User = UserResponse{
		ID:       review.User.ID,
		Username: review.User.Username,
		Email:    review.User.Email,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": reponse})
}
func UpdateReview(c *fiber.Ctx) error {
	reviewId := c.Params("reviewId")
	db := database.DB.Db
	var review model.Review

	// Check if the review exists
	if err := db.First(&review, "id = ?", reviewId).Error; err != nil {
		customeError := middleware.HandleDBError(err)
		return c.Status(fiber.StatusNotFound).JSON(customeError)
	}

	updatereview := new(model.Review)

	// user, ok := c.Locals("user").(model.User)

	// if !ok {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "User not found in context"})
	// }
	// userID := user.ID

	if err := c.BodyParser(&updatereview); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error while checking the input", "error": err.Error()})
	}

	// Update the review fields
	review.Title = updatereview.Title
	review.Text = updatereview.Text
	review.Rating = updatereview.Rating

	// Save the updated review
	if err := db.Save(&review).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error while saving the data", "data": err.Error()})
	}

	// Prepare the response
	response := ReviewResponse{
		ID:         review.ID,
		Title:      review.Title,
		Text:       review.Text,
		Rating:     review.Rating,
		CreatedAt:  review.CreatedAt,
		UpdatedAt:  review.UpdatedAt,
		BootcampID: review.BootcampID,
		User: UserResponse{
			ID:       review.User.ID,
			Username: review.User.Username,
			Email:    review.User.Email,
		},
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Review updated successfully", "data": response})
}

func DeleteReview(c *fiber.Ctx) error {
	reviewId := c.Params("reviewId")
	db := database.DB.Db
	var review model.Review
	if err := db.First(&review, "id = ?", reviewId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "could not get the id of the review", "error": err.Error()})
	}
	if err := db.Delete(&review).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "could not delete the review", "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "review deleted successfully"})
}
