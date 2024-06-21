package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/city/repositories"
	"na-hora/api/internal/utils"
)

type CityServiceInterface interface {
	ListAllByState(stateID uint) ([]entity.City, *utils.AppError)
	GetByIBGE(ibge string) (*entity.City, *utils.AppError)
}

type CityService struct {
	cityRepository repositories.CityRepositoryInterface
}

func GetCityService(repo repositories.CityRepositoryInterface) CityServiceInterface {
	return &CityService{
		repo,
	}
}

func (cs *CityService) ListAllByState(stateID uint) ([]entity.City, *utils.AppError) {
	allCities, err := cs.cityRepository.ListAllByState(stateID)
	if err != nil {
		return nil, err
	}
	return allCities, nil
}

func (cs *CityService) GetByIBGE(ibge string) (*entity.City, *utils.AppError) {
	cityFound, err := cs.cityRepository.GetByIBGE(ibge)
	if err != nil {
		return nil, err
	}
	return cityFound, nil
}
