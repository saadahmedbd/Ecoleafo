package util

import (
	"errors"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func init() {
	validate = validator.New()

	// Register custom validators
	validate.RegisterValidation("phone", validatePhone)
}

// ValidateStruct validates a struct using validator tags

// formatValidationError formats validation error messages
func formatValidationError(err validator.FieldError) string {
	field := err.Field()
	tag := err.Tag()

	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email"
	case "min":
		return field + " must be at least " + err.Param() + " characters"
	case "max":
		return field + " must not exceed " + err.Param() + " characters"
	case "phone":
		return field + " must be a valid phone number"
	case "url":
		return field + " must be a valid URL"
	default:
		return field + " is invalid"
	}
}

// validatePhone validates Bangladesh phone numbers
func validatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	// Remove spaces
	phone = strings.ReplaceAll(phone, " ", "")

	// Bangladesh phone pattern: +8801XXXXXXXXX or 01XXXXXXXXX
	phoneRegex := regexp.MustCompile(`^(\+880|0)?1[3-9]\d{8}$`)
	return phoneRegex.MatchString(phone)
}

// ValidateEmail validates email format
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	// Check for at least one uppercase letter
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Check for at least one lowercase letter
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	// Check for at least one number
	if !regexp.MustCompile(`\d`).MatchString(password) {
		return errors.New("password must contain at least one number")
	}

	return nil
}
