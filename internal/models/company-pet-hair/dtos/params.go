package dtos

import (
	"github.com/google/uuid"
)

type CreateCompanyPetHairParams struct {
	Name             string
	Description      *string
	CompanyID        uuid.UUID
	CompanyPetTypeID int
}

type UpdateCompanyPetHairParams struct {
	Name        string
	Description *string
}
