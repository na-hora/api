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
	ListByPetTypeID(int, *gorm.DB) ([]entity.CompanyPetHair, *utils.AppError)
	DeleteByID(petHairID int, tx *gorm.DB) *utils.AppError
	UpdateByID(petHairID int, update dtos.UpdateCompanyPetHairParams, tx *gorm.DB) *utils.AppError
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
			CompanyID:        data.CompanyID,
			Name:             data.Name,
			Description:      data.Description,
			CompanyPetTypeID: data.CompanyPetTypeID,
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
	data := tx.Where("company_id = ?", companyID).Preload("CompanyPetType").Find(&companyPetHairs)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return companyPetHairs, nil
}

func (chr *CompanyPetHairRepository) ListByPetTypeID(petTypeID int, tx *gorm.DB) ([]entity.CompanyPetHair, *utils.AppError) {
	if tx == nil {
		tx = chr.db
	}

	var companyPetHairs []entity.CompanyPetHair
	data := tx.Where("company_pet_type_id = ?", petTypeID).Preload("CompanyPetType").Find(&companyPetHairs)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return companyPetHairs, nil
}

func (cpt *CompanyPetHairRepository) DeleteByID(petHairID int, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = cpt.db
	}

	data := tx.Where("id = ?", petHairID).Delete(&entity.CompanyPetHair{})
	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if data.RowsAffected == 0 {
		return &utils.AppError{
			Message:    "Pet hair not found",
			StatusCode: http.StatusNotFound,
		}
	}

	return nil
}

func (cpt *CompanyPetHairRepository) UpdateByID(petHairID int, update dtos.UpdateCompanyPetHairParams, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = cpt.db
	}

	companyPetHair := entity.CompanyPetHair{
		ID:          petHairID,
		Name:        update.Name,
		Description: update.Description,
	}

	data := tx.Updates(&companyPetHair)
	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if data.RowsAffected == 0 {
		return &utils.AppError{
			Message:    "Pet hair not found",
			StatusCode: http.StatusNotFound,
		}
	}

	return nil
}
