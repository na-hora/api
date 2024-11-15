package services

import (
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
