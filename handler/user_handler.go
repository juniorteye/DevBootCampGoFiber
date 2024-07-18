package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juniorteye/devCamp/database"
	"github.com/juniorteye/devCamp/model"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	// Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User

	// find all users in the database
	db.Find(&users)

	// if no user found, return an error
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "no users found"})
	}
	return c.Status(201).JSON(fiber.Map{"status": "sucess", "data": users})
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var user model.User

	// find single user in the database by id
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "user not found"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not retrieve user", "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "data": user})
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	userId := c.Params("userId")
	var user model.User
	if err := db.First(&user, "id = ?", userId).Error; err != nil {
		if err.Error() == "record not found" {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "user not found"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not retrieve user", "data": err})
	}

	var updateUser model.User
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "invalid request body", "data": err})
	}

	// Update the user with the new data
	user.Username = updateUser.Username
	user.Email = updateUser.Email

	// Save the updated user back to the database
	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not update user", "data": err})
	}
	// Return the updated user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user updated", "data": user})
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB.Db
	userId := c.Params("userId")
	var user model.User

	// Attempt to find the user first to check if it exists
	if err := db.First(&user, "id = ?", userId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "user not found"})
	}

	// Delete the user
	if err := db.Delete(&user, "id = ?", userId).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not delete user", "data": err})
	}

	// Return success message
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user deleted successfully"})
}
