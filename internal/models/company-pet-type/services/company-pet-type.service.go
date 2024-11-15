package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-pet-type/dtos"
	"na-hora/api/internal/models/company-pet-type/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type CompanyPetTypeServiceInterface interface {
	CreatePetType(companyID uuid.UUID, name string, tx *gorm.DB) *utils.AppError
	GetByCompanyID(companyID uuid.UUID) ([]entity.CompanyPetType, *utils.AppError)
	DeleteByID(petTypeID int, tx *gorm.DB) *utils.AppError
	UpdateByID(petTypeID int, update dtos.CreateCompanyPetTypeParams, tx *gorm.DB) *utils.AppError
}

type CompanyPetTypeService struct {
	companyPetTypeRepository repositories.CompanyPetTypeRepositoryInterface
}

func GetCompanyPetTypeService(repo repositories.CompanyPetTypeRepositoryInterface) CompanyPetTypeServiceInterface {
	return &CompanyPetTypeService{
		repo,
	}
}

func (cpt *CompanyPetTypeService) CreatePetType(companyID uuid.UUID, name string, tx *gorm.DB) *utils.AppError {
	insertData := []dtos.CreateCompanyPetTypeParams{}

	caser := cases.Title(language.BrazilianPortuguese)
	nameTitled := caser.String(name)

	insertData = append(insertData, dtos.CreateCompanyPetTypeParams{
		Name:      nameTitled,
		CompanyID: companyID,
	})

	err := cpt.companyPetTypeRepository.CreateMany(insertData, tx)
	if err != nil {
		return err
	}

	return nil
}

func (cpt *CompanyPetTypeService) GetByCompanyID(companyID uuid.UUID) ([]entity.CompanyPetType, *utils.AppError) {
	return cpt.companyPetTypeRepository.List(companyID)
}

func (cpt *CompanyPetTypeService) DeleteByID(petTypeID int, tx *gorm.DB) *utils.AppError {
	return cpt.companyPetTypeRepository.DeleteByID(petTypeID, tx)
}

func (cpt *CompanyPetTypeService) UpdateByID(petTypeID int, update dtos.CreateCompanyPetTypeParams, tx *gorm.DB) *utils.AppError {
	caser := cases.Title(language.BrazilianPortuguese)
	nameTitled := caser.String(update.Name)

	update.Name = nameTitled

	return cpt.companyPetTypeRepository.UpdateByID(petTypeID, update, tx)
}
