package dtos

type CreateCompanyAddressRequestBody struct {
	ZipCode      string `json:"zipCode" validate:"required"`
	CityID       uint   `json:"cityId" validate:"required"`
	Neighborhood string `json:"neighborhood" validate:"required"`
	Street       string `json:"street" validate:"required"`
	Number       uint   `json:"number" validate:"required"`
	Complement   string `json:"complement"`
}

type CreateCompanyRequestBody struct {
	Name        string `json:"name" validate:"required"`
	FantasyName string `json:"fantasyName" validate:"required"`
	CNPJ        string `json:"cnpj" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone" validate:"required"`
	Password    string `json:"password" validate:"required"`
	AvatarUrl   string `json:"avatarUrl"`

	Address *CreateCompanyAddressRequestBody `json:"address"`
}
