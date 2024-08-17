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
	CreateDefaultCompanyPetHairs(companyID uuid.UUID, tx *gorm.DB) *utils.AppError
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

func (chs *CompanyPetHairService) CreateDefaultCompanyPetHairs(companyID uuid.UUID, tx *gorm.DB) *utils.AppError {
	var params = []dtos.CreateCompanyPetHairParams{}

	defaultHairs := []string{"Curto", "Médio", "Longo"}

	for _, hairName := range defaultHairs {
		params = append(params, dtos.CreateCompanyPetHairParams{
			CompanyID: companyID,
			Name:      hairName,
		})
	}

	err := chs.companyPetHairRepository.CreateMany(params, tx)
	if err != nil {
		return err
	}

	return nil
}

func (chs *CompanyPetHairService) ListByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetHair, *utils.AppError) {
	hairs, err := chs.companyPetHairRepository.ListByCompanyID(companyID, tx)
	if err != nil {
		return nil, err
	}

	return hairs, nil
}
