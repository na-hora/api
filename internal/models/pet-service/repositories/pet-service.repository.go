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
	RelateToType(ID int, PetTypeID int, tx *gorm.DB) *utils.AppError
	UnrelateFromType(ID int, PetTypeID int, tx *gorm.DB) *utils.AppError
	CreateConfiguration(companyPetServiceID int, insert dtos.CreateCompanyPetServiceConfigurationParams, tx *gorm.DB) (*entity.CompanyPetServiceValue, *utils.AppError)
	DeleteConfiguration(configurationID int, tx *gorm.DB) *utils.AppError
	UpdateConfiguration(companyPetServiceID int, insert dtos.UpdateCompanyPetServiceConfigurationParams, tx *gorm.DB) (*entity.CompanyPetServiceValue, *utils.AppError)
	GetByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetService, *utils.AppError)
	GetByID(ID int, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError)
	GetConfigurationBySizeAndHair(companyPetServiceID int, sizeID int, hairID int, tx *gorm.DB) (*entity.CompanyPetServiceValue, *utils.AppError)
	DeleteByID(petServiceID int, tx *gorm.DB) *utils.AppError
	GetConfigurationById(companyPetServiceID int) (*entity.CompanyPetServiceValue, *utils.AppError)
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

func (sr *PetServiceRepository) RelateToType(ID int, PetTypeID int, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = sr.db
	}

	data := tx.Create(&entity.CompanyPetServiceTypes{
		CompanyPetServiceID: ID,
		CompanyPetTypeID:    PetTypeID,
	})

	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}

func (sr *PetServiceRepository) UnrelateFromType(ID int, PetTypeID int, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = sr.db
	}

	data := tx.Delete(&entity.CompanyPetServiceTypes{
		CompanyPetServiceID: ID,
		CompanyPetTypeID:    PetTypeID,
	})

	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}

func (sr *PetServiceRepository) Update(
	companyID uuid.UUID,
	update dtos.UpdateCompanyPetServiceParams,
	tx *gorm.DB,
) (*entity.CompanyPetService, *utils.AppError) {
	if tx == nil {
		tx = sr.db
	}

	updateValue := entity.CompanyPetService{
		ID:          update.ID,
		CompanyID:   companyID,
		Name:        update.Name,
		Paralellism: update.Paralellism,
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

func (sr *PetServiceRepository) GetConfigurationById(
	companyPetServiceID int,
) (*entity.CompanyPetServiceValue, *utils.AppError) {
	petServiceValue := entity.CompanyPetServiceValue{}

	data := sr.db.Where(
		"company_pet_service_id = ?",
		companyPetServiceID,
	).First(&petServiceValue)

	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &petServiceValue, nil
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

func (sr *PetServiceRepository) DeleteConfiguration(
	configurationID int,
	tx *gorm.DB,
) *utils.AppError {
	if tx == nil {
		tx = sr.db
	}

	data := tx.Where("id = ?", configurationID).Delete(&entity.CompanyPetServiceValue{})

	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if data.RowsAffected == 0 {
		return &utils.AppError{
			Message:    "configuration not found",
			StatusCode: http.StatusNotFound,
		}
	}

	return nil
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
	data := tx.Where("company_id = ?", companyID).Preload("ServiceTypes").Preload("ServiceTypes.CompanyPetType").Find(&petService)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return petService, nil
}

func (sr *PetServiceRepository) GetByID(ID int, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError) {
	if tx == nil {
		tx = sr.db
	}

	petService := entity.CompanyPetService{}
	data := tx.Where("id = ?", ID).Preload("Configurations").Preload("ServiceTypes").First(&petService)

	if data.Error != nil {
		if data.Error != gorm.ErrRecordNotFound {
			return nil, &utils.AppError{
				Message:    data.Error.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		return nil, nil
	}

	return &petService, nil
}

func (sr *PetServiceRepository) GetConfigurationBySizeAndHair(
	companyPetServiceID int,
	sizeID int,
	hairID int,
	tx *gorm.DB,
) (*entity.CompanyPetServiceValue, *utils.AppError) {
	if tx == nil {
		tx = sr.db
	}

	petServiceValue := entity.CompanyPetServiceValue{}
	data := tx.Where(
		"company_pet_service_id = ? AND company_pet_size_id = ? AND company_pet_hair_id = ?",
		companyPetServiceID,
		sizeID,
		hairID,
	).First(&petServiceValue)

	if data.Error != nil {
		if data.Error != gorm.ErrRecordNotFound {
			return nil, &utils.AppError{
				Message:    data.Error.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		return nil, nil
	}

	return &petServiceValue, nil
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
