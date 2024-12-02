package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-hour/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyHourRepositoryInterface interface {
	ListByCompanyID(uuid.UUID) ([]dtos.ListHoursByCompanyIDResponse, *utils.AppError)
	CreateMany([]dtos.CreateCompanyHourParams, *gorm.DB) *utils.AppError
}

type CompanyHourRepository struct {
	db *gorm.DB
}

func GetCompanyHourRepository(db *gorm.DB) CompanyHourRepositoryInterface {
	return &CompanyHourRepository{db}
}

func (chr *CompanyHourRepository) ListByCompanyID(companyID uuid.UUID) ([]dtos.ListHoursByCompanyIDResponse, *utils.AppError) {
	var companyHours []entity.CompanyHour

	data := chr.db.Where("company_id = ?", companyID).Find(&companyHours).Order("weekday").Order("start_minute").Order("end_minute")
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	var response []dtos.ListHoursByCompanyIDResponse
	for _, hour := range companyHours {
		response = append(response, dtos.ListHoursByCompanyIDResponse{
			Weekday:     hour.Weekday,
			StartMinute: hour.StartMinute,
			EndMinute:   hour.EndMinute,
		})
	}

	return response, nil
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
