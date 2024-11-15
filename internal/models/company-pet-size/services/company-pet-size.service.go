package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-pet-size/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyPetSizeServiceInterface interface {
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

func (chs *CompanyPetSizeService) ListByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetSize, *utils.AppError) {
	sizes, err := chs.companyPetSizeRepository.ListByCompanyID(companyID, tx)
	if err != nil {
		return nil, err
	}

	return sizes, nil
}
