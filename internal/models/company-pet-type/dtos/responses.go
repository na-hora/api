package dtos

type ListPetTypesByCompanyResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListPetTypeCombinationsResponse struct {
	Hair          PetTypeCombinationsHair `json:"hair"`
	Size          PetTypeCombinationsSize `json:"size"`
	Price         *float64                `json:"price"`
	ExecutionTime *int                    `json:"executionTime"`
}

type PetTypeCombinationsHair struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PetTypeCombinationsSize struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
