package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/appointment/dtos"
	"na-hora/api/internal/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppointmentRepositoryInterface interface {
	List(companyID uuid.UUID, startDate time.Time, endDate time.Time) ([]entity.Appointment, *utils.AppError)
	Create(companyID uuid.UUID, insert dtos.CreateAppointmentParams, tx *gorm.DB) (*entity.Appointment, *utils.AppError)
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

func (ar *AppointmentRepository) Create(
	companyID uuid.UUID,
	insert dtos.CreateAppointmentParams,
	tx *gorm.DB,
) (*entity.Appointment, *utils.AppError) {
	if tx == nil {
		tx = ar.db
	}

	appointment := entity.Appointment{
		CompanyID:                companyID,
		StartTime:                insert.StartTime,
		ClientID:                 insert.ClientID,
		PetName:                  insert.PetName,
		Note:                     insert.Note,
		PaymentMode:              insert.PaymentMode,
		CompanyPetServiceValueID: insert.CompanyPetServiceValueID,
		TotalTime:                insert.TotalTime,
		TotalPrice:               insert.TotalPrice,
		Canceled:                 insert.Canceled,
	}

	data := tx.Create(&appointment)

	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &appointment, nil
}
