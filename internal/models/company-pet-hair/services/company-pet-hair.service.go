package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-pet-hair/dtos"
	"na-hora/api/internal/models/company-pet-hair/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type CompanyPetHairServiceInterface interface {
	Create(companyID uuid.UUID, petHairCreate dtos.CreateCompanyPetHairRequestBody, tx *gorm.DB) *utils.AppError
	ListByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetHair, *utils.AppError)
	ListByPetTypeID(petTypeID int, tx *gorm.DB) ([]entity.CompanyPetHair, *utils.AppError)
	DeleteByID(petHairID int, tx *gorm.DB) *utils.AppError
	UpdateByID(petHairID int, update dtos.UpdateCompanyPetHairParams, tx *gorm.DB) *utils.AppError
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

	caser := cases.Title(language.BrazilianPortuguese)
	nameTitled := caser.String(petHairCreate.Name)

	insertData = append(insertData, dtos.CreateCompanyPetHairParams{
		Name:             nameTitled,
		Description:      petHairCreate.Description,
		CompanyID:        companyID,
		CompanyPetTypeID: petHairCreate.CompanyPetTypeID,
	})

	err := cphs.companyPetHairRepository.CreateMany(insertData, tx)

	if err != nil {
		return err
	}

	return nil
}

func (chs *CompanyPetHairService) ListByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetHair, *utils.AppError) {
	return chs.companyPetHairRepository.ListByCompanyID(companyID, tx)
}

func (chs *CompanyPetHairService) ListByPetTypeID(petTypeID int, tx *gorm.DB) ([]entity.CompanyPetHair, *utils.AppError) {
	return chs.companyPetHairRepository.ListByPetTypeID(petTypeID, tx)
}

func (chs *CompanyPetHairService) DeleteByID(petHairID int, tx *gorm.DB) *utils.AppError {
	return chs.companyPetHairRepository.DeleteByID(petHairID, tx)
}

func (cpt *CompanyPetHairService) UpdateByID(petHairID int, update dtos.UpdateCompanyPetHairParams, tx *gorm.DB) *utils.AppError {
	caser := cases.Title(language.BrazilianPortuguese)
	nameTitled := caser.String(update.Name)

	return cpt.companyPetHairRepository.UpdateByID(petHairID, dtos.UpdateCompanyPetHairParams{
		Name:        nameTitled,
		Description: update.Description,
	}, tx)
}
