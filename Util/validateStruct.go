package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct using validator tags
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		// Format validation errors into readable message
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				return fmt.Errorf("%s is %s", e.Field(), e.Tag())
			}
		}
		return err
	}
	return nil
}
