package dtos

import "github.com/google/uuid"

type ListPetServicesByCompanyResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CreatePetServiceResponse struct {
	ID    uuid.UUID `json:"id"`
	Token string    `json:"token"`
}
