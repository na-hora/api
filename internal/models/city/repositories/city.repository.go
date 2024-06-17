package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/utils"
	"net/http"

	"gorm.io/gorm"
)

type CityRepositoryInterface interface {
	ListAllByState(stateID uint) ([]entity.City, *utils.AppError)
}

type CityRepository struct {
	db *gorm.DB
}

func GetCityRepository(db *gorm.DB) CityRepositoryInterface {
	return &CityRepository{db}
}

func (t *CityRepository) ListAllByState(stateID uint) ([]entity.City, *utils.AppError) {
	allCities := []entity.City{}

	data := t.db.Where("state_id = ?", stateID).Find(&allCities)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return allCities, nil
}
