package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-pet-type/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyPetTypeRepositoryInterface interface {
	CreateMany([]dtos.CreateCompanyPetTypeParams, *gorm.DB) *utils.AppError
	List(companyID uuid.UUID) ([]entity.CompanyPetType, *utils.AppError)
}

type CompanyPetTypeRepository struct {
	db *gorm.DB
}

func GetCompanyPetTypeRepository(db *gorm.DB) CompanyPetTypeRepositoryInterface {
	return &CompanyPetTypeRepository{db}
}

func (cpt *CompanyPetTypeRepository) CreateMany(insert []dtos.CreateCompanyPetTypeParams, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = cpt.db
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

func (cpt *CompanyPetTypeRepository) List(companyID uuid.UUID) ([]entity.CompanyPetType, *utils.AppError) {
	companyPetTypes := []entity.CompanyPetType{}
	data := cpt.db.Where("company_id = ?", companyID).Find(&companyPetTypes)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return companyPetTypes, nil
}
