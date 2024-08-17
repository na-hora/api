package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-pet-hair/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyPetHairRepositoryInterface interface {
	CreateMany([]dtos.CreateCompanyPetHairParams, *gorm.DB) *utils.AppError
	ListByCompanyID(uuid.UUID, *gorm.DB) ([]entity.CompanyPetHair, *utils.AppError)
}

type CompanyPetHairRepository struct {
	db *gorm.DB
}

func GetCompanyPetHairRepository(db *gorm.DB) CompanyPetHairRepositoryInterface {
	return &CompanyPetHairRepository{db}
}

func (chr *CompanyPetHairRepository) CreateMany(insert []dtos.CreateCompanyPetHairParams, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = chr.db
	}

	var treatedInserts []entity.CompanyPetHair
	total := 0

	for _, data := range insert {
		treatedInserts = append(treatedInserts, entity.CompanyPetHair{
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

func (chr *CompanyPetHairRepository) ListByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetHair, *utils.AppError) {
	if tx == nil {
		tx = chr.db
	}

	var companyPetHairs []entity.CompanyPetHair
	data := tx.Where("company_id = ?", companyID).Find(&companyPetHairs)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return companyPetHairs, nil
}
