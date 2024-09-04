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
		return fieldName + " is required<br/>"
	case "email":
		return fieldName + " must be a valid email address<br/>"
	case "e164":
		return fieldName + " must be a valid Phone no<br/>"
	case "uuid":
		return fieldName + " must be a valid UUID.<br/>"
	case "url":
		return fieldName + " must be a valid URL.<br/>"
	case "http_url":
		return fieldName + " must be a valid HTTP URL.<br/>"
	case "https_url":
		return fieldName + " must be a valid HTTPS URL.<br/>"
	case "date":
		return fieldName + " must be a valid date.<br/>"
	case "datetime":
		return fieldName + " must be a valid datetime.<br/>"
	case "ascii":
		return fieldName + " must contain only ASCII characters.<br/>"
	case "alphanum":
		return fieldName + " must be alphanumeric.<br/>"
	case "alpha":
		return fieldName + " must contain only alphabetic characters.<br/>"
	case "number":
		return fieldName + " must be a number.<br/>"
	case "len":
		return fieldName + " length must be exactly as defined by the rule.<br/>"
	case "min":
		return fieldName + " length must be at least as defined by the rule.<br/>"
	case "max":
		return fieldName + " length must be at most as defined by the rule.<br/>"
	case "gte":
		return fieldName + " must be greater than or equal to as defined by the rule.<br/>"
	case "lte":
		return fieldName + " must be less than or equal to as defined by the rule.<br/>"
	case "gt":
		return fieldName + " must be greater than as defined by the rule.<br/>"
	case "lt":
		return fieldName + " must be less than as defined by the rule.<br/>"
	default:
		return e.Error()
	}
}
