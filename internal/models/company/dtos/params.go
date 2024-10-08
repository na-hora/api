package dtos

type CreateCompanyAddressParams struct {
	ZipCode      string `json:"zipCode" `
	CityID       uint   `json:"cityId" `
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
}
