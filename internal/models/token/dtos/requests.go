package dtos

type GenerateTokenRequestBody struct {
	Note string `json:"note" validate:"required"`
}
