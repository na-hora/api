package dtos

import "github.com/google/uuid"

type CreateCompanyAddressRequestBody struct {
	ZipCode      string `json:"zipCode" validate:"required"`
	CityIBGE     string `json:"cityIbge" validate:"required"`
	Neighborhood string `json:"neighborhood" validate:"required"`
	Street       string `json:"street" validate:"required"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
}

type CreateCompanyRequestBody struct {
	Name        string    `json:"name" validate:"required"`
	FantasyName string    `json:"fantasyName" validate:"required"`
	CNPJ        string    `json:"cnpj" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	Phone       string    `json:"phone" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Validator   uuid.UUID `json:"validator" validate:"required"`
	AvatarUrl   string    `json:"avatarUrl"`

	Address *CreateCompanyAddressRequestBody `json:"address"`
}
