package dtos

import "github.com/google/uuid"

type CreateUserRequestBody struct {
	Username  string    `json:"username" gorm:"not null;unique" validate:"required"`
	Password  string    `json:"password" gorm:"not null" validate:"required"`
	CompanyID uuid.UUID `json:"companyId" gorm:"not null" validate:"required"`
}

type LoginUserRequestBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ForgotUserPasswordRequestBody struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetUserPasswordRequestBody struct {
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	Validator uuid.UUID `json:"validator" validate:"required"`
}
