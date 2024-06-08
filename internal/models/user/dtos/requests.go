package dtos

import "github.com/google/uuid"

type CreateUserRequestBody struct {
	Username  string    `json:"username" gorm:"not null;unique" validate:"required"`
	Password  string    `json:"password" gorm:"not null" validate:"required"`
	CompanyID uuid.UUID `json:"companyId" gorm:"not null" validate:"required"`
}
