package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/appointment/dtos"
	appointmentRepository "na-hora/api/internal/models/appointment/repositories"
	petServiceRepository "na-hora/api/internal/models/pet-service/repositories"
	"na-hora/api/internal/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppointmentServiceInterface interface {
	List(companyID uuid.UUID, startDate time.Time, endDate time.Time) ([]entity.Appointment, *utils.AppError)
	Create(companyID uuid.UUID, insert dtos.CreateAppointmentsRequestBody, tx *gorm.DB) (*entity.Appointment, *utils.AppError)
}

type AppointmentService struct {
	appointmentRepo appointmentRepository.AppointmentRepositoryInterface
	petServiceRepo  petServiceRepository.PetServiceRepositoryInterface
}

func GetAppointmentService(
	appointmentRepo appointmentRepository.AppointmentRepositoryInterface,
	petServiceRepo petServiceRepository.PetServiceRepositoryInterface,
) AppointmentServiceInterface {
	return &AppointmentService{
		appointmentRepo,
		petServiceRepo,
	}
}

func (as *AppointmentService) List(companyID uuid.UUID, startDate time.Time, endDate time.Time) ([]entity.Appointment, *utils.AppError) {
	allAppointments, err := as.appointmentRepo.List(companyID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return allAppointments, nil
}

func (as *AppointmentService) Create(companyID uuid.UUID, insert dtos.CreateAppointmentsRequestBody, tx *gorm.DB) (*entity.Appointment, *utils.AppError) {
	companyPetServiceValue, appErr := as.petServiceRepo.GetConfigurationBySizeAndHair(
		insert.CompanyPetServiceID,
		insert.CompanyPetSizeID,
		insert.CompanyPetHairID,
		tx,
	)

	if appErr != nil {
		return nil, appErr
	}

	if companyPetServiceValue == nil {
		return nil, &utils.AppError{
			Message:    "pet service value not found",
			StatusCode: http.StatusNotFound,
		}
	}

	createParams := dtos.CreateAppointmentParams{
		ClientID:                 insert.ClientID,
		CompanyPetServiceValueID: companyPetServiceValue.ID,
		StartTime:                insert.StartTime,
		PetName:                  insert.PetName,
		PaymentMode:              insert.PaymentMode,
		Note:                     insert.Note,
		TotalTime:                companyPetServiceValue.ExecutionTime,
		TotalPrice:               companyPetServiceValue.Price,
		Canceled:                 false,
	}

	appointmentCreated, err := as.appointmentRepo.Create(
		companyID,
		createParams,
		tx,
	)

	if err != nil {
		return nil, err
	}

	return appointmentCreated, nil
}
