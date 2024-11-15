package dtos

import (
	"github.com/google/uuid"
)

type CreateCompanyPetTypeParams struct {
	Name      string
	CompanyID uuid.UUID
}
