package dtos

type ListPetServicesByCompanyResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CreatePetServiceResponse struct {
	ID uint `json:"id"`
}
