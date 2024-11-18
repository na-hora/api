package dtos

type ListPetHairsByCompanyResponse struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	CompanyPetTypeID   int    `json:"companyPetTypeId"`
	CompanyPetTypeName string `json:"companyPetTypeName"`
}
