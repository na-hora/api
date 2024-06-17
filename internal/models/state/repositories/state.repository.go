package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/utils"
	"net/http"

	"gorm.io/gorm"
)

type StateRepositoryInterface interface {
	ListAll() ([]entity.State, *utils.AppError)
}

type StateRepository struct {
	db *gorm.DB
}

func GetStateRepository(db *gorm.DB) StateRepositoryInterface {
	return &StateRepository{db}
}

func (t *StateRepository) ListAll() ([]entity.State, *utils.AppError) {
	allStates := []entity.State{}

	data := t.db.Find(&allStates)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return allStates, nil
}
