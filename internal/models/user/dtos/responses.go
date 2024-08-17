package dtos

import "github.com/google/uuid"

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type CompanyPetSizesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CompanyPetHairsResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LoginCompanyData struct {
	ID          uuid.UUID                 `json:"id"`
	FantasyName string                    `json:"fantasyName"`
	AvatarURL   string                    `json:"avatarUrl"`
	PetSizes    []CompanyPetSizesResponse `json:"petSizes"`
	PetHairs    []CompanyPetHairsResponse `json:"petHairs"`
}

type LoginUserResponse struct {
	ID      uuid.UUID        `json:"id"`
	Token   string           `json:"token"`
	Company LoginCompanyData `json:"company"`
}
