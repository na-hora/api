package dtos

type CreateCompanyPetServiceConfigurationParams struct {
	Price               float64
	ExecutionTime       int
	CompanyPetServiceID int
	CompanyPetSizeID    int
	CompanyPetHairID    int
}

type UpdateCompanyPetServiceParams struct {
	ID          int
	Name        string
	Paralellism int
}

type UpdateCompanyPetServiceConfigurationParams struct {
	ID            int
	Price         float64
	ExecutionTime int
}
