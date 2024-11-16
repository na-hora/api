package dtos

import "github.com/go-playground/validator/v10"

type CreateCompanyPetSizeRequestBody struct {
	Name             string `json:"name" validate:"required"`
	CompanyPetTypeID int    `json:"companyPetTypeID" validate:"required"`
}

func (dto *CreateCompanyPetSizeRequestBody) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}

type UpdateCompanyPetSizeRequestBody struct {
	Name string `json:"name" validate:"required"`
}

func (dto *UpdateCompanyPetSizeRequestBody) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}
