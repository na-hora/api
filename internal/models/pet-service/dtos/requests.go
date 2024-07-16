package dtos

import "github.com/google/uuid"

type CreatePetServiceRequestBody struct {
	CompanyID   uuid.UUID `json:"companyId" gorm:"not null" validate:"required"`
	Name        string    `json:"name" gorm:"not null" validate:"required"`
	Paralellism int       `json:"paralellism" gorm:"not null;default:1" validate:"required"`
}
