package helper

import (
	"fmt"

	"github.com/go-playground/validator"
)


type CustomValidator struct {
    Validator *validator.Validate
  }


func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return fmt.Errorf("%s", validationErrors.Error())
	}
  return nil
}