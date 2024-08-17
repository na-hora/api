package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/pet-service/dtos"
	"na-hora/api/internal/models/pet-service/repositories"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PetServiceServiceInterface interface {
	CreatePetService(petServiceCreate dtos.CreatePetServiceRequestBody, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError)
	GetByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetService, *utils.AppError)
	DeleteByID(petServiceID int, tx *gorm.DB) *utils.AppError
}

type PetServiceService struct {
	petServiceRepository repositories.PetServiceRepositoryInterface
}

func GetPetServiceService(repo repositories.PetServiceRepositoryInterface, tx *gorm.DB) PetServiceServiceInterface {
	return &PetServiceService{
		repo,
	}
}

func (ps *PetServiceService) CreatePetService(
	petServiceCreate dtos.CreatePetServiceRequestBody,
	tx *gorm.DB,
) (*entity.CompanyPetService, *utils.AppError) {
	petServiceCreated, err := ps.petServiceRepository.Create(petServiceCreate, tx)
	if err != nil {
		return nil, err
	}

	if petServiceCreate.Configurations != nil {
		for _, configurationParams := range petServiceCreate.Configurations {
			configuration := dtos.CreateCompanyPetServiceConfigurationParams{
				Price:               configurationParams.Price,
				ExecutionTime:       configurationParams.ExecutionTime,
				CompanyPetServiceID: petServiceCreated.ID,
				CompanyPetSizeID:    configurationParams.CompanyPetSizeID,
				CompanyPetHairID:    configurationParams.CompanyPetHairID,
			}

			_, err := ps.petServiceRepository.CreateConfiguration(
				petServiceCreated.ID,
				configuration,
				tx,
			)

			if err != nil {
				return nil, err
			}
		}
	}

	return petServiceCreated, nil
}

func (ps *PetServiceService) GetByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetService, *utils.AppError) {
	petServices, err := ps.petServiceRepository.GetByCompanyID(companyID, tx)
	if err != nil {
		return nil, err
	}

	if len(petServices) == 0 {
		return nil, &utils.AppError{
			Message:    "pet services not found",
			StatusCode: http.StatusNotFound,
		}
	}

	return petServices, nil
}

func (ps *PetServiceService) DeleteByID(petServiceID int, tx *gorm.DB) *utils.AppError {
	err := ps.petServiceRepository.DeleteByID(petServiceID, tx)
	if err != nil {
		return err
	}

	return nil
}
