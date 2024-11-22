package dtos

type CreatePetServiceRequestBody struct {
	Name        string `json:"name" validate:"required"`
	Paralellism int    `json:"paralellism" validate:"required"`
	PetTypes    []int  `json:"petTypes" validate:"required"`
}

type PetServiceValuesRelateRequestBody struct {
	Relations []PetServiceValuesToRelate `json:"relations" validate:"required"`
}

type PetServiceValuesToRelate struct {
	Price            float64 `json:"price" validate:"required"`
	ExecutionTime    int     `json:"executionTime" validate:"required"`
	CompanyPetSizeID int     `json:"companyPetSizeId" validate:"required"`
	CompanyPetHairID int     `json:"companyPetHairId" validate:"required"`
}

type UpdatePetServiceRequestBody struct {
	Name        string `json:"name" validate:"required"`
	Paralellism int    `json:"paralellism" validate:"required"`
	PetTypes    []int  `json:"petTypes" validate:"required"`
}
