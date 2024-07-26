package lib

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var instance *validator.Validate

func Validate(c *fiber.Ctx, input interface{}) *map[string]string {
	if c.BodyParser(input) != nil {
		return &map[string]string{
			"error": "unable to parse body to json",
		}
	}

	if instance == nil {
		instance = validator.New(validator.WithRequiredStructEnabled())
	}

	err := instance.Struct(input)

	if err == nil {
		return nil
	}

	errs := map[string]string{}

	t := reflect.TypeOf(input)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for index, err := range err.(validator.ValidationErrors) {
		field := t.Field(index).Tag.Get("json")
		errs[field] = err.ActualTag()
	}

	return &errs
}
