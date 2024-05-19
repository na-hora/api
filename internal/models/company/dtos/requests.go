package dtos

type CreateCompanyRequestBody struct {
	Name        string `json:"name" gorm:"not null;unique" validate:"required"`
	FantasyName string `json:"fantasyName" gorm:"not null" validate:"required"`
	CNPJ        string `json:"cnpj" gorm:"not null;unique" validate:"required"`
	Email       string `json:"email" gorm:"not null" validate:"required,email"`
	Phone       string `json:"phone" gorm:"not null" validate:"required"`
	AvatarUrl   string `json:"avatarUrl"`
}
