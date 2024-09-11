package pkg

import (
	"github.com/go-playground/validator/v10"
)

func ValidateForOnlineRegistration(err error, pageNo int) map[string]interface{} {
	errors := make(map[string]interface{})
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			field := LowercaseFirstChar(e.StructField())
			errors[field] = map[string]interface{}{
				"page":    pageNo,
				"message": getMessageForTag(e),
			}
		}
	}
	return errors
}

func ValidateForm(err error) map[string]interface{} {
	errors := make(map[string]interface{})
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			field := LowercaseFirstChar(e.StructField())
			errors[field] = map[string]interface{}{
				"message": getMessageForTag(e),
			}
		}
	}
	return errors
}

func getMessageForTag(e validator.FieldError) string {
	var fieldName = LowercaseFirstChar(e.StructField())
	switch e.Tag() {
	case "required":
		return fieldName + " is required"
	case "email":
		return fieldName + " must be a valid email address"
	case "e164":
		return fieldName + " must be a valid Phone no"
	case "uuid":
		return fieldName + " must be a valid UUID."
	case "url":
		return fieldName + " must be a valid URL."
	case "http_url":
		return fieldName + " must be a valid HTTP URL."
	case "https_url":
		return fieldName + " must be a valid HTTPS URL."
	case "date":
		return fieldName + " must be a valid date."
	case "datetime":
		return fieldName + " must be a valid datetime."
	case "ascii":
		return fieldName + " must contain only ASCII characters."
	case "alphanum":
		return fieldName + " must be alphanumeric."
	case "alpha":
		return fieldName + " must contain only alphabetic characters."
	case "number":
		return fieldName + " must be a number."
	case "len":
		return fieldName + " length must be exactly as defined by the rule."
	case "min":
		return fieldName + " length must be at least as defined by the rule."
	case "max":
		return fieldName + " length must be at most as defined by the rule."
	case "gte":
		return fieldName + " must be greater than or equal to as defined by the rule."
	case "lte":
		return fieldName + " must be less than or equal to as defined by the rule."
	case "gt":
		return fieldName + " must be greater than as defined by the rule."
	case "lt":
		return fieldName + " must be less than as defined by the rule."
	default:
		return e.Error()
	}
}
