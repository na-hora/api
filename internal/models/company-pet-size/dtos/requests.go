package dtos

import "github.com/go-playground/validator/v10"

type CreateCompanyPetSizeRequestBody struct {
	Name             string  `json:"name" validate:"required"`
	Description      *string `json:"description"`
	CompanyPetTypeID int     `json:"companyPetTypeID" validate:"required"`
}

func (dto *CreateCompanyPetSizeRequestBody) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}

type UpdateCompanyPetSizeRequestBody struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

func (dto *UpdateCompanyPetSizeRequestBody) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}
