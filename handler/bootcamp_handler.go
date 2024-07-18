package handler

import (
	"log"
	// "os"

	// "strconv"
	// "strings"

	"github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"
	"github.com/juniorteye/devCamp/database"
	"github.com/juniorteye/devCamp/middleware"
	"github.com/juniorteye/devCamp/model"
	"github.com/juniorteye/devCamp/validation"
)

// CreateBootCamp godoc
// @Summary Create a new BootCamp
// @Description Create a new BootCamp
// @Tags BootCamps
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Name"
// @Param description formData string true "Description"
// @Param photo formData file true "Photo"
// @Success 201 {object} model.Bootcamp
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /bootcamp [post]
func CreateBootCamp(c *fiber.Ctx) error {
	db := database.DB.Db
	newBootcamp := new(model.Bootcamp)

	// Define the file variable outside of the if block

	// file, err := c.FormFile("photo")
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"message": "no file found", "status": "failed"})
	// }

	// mimeType := file.Header.Get("Content-Type")
	// fileType := strings.Split(mimeType, "/")[0]
	// if fileType != "image" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "please upload an image"})
	// }
	// // size, err := strconv.Atoi(os.Getenv("MAX_FILE_UPLOAD"))
	// size, err := strconv.Atoi(os.Getenv("MAX_FILE_UPLOAD"))

	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"message": "upload went wrong"})
	// }
	// maxSize := size * 1024 * 1024
	// if file.Size > int64(maxSize) {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "file size exceeds limit of 2MB"})
	// }

	// filePath := os.Getenv("FILE_UPLOAD_PATH") + file.Filename
	// err = c.SaveFile(file, filePath)
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"message": "file not saved", "status": "error", "data": err.Error()})
	// }

	//Parse the request body into the user struct
	if err := c.BodyParser(newBootcamp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "something's is wrong with input", "data": err.Error()})
	}

	// newBootcamp.Photo = filePath

	// Validate the struct
	err := validation.Validator().Struct(newBootcamp)
	if err != nil {
		formattedErros := validation.FormatValidationError(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "validation failed", "data": formattedErros})
	}
	log.Printf("the new Bootcamp is %v\n", newBootcamp)
	err = db.Create(&newBootcamp).Error
	if err != nil {
		log.Printf("the fail message is: %s\n", err.Error())
		customError := middleware.HandleDBError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(customError)
	}
	// Preload the associated User data
	err = db.Preload("User").First(&newBootcamp, newBootcamp.ID).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "could not preload user data", "data": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "bootcamp Created", "data": newBootcamp})
}

// GetAllBootCamp godoc
// @Summary Get all BootCamps
// @Description Get all BootCamps
// @Tags BootCamps
// @Produce json
// @Success 200 {array} model.Bootcamp
// @Failure 404 {object} fiber.Map
// @Router /bootcamps [get]
func GetAllBootCamp(c *fiber.Ctx) error {
	db := database.DB.Db
	var bootcamps []model.Bootcamp

	// find all bootcamps
	db.Find(&bootcamps)

	// if no bootcamp return an error
	if len(bootcamps) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "no bootcamp was found"})
	}
	return c.Status(201).JSON(fiber.Map{"status": "sucess", "data": bootcamps})
}

// GetBootCamp godoc
// @Summary Get a BootCamp by ID
// @Description Get a BootCamp by ID
// @Tags BootCamps
// @Produce json
// @Param bootcampId path string true "BootCamp ID"
// @Success 200 {object} model.Bootcamp
// @Failure 404 {object} fiber.Map
// @Router /bootcamp/{bootcampId} [get]
func GetBootCamp(c *fiber.Ctx) error {
	db := database.DB.Db
	bootcampId := c.Params("bootcampId")
	var bootcamp model.Bootcamp

	err := db.First(&bootcamp, "id = ?", bootcampId).Error
	if err != nil {
		log.Printf("the id is %s\n", err.Error())
		if err.Error() == "record not found" {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "bootcamp not found"})
		}
		customError := middleware.HandleDBError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(customError)
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "data": bootcamp})
}

// UpDateBootCamp godoc
// @Summary Update a BootCamp by ID
// @Description Update a BootCamp by ID
// @Tags BootCamps
// @Accept json
// @Produce json
// @Param bootcampId path string true "BootCamp ID"
// @Param bootcamp body model.Bootcamp true "BootCamp"
// @Success 200 {object} model.Bootcamp
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /bootcamp/{bootcampId} [put]
func UpDateBootCamp(c *fiber.Ctx) error {
	db := database.DB.Db
	bootcampId := c.Params("bootcampId")
	var bootcamp model.Bootcamp

	// Find the bootcamp by ID
	err := db.First(&bootcamp, "id = ?", bootcampId).Error
	if err != nil {
		log.Printf("the id error is %s\n", err.Error())
		if err.Error() == "record not found" {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "bootcamp not found"})
		}
		customError := middleware.HandleDBError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(customError)
	}

	// Parse the request body into the updateBootcamp struct
	updateBootcamp := new(model.Bootcamp)
	if err := c.BodyParser(updateBootcamp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "something's wrong with input", "data": err.Error()})
	}

	// Update the fields of the retrieved bootcamp with the new values
	bootcamp.Name = updateBootcamp.Name
	bootcamp.Description = updateBootcamp.Description
	bootcamp.Website = updateBootcamp.Website
	// bootcamp.Photo = updateBootcamp.Photo

	// Save the updated bootcamp back to the database
	if err := db.Save(&bootcamp).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not update bootcamp", "data": err.Error()})
	}
	// Return the updated bootcamp
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "bootcamp updated", "data": bootcamp})
}

// DeleteBootCamp godoc
// @Summary Delete a BootCamp by ID
// @Description Delete a BootCamp by ID
// @Tags BootCamps
// @Produce json
// @Param bootcampId path string true "BootCamp ID"
// @Success 200 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /bootcamp/{bootcampId} [delete]
func DeleteBootCamp(c *fiber.Ctx) error {
	db := database.DB.Db
	bootcampId := c.Params("bootcampId")
	var bootcamp model.Bootcamp

	// Attempt to find the user first to check if it exists
	if err := db.First(&bootcamp, "id = ?", bootcampId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "user not found"})
	}

	// Deleting the bootcamp with that id
	if err := db.Delete(&bootcamp, "id = ?", bootcampId).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not delete user", "data": err})
	}

	// Verify the bootcamp has been deleted
	var checkBootcamp model.Bootcamp
	if err := db.Unscoped().First(&checkBootcamp, "id = ?", bootcampId).Error; err == nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "bootcamp still exists after deletion attempt"})
	}
	// Return status Message
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "bootcamp deleted successfully"})
}
