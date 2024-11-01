package dtos

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateClientRequestBody struct {
	CompanyID uuid.UUID `json:"companyId" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required"`
}

func (dto *CreateClientRequestBody) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}

type UpdateClientRequestBody struct {
	ID    uuid.UUID `json:"id" validate:"required"`
	Name  string    `json:"name" validate:"required"`
	Email string    `json:"email" validate:"required,email"`
	Phone string    `json:"phone" validate:"required"`
}

func (dto *UpdateClientRequestBody) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}
