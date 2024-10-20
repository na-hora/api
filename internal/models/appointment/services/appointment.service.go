package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/appointment/repositories"
	"na-hora/api/internal/utils"
	"time"

	"github.com/google/uuid"
)

type AppointmentServiceInterface interface {
	List(companyID uuid.UUID, startDate time.Time, endDate time.Time) ([]entity.Appointment, *utils.AppError)
}

type AppointmentService struct {
	appointmentRepository repositories.AppointmentRepositoryInterface
}

func GetAppointmentService(repo repositories.AppointmentRepositoryInterface) AppointmentServiceInterface {
	return &AppointmentService{
		repo,
	}
}

func (cs *AppointmentService) List(companyID uuid.UUID, startDate time.Time, endDate time.Time) ([]entity.Appointment, *utils.AppError) {
	allAppointments, err := cs.appointmentRepository.List(companyID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return allAppointments, nil
}
