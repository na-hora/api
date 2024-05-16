package dtos

import "github.com/google/uuid"

type CreateCompanyResponse struct {
	ID uuid.UUID `json:"id"`
}
