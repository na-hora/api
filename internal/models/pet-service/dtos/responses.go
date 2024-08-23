package dtos

type ListPetServicesByCompanyResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreatePetServiceResponse struct {
	ID int `json:"id"`
}

type UpdatePetServiceResponse struct {
	ID int `json:"id"`
}
