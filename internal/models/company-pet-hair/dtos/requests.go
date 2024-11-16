package dtos

import "github.com/go-playground/validator/v10"

type CreateCompanyPetHairRequestBody struct {
	Name             string `json:"name" validate:"required"`
	CompanyPetTypeID int    `json:"companyPetTypeID" validate:"required"`
}

func (dto *CreateCompanyPetHairRequestBody) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}

type UpdateCompanyPetHairRequestBody struct {
	Name string `json:"name" validate:"required"`
}

func (dto *UpdateCompanyPetHairRequestBody) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}
