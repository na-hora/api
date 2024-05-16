package utils

import (
	"github.com/klassmann/cpfcnpj"

	"net/http"
)

func ValidateCNPJ(cnpj string) *AppError {
	r := cpfcnpj.ValidateCNPJ(cnpj)

	if !r {
		return &AppError{
			Message:    "invalid CNPJ",
			StatusCode: http.StatusBadRequest,
		}
	}

	return nil
}
