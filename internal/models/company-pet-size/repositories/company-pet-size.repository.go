package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-pet-size/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyPetSizeRepositoryInterface interface {
	CreateMany([]dtos.CreateCompanyPetSizeParams, *gorm.DB) *utils.AppError
	ListByCompanyID(uuid.UUID, *gorm.DB) ([]entity.CompanyPetSize, *utils.AppError)
}

type CompanyPetSizeRepository struct {
	db *gorm.DB
}

func GetCompanyPetSizeRepository(db *gorm.DB) CompanyPetSizeRepositoryInterface {
	return &CompanyPetSizeRepository{db}
}

func (chr *CompanyPetSizeRepository) CreateMany(insert []dtos.CreateCompanyPetSizeParams, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = chr.db
	}

	var treatedInserts []entity.CompanyPetSize
	total := 0

	for _, data := range insert {
		treatedInserts = append(treatedInserts, entity.CompanyPetSize{
			CompanyID: data.CompanyID,
			Name:      data.Name,
		})

		total = total + 1
	}

	data := tx.CreateInBatches(treatedInserts, total)
	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}

func (chr *CompanyPetSizeRepository) ListByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetSize, *utils.AppError) {
	if tx == nil {
		tx = chr.db
	}

	var companyPetSizes []entity.CompanyPetSize
	data := tx.Where("company_id = ?", companyID).Find(&companyPetSizes)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return companyPetSizes, nil
}
