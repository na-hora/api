package dtos

import "github.com/google/uuid"

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type LoginUserResponse struct {
	ID    uuid.UUID `json:"id"`
	Token string    `json:"token"`
}
