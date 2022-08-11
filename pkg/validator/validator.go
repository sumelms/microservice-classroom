package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func (v Validator) Validate(i interface{}) error {
	err := v.validator.RegisterValidation("classroom_format", ValidateClassroomFormat)
	if err != nil {
		return fmt.Errorf("unable to register classroom_format validator")
	}

	return v.validator.Struct(i)
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// ValidateClassroomFormat validates classroom format value
func ValidateClassroomFormat(fl validator.FieldLevel) bool {
	return fl.Field().String() == "online" || fl.Field().String() == "hybrid"
}
