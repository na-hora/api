package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-pet-size/dtos"
	"na-hora/api/internal/models/company-pet-size/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyPetSizeServiceInterface interface {
	Create(companyID uuid.UUID, petHairCreate dtos.CreateCompanyPetSizeRequestBody, tx *gorm.DB) *utils.AppError
	ListByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetSize, *utils.AppError)
}

type CompanyPetSizeService struct {
	companyPetSizeRepository repositories.CompanyPetSizeRepositoryInterface
}

func GetCompanyPetSizeService(repo repositories.CompanyPetSizeRepositoryInterface) CompanyPetSizeServiceInterface {
	return &CompanyPetSizeService{
		repo,
	}
}

func (cphs *CompanyPetSizeService) Create(
	companyID uuid.UUID,
	petSizeCreate dtos.CreateCompanyPetSizeRequestBody,
	tx *gorm.DB,
) *utils.AppError {
	insertData := []dtos.CreateCompanyPetSizeParams{}

	insertData = append(insertData, dtos.CreateCompanyPetSizeParams{
		Name:             petSizeCreate.Name,
		CompanyID:        companyID,
		CompanyPetTypeID: petSizeCreate.CompanyPetTypeID,
	})

	err := cphs.companyPetSizeRepository.CreateMany(insertData, tx)

	if err != nil {
		return err
	}

	return nil
}

func (chs *CompanyPetSizeService) ListByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetSize, *utils.AppError) {
	sizes, err := chs.companyPetSizeRepository.ListByCompanyID(companyID, tx)
	if err != nil {
		return nil, err
	}

	return sizes, nil
}
