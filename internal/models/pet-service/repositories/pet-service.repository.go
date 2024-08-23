package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/pet-service/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PetServiceRepositoryInterface interface {
	Create(companyID uuid.UUID, insert dtos.CreatePetServiceRequestBody, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError)
	Update(companyID uuid.UUID, update dtos.UpdateCompanyPetServiceParams, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError)
	CreateConfiguration(ID int, insert dtos.CreateCompanyPetServiceConfigurationParams, tx *gorm.DB) (*entity.CompanyPetServiceValue, *utils.AppError)
	UpdateConfiguration(ID int, insert dtos.UpdateCompanyPetServiceConfigurationParams, tx *gorm.DB) (*entity.CompanyPetServiceValue, *utils.AppError)
	GetByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetService, *utils.AppError)
	DeleteByID(petServiceID int, tx *gorm.DB) *utils.AppError
}

type PetServiceRepository struct {
	db *gorm.DB
}

func GetPetServiceRepository(db *gorm.DB) PetServiceRepositoryInterface {
	return &PetServiceRepository{db}
}

func (sr *PetServiceRepository) Create(
	companyID uuid.UUID,
	insert dtos.CreatePetServiceRequestBody,
	tx *gorm.DB,
) (*entity.CompanyPetService, *utils.AppError) {
	if tx == nil {
		tx = sr.db
	}

	insertValue := entity.CompanyPetService{
		CompanyID:   companyID,
		Name:        insert.Name,
		Paralellism: insert.Paralellism,
	}

	data := tx.Create(&insertValue)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &insertValue, nil
}

func (sr *PetServiceRepository) Update(
	companyID uuid.UUID,
	insert dtos.UpdateCompanyPetServiceParams,
	tx *gorm.DB,
) (*entity.CompanyPetService, *utils.AppError) {
	if tx == nil {
		tx = sr.db
	}

	updateValue := entity.CompanyPetService{
		ID:          insert.ID,
		CompanyID:   companyID,
		Name:        insert.Name,
		Paralellism: insert.Paralellism,
	}

	data := tx.Updates(&updateValue)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &updateValue, nil
}

func (sr *PetServiceRepository) CreateConfiguration(
	companyPetServiceID int,
	insert dtos.CreateCompanyPetServiceConfigurationParams,
	tx *gorm.DB,
) (*entity.CompanyPetServiceValue, *utils.AppError) {
	if tx == nil {
		tx = sr.db
	}

	insertValue := entity.CompanyPetServiceValue{
		CompanyPetServiceID: companyPetServiceID,
		CompanyPetSizeID:    insert.CompanyPetSizeID,
		CompanyPetHairID:    insert.CompanyPetHairID,
		Price:               insert.Price,
		ExecutionTime:       insert.ExecutionTime,
	}

	data := tx.Create(&insertValue)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &insertValue, nil
}

func (sr *PetServiceRepository) UpdateConfiguration(
	companyPetServiceID int,
	update dtos.UpdateCompanyPetServiceConfigurationParams,
	tx *gorm.DB,
) (*entity.CompanyPetServiceValue, *utils.AppError) {
	if tx == nil {
		tx = sr.db
	}

	updateValue := entity.CompanyPetServiceValue{
		CompanyPetServiceID: companyPetServiceID,
		ID:                  update.ID,
		Price:               update.Price,
		ExecutionTime:       update.ExecutionTime,
	}

	data := tx.Updates(&updateValue)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &updateValue, nil
}

func (sr *PetServiceRepository) GetByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetService, *utils.AppError) {
	if tx == nil {
		tx = sr.db
	}

	petService := []entity.CompanyPetService{}
	data := tx.Where("company_id = ?", companyID).Find(&petService)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return petService, nil
}

func (sr *PetServiceRepository) DeleteByID(petServiceID int, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = sr.db
	}

	data := tx.Where("id = ?", petServiceID).Delete(&entity.CompanyPetService{})
	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if data.RowsAffected == 0 {
		return &utils.AppError{
			Message:    "Pet service not found",
			StatusCode: http.StatusNotFound,
		}
	}

	return nil
}
