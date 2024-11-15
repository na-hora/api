package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-pet-hair/dtos"
	"na-hora/api/internal/models/company-pet-hair/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyPetHairServiceInterface interface {
	Create(companyID uuid.UUID, petHairCreate dtos.CreateCompanyPetHairRequestBody, tx *gorm.DB) *utils.AppError
	ListByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetHair, *utils.AppError)
}

type CompanyPetHairService struct {
	companyPetHairRepository repositories.CompanyPetHairRepositoryInterface
}

func GetCompanyPetHairService(repo repositories.CompanyPetHairRepositoryInterface) CompanyPetHairServiceInterface {
	return &CompanyPetHairService{
		repo,
	}
}

func (cphs *CompanyPetHairService) Create(
	companyID uuid.UUID,
	petHairCreate dtos.CreateCompanyPetHairRequestBody,
	tx *gorm.DB,
) *utils.AppError {
	insertData := []dtos.CreateCompanyPetHairParams{}

	insertData = append(insertData, dtos.CreateCompanyPetHairParams{
		Name:             petHairCreate.Name,
		CompanyID:        companyID,
		CompanyPetTypeID: petHairCreate.CompanyPetTypeID,
	})

	err := cphs.companyPetHairRepository.CreateMany(insertData, tx)

	if err != nil {
		return err
	}

	return nil
}

func (cphs *CompanyPetHairService) ListByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetHair, *utils.AppError) {
	hairs, err := cphs.companyPetHairRepository.ListByCompanyID(companyID, tx)
	if err != nil {
		return nil, err
	}

	return hairs, nil
}
