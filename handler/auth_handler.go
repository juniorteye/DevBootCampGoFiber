package handler

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/juniorteye/devCamp/database"
	"github.com/juniorteye/devCamp/model"
	mailer "github.com/juniorteye/devCamp/utils"
	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordRequest struct {
	Email           string `json:"email"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

func SignUp(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	// Parse the request body into the user struct
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot parse JSON", "data": err.Error()})
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to hash password", "data": err.Error()})
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err.Error()})
	}

	// Create JWT token claims
	expireDuration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	claims := jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * 24 * time.Duration(expireDuration)).Unix(),
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to sign token", "data": err.Error()})
	}

	cookieExpireDuration, _ := strconv.Atoi(os.Getenv("JWT_COOKIE_EXPIRE"))
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24 * time.Duration(cookieExpireDuration)),
		HTTPOnly: true,
		SameSite: "lax",
	})

	// Return the generated token
	return c.JSON(fiber.Map{"token": t})
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var dbUser model.User
	db := database.DB.Db
	if err := db.Where("username = ?", input.Username).First(&dbUser).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Replace "your_secret_key" with os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"userID": dbUser.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 20).Unix(), // 20 days expiry
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET"))) // Use the secret key from environment variable
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
	}

	return c.JSON(fiber.Map{"token": t})
}

func ChangePassword(c *fiber.Ctx) error {
	db := database.DB.Db

	// Parse the request body into ChangePasswordRequest
	var req ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "The input is wrong"})
	}

	// Fetch the user from the database using the email
	var dbUser model.User
	if err := db.Where("email = ?", req.Email).First(&dbUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}

	// Compare the provided current password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(req.CurrentPassword)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Check your password and input it again"})
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to hash password", "data": err.Error()})
	}

	// Update the user's password in the database
	dbUser.Password = string(hashedPassword)
	if err := db.Save(&dbUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to update password", "data": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Password updated successfully"})
}

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	NewPassword string `json:"new_password"`
}

func ForgetPassword(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User

	// Get email from request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request body"})
	}

	if err := db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}

	// Generate the token to be sent
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Email body
	emailBody := "Click the link to reset your password: <a href='http://yourdomain.com/reset-password?token=" + t + "'>Reset Password</a>"

	// Send the token to the user
	if err := mailer.SendMail(user.Email, "Password Reset", emailBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to send email", "data": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Password reset email sent"})
}

func ResetPassword(c *fiber.Ctx) error {
	db := database.DB.Db

	// Get the token from the query parameters
	tokenString := c.Query("token")
	if tokenString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Token is required"})
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid or expired token"})
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid token"})
	}

	// Get the username from the token claims
	username, ok := claims["username"].(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid token claims"})
	}

	//Find the user by username
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	// Parse the request body to get the new password
	var request ResetPasswordRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request body"})
	}

	// Hash the new user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to hash password", "data": err.Error()})
	}

	// Update the user's password
	user.Password = string(hashedPassword)
	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Password updated successfully"})
}
