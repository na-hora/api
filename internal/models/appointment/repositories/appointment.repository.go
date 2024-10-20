package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppointmentRepositoryInterface interface {
	List(companyID uuid.UUID, startDate time.Time, endDate time.Time) ([]entity.Appointment, *utils.AppError)
}

type AppointmentRepository struct {
	db *gorm.DB
}

func GetAppointmentRepository(db *gorm.DB) AppointmentRepositoryInterface {
	return &AppointmentRepository{db}
}

func (cr *AppointmentRepository) List(companyID uuid.UUID, startDate time.Time, endDate time.Time) ([]entity.Appointment, *utils.AppError) {
	allAppointments := []entity.Appointment{}

	data := cr.db.Where("company_id = ? AND start_time BETWEEN ? AND ?", companyID, startDate, endDate).Find(&allAppointments)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return allAppointments, nil
}
