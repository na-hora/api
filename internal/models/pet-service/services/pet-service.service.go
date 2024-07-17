package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/pet-service/dtos"
	"na-hora/api/internal/models/pet-service/repositories"
	"na-hora/api/internal/utils"

	"gorm.io/gorm"
)

type PetServiceInterface interface {
	CreatePetService(petServiceCreate dtos.CreatePetServiceRequestBody, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError)
}

type PetServiceService struct {
	petServiceRepository repositories.PetServiceRepositoryInterface
}

func GetPetServiceServicd(repo repositories.PetServiceRepositoryInterface, tx *gorm.DB) PetServiceInterface {
	return &PetServiceService{
		repo,
	}
}

func (ps *PetServiceService) CreatePetService(petServiceCreate dtos.CreatePetServiceRequestBody, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError) {
	petServiceCreated, err := ps.petServiceRepository.Create(petServiceCreate, tx)
	if err != nil {
		return nil, err
	}

	return petServiceCreated, nil
}
