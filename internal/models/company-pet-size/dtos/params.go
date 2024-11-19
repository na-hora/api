package dtos

import (
	"github.com/google/uuid"
)

type CreateCompanyPetSizeParams struct {
	Name             string
	Description      *string
	CompanyID        uuid.UUID
	CompanyPetTypeID int
}

type UpdateCompanyPetSizeParams struct {
	Name        string
	Description *string
}
