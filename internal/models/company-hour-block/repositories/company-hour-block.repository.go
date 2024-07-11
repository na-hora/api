package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company-hour-block/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"gorm.io/gorm"
)

type CompanyHourBlockRepositoryInterface interface {
	CreateMany([]dtos.CreateCompanyHourBlockParams, *gorm.DB) *utils.AppError
}

type CompanyHourBlockRepository struct {
	db *gorm.DB
}

func GetCompanyHourBlockRepository(db *gorm.DB) CompanyHourBlockRepositoryInterface {
	return &CompanyHourBlockRepository{db}
}

func (chr *CompanyHourBlockRepository) CreateMany(insert []dtos.CreateCompanyHourBlockParams, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = chr.db
	}

	var treatedInserts []entity.CompanyHourBlock
	total := 0

	for _, data := range insert {
		treatedInserts = append(treatedInserts, entity.CompanyHourBlock{
			CompanyID: data.CompanyID,
			Day:       data.Day,
			StartHour: data.StartHour,
			EndHour:   data.EndHour,
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
