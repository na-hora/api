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
	GetByID(ID uuid.UUID) (*entity.Client, *utils.AppError)
	List(companyID uuid.UUID) ([]entity.Client, *utils.AppError)
	Create(companyID uuid.UUID, insert dtos.CreateClientRequestBody, tx *gorm.DB) (*entity.Client, *utils.AppError)
	GetByEmail(companyID uuid.UUID, email string) (*entity.Client, *utils.AppError)
	Update(companyID uuid.UUID, update dtos.UpdateClientRequestBody, tx *gorm.DB) (*entity.Client, *utils.AppError)
}

type ClientRepository struct {
	db *gorm.DB
}

func GetClientRepository(db *gorm.DB) ClientRepositoryInterface {
	return &ClientRepository{db}
}

func (cr *ClientRepository) GetByID(ID uuid.UUID) (*entity.Client, *utils.AppError) {
	var client entity.Client
	data := cr.db.First(&client, ID)

	if data.Error != nil {
		if data.Error != gorm.ErrRecordNotFound {
			return nil, &utils.AppError{
				Message:    data.Error.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		return nil, nil
	}

	return &client, nil
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

func (cr *ClientRepository) Update(
	companyID uuid.UUID,
	update dtos.UpdateClientRequestBody,
	tx *gorm.DB,
) (*entity.Client, *utils.AppError) {
	if tx == nil {
		tx = cr.db
	}

	client := entity.Client{
		ID:        update.ID,
		CompanyID: companyID,
		Name:      update.Name,
		Phone:     update.Phone,
		Email:     update.Email,
	}

	data := tx.Updates(&client)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &client, nil
}

func (cr *ClientRepository) GetByEmail(companyID uuid.UUID, email string) (*entity.Client, *utils.AppError) {
	var client entity.Client

	data := cr.db.Where("company_id = ? AND email = ?", companyID, email).First(&client)
	if data.Error != nil {
		if data.Error != gorm.ErrRecordNotFound {
			return nil, &utils.AppError{
				Message:    data.Error.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		return nil, nil
	}

	return &client, nil
}
