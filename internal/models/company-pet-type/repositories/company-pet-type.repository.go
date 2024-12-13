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
	ListByCompanyID(companyID uuid.UUID) ([]entity.CompanyPetType, *utils.AppError)
	GetByID(ID int, tx *gorm.DB) (*entity.CompanyPetType, *utils.AppError)
	DeleteByID(petTypeID int, tx *gorm.DB) *utils.AppError
	UpdateByID(petTypeID int, update dtos.CreateCompanyPetTypeParams, tx *gorm.DB) *utils.AppError
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

func (cpt *CompanyPetTypeRepository) GetByID(ID int, tx *gorm.DB) (*entity.CompanyPetType, *utils.AppError) {
	if tx == nil {
		tx = cpt.db
	}

	petType := entity.CompanyPetType{}
	data := tx.Where("id = ?", ID).First(&petType)
	if data.Error != nil {
		if data.Error != gorm.ErrRecordNotFound {
			return nil, &utils.AppError{
				Message:    data.Error.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		return nil, nil
	}

	return &petType, nil
}

func (cpt *CompanyPetTypeRepository) ListByCompanyID(companyID uuid.UUID) ([]entity.CompanyPetType, *utils.AppError) {
	companyPetTypes := []entity.CompanyPetType{}
	data := cpt.db.Where("company_id = ?", companyID).Order("name ASC").Find(&companyPetTypes)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return companyPetTypes, nil
}

func (cpt *CompanyPetTypeRepository) DeleteByID(petTypeID int, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = cpt.db
	}

	data := tx.Where("id = ?", petTypeID).Delete(&entity.CompanyPetType{})
	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if data.RowsAffected == 0 {
		return &utils.AppError{
			Message:    "Pet type not found",
			StatusCode: http.StatusNotFound,
		}
	}

	return nil
}

func (cpt *CompanyPetTypeRepository) UpdateByID(petTypeID int, update dtos.CreateCompanyPetTypeParams, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = cpt.db
	}

	companyPetType := entity.CompanyPetType{
		ID:   petTypeID,
		Name: update.Name,
	}

	data := tx.Updates(&companyPetType)
	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if data.RowsAffected == 0 {
		return &utils.AppError{
			Message:    "Pet type not found",
			StatusCode: http.StatusNotFound,
		}
	}

	return nil
}
