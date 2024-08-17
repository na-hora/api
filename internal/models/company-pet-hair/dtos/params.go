package dtos

import (
	"github.com/google/uuid"
)

type CreateCompanyPetHairParams struct {
	Name      string
	CompanyID uuid.UUID
}
