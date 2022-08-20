package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewCustomValidator(validator *validator.Validate) CustomValidator {
	cv := CustomValidator{
		Validator: validator,
	}
	return cv
}

func (c *CustomValidator) Validate(i interface{}) error {
	return c.Validator.Struct(i)
}

func (c *CustomValidator) Message(error interface{}) string {
	var message string
	for _, err := range error.(validator.ValidationErrors) {
		if err.Tag() == "gt" {
			message = fmt.Sprintf("field %s need min %s", err.Field(), err.Param())
			//fmt.Println(err.Namespace())
			//fmt.Println(err.Field())
			//fmt.Println(err.StructNamespace())
			//fmt.Println(err.StructField())
			//fmt.Println(err.Tag())
			//fmt.Println(err.ActualTag())
			//fmt.Println(err.Kind())
			//fmt.Println(err.Type())
			//fmt.Println(err.Value())
			fmt.Println(err.Param())
		} else {
			message = fmt.Sprintf("field %s is %s", err.Field(), err.Tag())
		}

	}
	return message
}
