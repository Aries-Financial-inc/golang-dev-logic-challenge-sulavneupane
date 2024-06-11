package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
)

func ParseValidationError(errs ...error) string {
	var out []string
	for _, err := range errs {
		switch typedError := any(err).(type) {
		case validator.ValidationErrors:
			for _, e := range typedError {
				out = append(out, parseFieldError(e))
			}
		case *json.UnmarshalTypeError:
			out = append(out, parseMarshallingError(*typedError))
		default:
			out = append(out, err.Error())
		}
	}
	return strings.Join(out, ", ")
}

func parseFieldError(e validator.FieldError) string {
	fieldPrefix := fmt.Sprintf("%s", e.Field())
	tag := strings.Split(e.Tag(), "|")[0]
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", fieldPrefix)
	case "min":
		return fmt.Sprintf("%s must be %s characters long", fieldPrefix, e.Param())
	case "max":
		return fmt.Sprintf("%s cannot exceed %s characters", fieldPrefix, e.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fieldPrefix)
	case "uuid":
		return fmt.Sprintf("%s must be a valid UUID", fieldPrefix)
	case "required_without":
		return fmt.Sprintf("%s is required if %s is not supplied", fieldPrefix, e.Param())
	case "lt", "ltfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be less than %s", fieldPrefix, param)
	case "gt", "gtfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be greater than %s", fieldPrefix, param)
	default:
		english := en.New()
		translator := ut.New(english, english)
		if translatorInstance, found := translator.GetTranslator("en"); found {
			return e.Translate(translatorInstance)
		} else {
			return fmt.Errorf("%v", e).Error()
		}
	}
}
func parseMarshallingError(e json.UnmarshalTypeError) string {
	return fmt.Sprintf("The field %s must be a %s", e.Field, e.Type.String())
}
