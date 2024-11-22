package services

import (
	"na-hora/api/internal/entity"
	petTypeRepos "na-hora/api/internal/models/company-pet-type/repositories"
	"na-hora/api/internal/models/pet-service/dtos"
	"na-hora/api/internal/models/pet-service/repositories"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PetServiceServiceInterface interface {
	CreatePetService(companyID uuid.UUID, petServiceCreate dtos.CreatePetServiceRequestBody, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError)
	UpdatePetService(companyID uuid.UUID, ID int, petServiceUpdate dtos.UpdatePetServiceRequestBody, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError)
	GetByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetService, *utils.AppError)
	GetByID(ID int, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError)
	DeleteByID(petServiceID int, tx *gorm.DB) *utils.AppError
}

type PetServiceService struct {
	petServiceRepository repositories.PetServiceRepositoryInterface
	petTypeRepository    petTypeRepos.CompanyPetTypeRepositoryInterface
}

func GetPetServiceService(
	petServiceRepo repositories.PetServiceRepositoryInterface,
	petTypeRepo petTypeRepos.CompanyPetTypeRepositoryInterface,
) PetServiceServiceInterface {
	return &PetServiceService{
		petServiceRepo,
		petTypeRepo,
	}
}

func (ps *PetServiceService) CreatePetService(
	companyID uuid.UUID,
	petServiceCreate dtos.CreatePetServiceRequestBody,
	tx *gorm.DB,
) (*entity.CompanyPetService, *utils.AppError) {
	petServiceCreated, err := ps.petServiceRepository.Create(companyID, petServiceCreate, tx)
	if err != nil {
		return nil, err
	}

	if petServiceCreate.PetTypes != nil {
		for _, petTypeID := range petServiceCreate.PetTypes {
			typeFound, findErr := ps.petTypeRepository.GetByID(petTypeID, tx)

			if findErr != nil {
				return nil, findErr
			}

			if typeFound == nil {
				return nil, &utils.AppError{
					Message:    "Pet type not found",
					StatusCode: http.StatusBadRequest,
				}
			}

			if typeFound.CompanyID != companyID {
				return nil, &utils.AppError{
					Message:    "Invalid pet type provided",
					StatusCode: http.StatusBadRequest,
				}
			}

			relateErr := ps.petServiceRepository.RelateToType(petServiceCreated.ID, typeFound.ID, tx)

			if relateErr != nil {
				return nil, relateErr
			}
		}
	}

	return petServiceCreated, nil
}

func (ps *PetServiceService) UpdatePetService(
	companyID uuid.UUID,
	ID int,
	petServiceUpdate dtos.UpdatePetServiceRequestBody,
	tx *gorm.DB,
) (*entity.CompanyPetService, *utils.AppError) {
	detailedPetService, err := ps.petServiceRepository.GetByID(ID, tx)
	if err != nil {
		return nil, err
	}

	if detailedPetService == nil {
		return nil, &utils.AppError{
			Message:    "pet service not found",
			StatusCode: http.StatusNotFound,
		}
	}

	petServiceUpdated, err := ps.petServiceRepository.Update(companyID, dtos.UpdateCompanyPetServiceParams{
		ID:          ID,
		Name:        petServiceUpdate.Name,
		Paralellism: petServiceUpdate.Paralellism,
	}, tx)

	if err != nil {
		return nil, err
	}

	if petServiceUpdate.PetTypes != nil {
		var existingPetTypeIDs []int

		if petServiceUpdate.PetTypes != nil {
			for _, relation := range detailedPetService.ServiceTypes {
				existingPetTypeIDs = append(existingPetTypeIDs, relation.CompanyPetTypeID)
			}
		}

		existingMap := make(map[int]bool)
		for _, id := range existingPetTypeIDs {
			existingMap[id] = true
		}

		newMap := make(map[int]bool)
		for _, id := range petServiceUpdate.PetTypes {
			newMap[id] = true
		}

		// Add new relationships
		for _, petTypeID := range petServiceUpdate.PetTypes {
			if !existingMap[petTypeID] {
				typeFound, findErr := ps.petTypeRepository.GetByID(petTypeID, tx)
				if findErr != nil {
					return nil, findErr
				}

				if typeFound == nil {
					return nil, &utils.AppError{
						Message:    "pet type not found",
						StatusCode: http.StatusBadRequest,
					}
				}

				if typeFound.CompanyID != companyID {
					return nil, &utils.AppError{
						Message:    "invalid pet type provided",
						StatusCode: http.StatusBadRequest,
					}
				}

				relateErr := ps.petServiceRepository.RelateToType(ID, petTypeID, tx)
				if relateErr != nil {
					return nil, relateErr
				}
			}
		}

		// Remove obsolete relationships
		for _, petTypeID := range existingPetTypeIDs {
			if !newMap[petTypeID] {
				unrelateErr := ps.petServiceRepository.UnrelateFromType(ID, petTypeID, tx)
				if unrelateErr != nil {
					return nil, unrelateErr
				}
			}
		}
	}

	return petServiceUpdated, nil
}

func (ps *PetServiceService) GetByCompanyID(companyID uuid.UUID, tx *gorm.DB) ([]entity.CompanyPetService, *utils.AppError) {
	petServices, err := ps.petServiceRepository.GetByCompanyID(companyID, tx)
	if err != nil {
		return nil, err
	}

	return petServices, nil
}

func (ps *PetServiceService) GetByID(ID int, tx *gorm.DB) (*entity.CompanyPetService, *utils.AppError) {
	petService, err := ps.petServiceRepository.GetByID(ID, tx)
	if err != nil {
		return nil, err
	}

	return petService, nil
}

func (ps *PetServiceService) DeleteByID(petServiceID int, tx *gorm.DB) *utils.AppError {
	err := ps.petServiceRepository.DeleteByID(petServiceID, tx)
	if err != nil {
		return err
	}

	return nil
}
