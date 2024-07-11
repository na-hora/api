package services

import (
	"na-hora/api/internal/models/company-hour-block/dtos"
	"na-hora/api/internal/models/company-hour-block/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyHourBlockServiceInterface interface {
	CreateManyCompanyHourBlock(hourBlockCreate dtos.CreateCompanyHourBlockRequestBody, companyID uuid.UUID, tx *gorm.DB) *utils.AppError
}

type CompanyHourBlockService struct {
	companyHourBlockRepository repositories.CompanyHourBlockRepositoryInterface
}

func GetCompanyHourBlockService(repo repositories.CompanyHourBlockRepositoryInterface) CompanyHourBlockServiceInterface {
	return &CompanyHourBlockService{
		repo,
	}
}

func (chs *CompanyHourBlockService) CreateManyCompanyHourBlock(hourBlockCreate dtos.CreateCompanyHourBlockRequestBody, companyID uuid.UUID, tx *gorm.DB) *utils.AppError {
	var params = []dtos.CreateCompanyHourBlockParams{}
	for _, register := range hourBlockCreate.Registers {
		params = append(params, dtos.CreateCompanyHourBlockParams{
			CompanyID: companyID,
			Day:       register.Day,
			StartHour: register.StartHour,
			EndHour:   register.EndHour,
		})
	}

	err := chs.companyHourBlockRepository.CreateMany(params, tx)
	if err != nil {
		return err
	}

	return nil
}
