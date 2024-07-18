package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate *validator.Validate

// Initialize validator
func InitValidator() {
	validate = validator.New()
	RegisterCustomValidations()
}

// Register custom validation functions
func RegisterCustomValidations() {
	validate.RegisterValidation("url", validateURL)
	validate.RegisterValidation("email", validateEmail)
	validate.RegisterValidation("rating", validateRating)
}

// Custom validation for URL
func validateURL(fl validator.FieldLevel) bool {
	url := fl.Field().String()
	regex := `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	re := regexp.MustCompile(regex)
	return re.MatchString(url)
}

// Custom validation for Email
func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	regex := `^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

// Custom validation for Rating
func validateRating(fl validator.FieldLevel) bool {
	rating := fl.Field().Int()
	return rating >= 1 && rating <= 5
}

// Expose the validator instance
func Validator() *validator.Validate {
	return validate
}

func FormatValidationError(err error) map[string]string {
	validationErrors := err.(validator.ValidationErrors)
	formattedErrors := make(map[string]string)

	for _, validationErr := range validationErrors {
		field := strings.ToLower(validationErr.Field())
		tag := validationErr.Tag()
		formattedErrors[field] = fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", field, tag)
	}
	return formattedErrors
}
