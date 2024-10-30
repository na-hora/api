package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/client/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientRepositoryInterface interface {
	List(companyID uuid.UUID) ([]entity.Client, *utils.AppError)
	Create(companyID uuid.UUID, insert dtos.CreateClientRequestBody, tx *gorm.DB) (*entity.Client, *utils.AppError)
}

type ClientRepository struct {
	db *gorm.DB
}

func GetClientRepository(db *gorm.DB) ClientRepositoryInterface {
	return &ClientRepository{db}
}

func (cr *ClientRepository) List(companyID uuid.UUID) ([]entity.Client, *utils.AppError) {
	clients := []entity.Client{}

	data := cr.db.Where("company_id = ?", companyID).Find(&clients)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return clients, nil
}

func (cr *ClientRepository) Create(
	companyID uuid.UUID,
	insert dtos.CreateClientRequestBody,
	tx *gorm.DB,
) (*entity.Client, *utils.AppError) {
	if tx == nil {
		tx = cr.db
	}

	client := entity.Client{
		CompanyID: companyID,
		Name:      insert.Name,
		Phone:     insert.Phone,
		Email:     insert.Email,
	}

	data := tx.Create(&client)

	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &client, nil
}
