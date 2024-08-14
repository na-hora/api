package dtos

import "github.com/google/uuid"

type CreatePetServiceRequestBody struct {
	CompanyID      uuid.UUID                 `json:"companyId" gorm:"not null" validate:"required"`
	Name           string                    `json:"name" gorm:"not null" validate:"required"`
	Paralellism    int                       `json:"paralellism" gorm:"not null;default:1" validate:"required"`
	Configurations []PetServiceConfiguration `json:"configurations" validate:"required"`
}

type PetServiceConfiguration struct {
	Price            float64 `json:"price" validate:"required"`
	ExecutionTime    int     `json:"executionTime" validate:"required"`
	CompanyPetSizeID uint    `json:"companyPetSizeId" validate:"required"`
	CompanyPetHairID uint    `json:"companyPetHairId" validate:"required"`
}
