package dtos

type CreatePetServiceRequestBody struct {
	Name           string                            `json:"name" validate:"required"`
	Paralellism    int                               `json:"paralellism" validate:"required"`
	Configurations []PetServiceConfigurationToCreate `json:"configurations" validate:"required"`
}

type PetServiceConfigurationToCreate struct {
	Price            float64 `json:"price" validate:"required"`
	ExecutionTime    int     `json:"executionTime" validate:"required"`
	CompanyPetSizeID int     `json:"companyPetSizeId" validate:"required"`
	CompanyPetHairID int     `json:"companyPetHairId" validate:"required"`
}

type UpdatePetServiceRequestBody struct {
	Name           string                            `json:"name" validate:"required"`
	Paralellism    int                               `json:"paralellism" validate:"required"`
	Configurations []PetServiceConfigurationToUpdate `json:"configurations"`
}

type PetServiceConfigurationToUpdate struct {
	CompanyPetServiceValueID int     `json:"companyPetServiceValueId" validate:"required"`
	Price                    float64 `json:"price" validate:"required"`
	ExecutionTime            int     `json:"executionTime" validate:"required"`
}
