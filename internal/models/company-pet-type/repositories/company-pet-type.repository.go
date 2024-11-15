package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-pet-type/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"gorm.io/gorm"
)

type CompanyPetTypeRepositoryInterface interface {
	CreateMany([]dtos.CreateCompanyPetTypeParams, *gorm.DB) *utils.AppError
}

type CompanyPetTypeRepository struct {
	db *gorm.DB
}

func GetCompanyPetTypeRepository(db *gorm.DB) CompanyPetTypeRepositoryInterface {
	return &CompanyPetTypeRepository{db}
}

func (this *CompanyPetTypeRepository) CreateMany(insert []dtos.CreateCompanyPetTypeParams, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = this.db
	}

	var treatedInserts []entity.CompanyPetType
	total := 0

	for _, data := range insert {
		treatedInserts = append(treatedInserts, entity.CompanyPetType{
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
