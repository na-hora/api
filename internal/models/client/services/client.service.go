package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/client/dtos"
	"na-hora/api/internal/models/client/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientServiceInterface interface {
	List(companyID uuid.UUID) ([]entity.Client, *utils.AppError)
	Create(companyID uuid.UUID, insert dtos.CreateClientRequestBody, tx *gorm.DB) (*entity.Client, *utils.AppError)
}

type ClientService struct {
	clientRepo repositories.ClientRepositoryInterface
}

func GetClientService(
	clientRepo repositories.ClientRepositoryInterface,
) ClientServiceInterface {
	return &ClientService{
		clientRepo,
	}
}

func (cs *ClientService) List(companyID uuid.UUID) ([]entity.Client, *utils.AppError) {
	clients, err := cs.clientRepo.List(companyID)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (cs *ClientService) Create(companyID uuid.UUID, insert dtos.CreateClientRequestBody, tx *gorm.DB) (*entity.Client, *utils.AppError) {
	clientCreated, err := cs.clientRepo.Create(
		companyID,
		insert,
		tx,
	)

	if err != nil {
		return nil, err
	}

	return clientCreated, nil
}
