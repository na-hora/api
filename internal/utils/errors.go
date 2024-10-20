package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Message    string
	StatusCode int
}

func ResponseValidationErrors(err error, w http.ResponseWriter, origin string) {
	if errors, ok := err.(validator.ValidationErrors); ok {
		errorsResponse := translateErrors(errors, origin)

		ResponseJSON(w, http.StatusBadRequest, errorsResponse)
		return
	}

	ResponseJSON(w, http.StatusInternalServerError, nil)
}

func translateErrors(errors validator.ValidationErrors, origin string) map[string]string {
	errorResponses := make(map[string]string)
	for _, error := range errors {
		errorResponses[strings.ToLower(error.Field())] = fmt.Sprintf(formatError(error)+" in %s", origin)
	}

	return errorResponses
}

func formatError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "required"
	case "email":
		return "invalid"
	default:
		return "validation failed"
	}
}
