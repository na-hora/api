package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-hour/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"gorm.io/gorm"
)

type CompanyHourRepositoryInterface interface {
	CreateMany([]dtos.CreateCompanyHourParams, *gorm.DB) *utils.AppError
}

type CompanyHourRepository struct {
	db *gorm.DB
}

func GetCompanyHourRepository(db *gorm.DB) CompanyHourRepositoryInterface {
	return &CompanyHourRepository{db}
}

func (chr *CompanyHourRepository) CreateMany(insert []dtos.CreateCompanyHourParams, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = chr.db
	}

	var treatedInserts []entity.CompanyHour
	total := 0

	for _, data := range insert {
		treatedInserts = append(treatedInserts, entity.CompanyHour{
			CompanyID:   data.CompanyID,
			Weekday:     data.Weekday,
			StartMinute: data.StartMinute,
			EndMinute:   data.EndMinute,
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
