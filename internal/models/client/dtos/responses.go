package dtos

import (
	"github.com/google/uuid"
)

type ListClientsResponse struct {
	Clients []Client `json:"clients"`
}

type CreateClientResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CompanyID uuid.UUID `json:"companyId"`
}

type Client struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CompanyID uuid.UUID `json:"companyId"`
}
