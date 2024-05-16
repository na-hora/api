package dtos

type CreateCompanyRequestBody struct {
	Name        string `json:"name" gorm:"not null;unique"`
	FantasyName string `json:"fantasyName" gorm:"not null"`
	CNPJ        string `json:"cnpj" gorm:"not null;unique"`
	Email       string `json:"email" gorm:"not null"`
	Phone       string `json:"phone" gorm:"not null"`
	AvatarUrl   string `json:"avatarUrl"`
}
