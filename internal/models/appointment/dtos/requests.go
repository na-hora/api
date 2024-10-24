package dtos

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateAppointmentsRequestBody struct {
	CompanyID                uuid.UUID `json:"companyId" validate:"required"`
	ClientID                 uuid.UUID `json:"clientId" validate:"required"`
	CompanyPetServiceValueID int       `json:"companyPetServiceValueId" validate:"required"`
	StartTime                time.Time `json:"startTime" validate:"required"`
	PetName                  string    `json:"petName"`
	PaymentMode              string    `json:"paymentMode"`
	Note                     string    `json:"note"`
}

func (dto *CreateAppointmentsRequestBody) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}
