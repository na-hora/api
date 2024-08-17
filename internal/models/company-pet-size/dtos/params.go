package dtos

import (
	"github.com/google/uuid"
)

type CreateCompanyPetSizeParams struct {
	Name      string
	CompanyID uuid.UUID
}
