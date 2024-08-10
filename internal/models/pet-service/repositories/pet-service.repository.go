package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/pet-service/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"gorm.io/gorm"
)

type PetServiceRepositoryInterface interface {
	Create(insert dtos.CreatePetServiceRequestBody, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError)
}

type PetServiceRepository struct {
	db *gorm.DB
}

func GetPetServiceRepository(db *gorm.DB) PetServiceRepositoryInterface {
	return &PetServiceRepository{db}
}

func (sr *PetServiceRepository) Create(insert dtos.CreatePetServiceRequestBody, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError) {
	if tx == nil {
		tx = sr.db
	}

	insertValue := entity.CompanyPetService{
		CompanyID:   insert.CompanyID,
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
