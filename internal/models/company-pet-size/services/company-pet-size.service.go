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
	CreateDefaultCompanyPetSizes(companyID uuid.UUID, tx *gorm.DB) *utils.AppError
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

func (chs *CompanyPetSizeService) CreateDefaultCompanyPetSizes(companyID uuid.UUID, tx *gorm.DB) *utils.AppError {
	var params = []dtos.CreateCompanyPetSizeParams{}

	defaultSizes := []string{"Pequeno", "Médio", "Grande"}

	for _, sizeName := range defaultSizes {
		params = append(params, dtos.CreateCompanyPetSizeParams{
			CompanyID: companyID,
			Name:      sizeName,
		})
	}

	err := chs.companyPetSizeRepository.CreateMany(params, tx)
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
