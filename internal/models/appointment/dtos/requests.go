package dtos

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ListAppointmentsRequestParams struct {
	StartDate time.Time `json:"startDate" validate:"required"`
	Test      string    `json:"test" validate:"required"`
}

func (dto *ListAppointmentsRequestParams) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}

type CreateAppointmentsRequestBody struct {
	StartDate time.Time `json:"startDate" validate:"required"`
	Test      string    `json:"test" validate:"required"`
}

func (dto *CreateAppointmentsRequestBody) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}
