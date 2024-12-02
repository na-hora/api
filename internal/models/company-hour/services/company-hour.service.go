package services

import (
	"na-hora/api/internal/models/company-hour/dtos"
	"na-hora/api/internal/models/company-hour/repositories"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyHourServiceInterface interface {
	ListByCompanyID(companyID uuid.UUID) ([]dtos.ListHoursByCompanyIDResponse, *utils.AppError)
	RelateCompanyHour(hourCreate dtos.CreateCompanyHourRequestBody, companyID uuid.UUID, tx *gorm.DB) *utils.AppError
}

type CompanyHourService struct {
	companyHourRepository repositories.CompanyHourRepositoryInterface
}

func GetCompanyHourService(repo repositories.CompanyHourRepositoryInterface) CompanyHourServiceInterface {
	return &CompanyHourService{
		repo,
	}
}

func (chs *CompanyHourService) ListByCompanyID(companyID uuid.UUID) ([]dtos.ListHoursByCompanyIDResponse, *utils.AppError) {
	return chs.companyHourRepository.ListByCompanyID(companyID)
}

func (chs *CompanyHourService) RelateCompanyHour(hourCreate dtos.CreateCompanyHourRequestBody, companyID uuid.UUID, tx *gorm.DB) *utils.AppError {
	registeredHours, err := chs.ListByCompanyID(companyID)
	if err != nil {
		return err
	}

	existingHoursMap := make(map[uint]dtos.ListHoursByCompanyIDResponse)
	for _, hour := range registeredHours {
		existingHoursMap[hour.ID] = hour
	}

	var (
		toCreate []dtos.CreateCompanyHourParams
		toUpdate []dtos.CreateCompanyHourParams
		toDelete []int
	)

	receivedIDs := make(map[int]bool)
	for _, register := range hourCreate.Registers {
		if register.EndMinute < register.StartMinute {
			return &utils.AppError{
				Message:    "O horário final não pode ser menor que o horário inicial",
				StatusCode: http.StatusBadRequest,
			}
		}

		receivedIDs[register.ID] = true

		if existingHour, exists := existingHoursMap[uint(register.ID)]; exists {
			if existingHour.Weekday != register.Weekday || existingHour.StartMinute != register.StartMinute || existingHour.EndMinute != register.EndMinute {
				toUpdate = append(toUpdate, dtos.CreateCompanyHourParams{
					ID:          uint(register.ID),
					CompanyID:   companyID,
					Weekday:     register.Weekday,
					StartMinute: register.StartMinute,
					EndMinute:   register.EndMinute,
				})
			}
		} else {
			toCreate = append(toCreate, dtos.CreateCompanyHourParams{
				CompanyID:   companyID,
				Weekday:     register.Weekday,
				StartMinute: register.StartMinute,
				EndMinute:   register.EndMinute,
			})
		}
	}

	// registros a serem deletados
	for _, hour := range registeredHours {
		if !receivedIDs[int(hour.ID)] {
			toDelete = append(toDelete, int(hour.ID))
		}
	}

	if len(toCreate) > 0 {
		err = chs.companyHourRepository.CreateMany(toCreate, tx)
		if err != nil {
			return err
		}
	}

	if len(toUpdate) > 0 {
		err = chs.companyHourRepository.UpdateMany(toUpdate, tx)
		if err != nil {
			return err
		}
	}

	if len(toDelete) > 0 {
		err = chs.companyHourRepository.DeleteMany(toDelete, tx)
		if err != nil {
			return err
		}
	}

	return nil
}
