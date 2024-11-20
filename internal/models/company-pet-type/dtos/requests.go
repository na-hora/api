package dtos

type CreatePetTypeRequestBody struct {
	Name string `json:"name" validate:"required"`
}
