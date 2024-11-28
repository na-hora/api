package services

import (
	"na-hora/api/internal/models/company-hour/dtos"
	"na-hora/api/internal/models/company-hour/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyHourServiceInterface interface {
	CreateManyCompanyHour(hourCreate dtos.CreateCompanyHourRequestBody, companyID uuid.UUID, tx *gorm.DB) *utils.AppError
}

type CompanyHourService struct {
	companyHourRepository repositories.CompanyHourRepositoryInterface
}

func GetCompanyHourService(repo repositories.CompanyHourRepositoryInterface) CompanyHourServiceInterface {
	return &CompanyHourService{
		repo,
	}
}

func (chs *CompanyHourService) CreateManyCompanyHour(hourCreate dtos.CreateCompanyHourRequestBody, companyID uuid.UUID, tx *gorm.DB) *utils.AppError {
	var params = []dtos.CreateCompanyHourParams{}
	for _, register := range hourCreate.Registers {
		params = append(params, dtos.CreateCompanyHourParams{
			CompanyID:   companyID,
			Weekday:     register.Weekday,
			StartMinute: register.StartMinute,
			EndMinute:   register.EndMinute,
		})
	}

	err := chs.companyHourRepository.CreateMany(params, tx)
	if err != nil {
		return err
	}

	return nil
}
