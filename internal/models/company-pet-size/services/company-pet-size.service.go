package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-pet-size/dtos"
	"na-hora/api/internal/models/company-pet-size/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type CompanyPetSizeServiceInterface interface {
	Create(companyID uuid.UUID, petHairCreate dtos.CreateCompanyPetSizeRequestBody, tx *gorm.DB) *utils.AppError
	ListByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetSize, *utils.AppError)
	ListByPetTypeID(petTypeID int, tx *gorm.DB) ([]entity.CompanyPetSize, *utils.AppError)
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

	caser := cases.Title(language.BrazilianPortuguese)
	nameTitled := caser.String(petSizeCreate.Name)

	insertData = append(insertData, dtos.CreateCompanyPetSizeParams{
		Name:             nameTitled,
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

func (chs *CompanyPetSizeService) ListByPetTypeID(petTypeID int, tx *gorm.DB) ([]entity.CompanyPetSize, *utils.AppError) {
	return chs.companyPetSizeRepository.ListByPetTypeID(petTypeID, tx)
}

func (chs *CompanyPetSizeService) DeleteByID(petSizeID int, tx *gorm.DB) *utils.AppError {
	return chs.companyPetSizeRepository.DeleteByID(petSizeID, tx)
}

func (cpt *CompanyPetSizeService) UpdateByID(petSizeID int, update dtos.UpdateCompanyPetSizeParams, tx *gorm.DB) *utils.AppError {
	caser := cases.Title(language.BrazilianPortuguese)
	nameTitled := caser.String(update.Name)

	return cpt.companyPetSizeRepository.UpdateByID(petSizeID, dtos.UpdateCompanyPetSizeParams{
		Name:        nameTitled,
		Description: update.Description,
	}, tx)
}
