package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type AppError struct {
	Message    string
	StatusCode int
}

func ResponseValidationErrors(err error, w http.ResponseWriter) {
	if errors, ok := err.(validator.ValidationErrors); ok {
		errorsResponse := translateErrors(errors)

		ResponseJSON(w, http.StatusBadRequest, errorsResponse)
		return
	}

	ResponseJSON(w, http.StatusInternalServerError, nil)
}

func translateErrors(errors validator.ValidationErrors) map[string]string {
	errorResponses := make(map[string]string)
	for _, error := range errors {
		errorResponses[strings.ToLower(error.Field())] = formatError(error)
	}

	return errorResponses
}

func formatError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("required")
	case "email":
		return "invalid"
	default:
		return fmt.Sprintf("validation failed")
	}
}
