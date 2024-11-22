package dtos

type ListPetServicesByCompanyResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetPetServiceByIDResponse struct {
	ID             int                       `json:"id"`
	Name           string                    `json:"name"`
	Paralellism    int                       `json:"paralellism"`
	Configurations []PetServiceConfiguration `json:"configurations"`
	PetTypes       []int                     `json:"petTypes"`
}

type PetServiceConfiguration struct {
	ID               int     `json:"id"`
	CompanyPetHairID int     `json:"companyPetHairId"`
	CompanyPetSizeID int     `json:"companyPetSizeId"`
	Price            float64 `json:"price"`
	ExecutionTime    int     `json:"executionTime"`
}

type CreatePetServiceResponse struct {
	ID int `json:"id"`
}

type UpdatePetServiceResponse struct {
	ID int `json:"id"`
}
