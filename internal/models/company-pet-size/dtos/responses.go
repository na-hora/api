package dtos

type ListPetSizesByCompanyResponse struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	CompanyPetTypeID   int    `json:"companyPetTypeId"`
	CompanyPetTypeName string `json:"companyPetTypeName"`
}
