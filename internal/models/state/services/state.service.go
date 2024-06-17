package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/state/repositories"
	"na-hora/api/internal/utils"
)

type StateServiceInterface interface {
	ListAll() ([]entity.State, *utils.AppError)
}

type StateService struct {
	stateRepository repositories.StateRepositoryInterface
}

func GetStateService(repo repositories.StateRepositoryInterface) StateServiceInterface {
	return &StateService{
		repo,
	}
}

func (ss *StateService) ListAll() ([]entity.State, *utils.AppError) {
	allStates, err := ss.stateRepository.ListAll()
	if err != nil {
		return nil, err
	}
	return allStates, nil
}
