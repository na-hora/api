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
	DeleteByID(petSizeID int, tx *gorm.DB) *utils.AppError
	UpdateByID(petSizeID int, update dtos.UpdateCompanyPetSizeParams, tx *gorm.DB) *utils.AppError
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
		Description:      petSizeCreate.Description,
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
	return chs.companyPetSizeRepository.ListByCompanyID(companyID, tx)
}

func (chs *CompanyPetSizeService) DeleteByID(petSizeID int, tx *gorm.DB) *utils.AppError {
	return chs.companyPetSizeRepository.DeleteByID(petSizeID, tx)
}

func (cpt *CompanyPetSizeService) UpdateByID(petSizeID int, update dtos.UpdateCompanyPetSizeParams, tx *gorm.DB) *utils.AppError {
	return cpt.companyPetSizeRepository.UpdateByID(petSizeID, update, tx)
}
