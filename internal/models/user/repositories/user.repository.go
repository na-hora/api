package repositories

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/user/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(dtos.CreateUserRequestBody) (*entity.User, *utils.AppError)
}

type userRepository struct {
	db *gorm.DB
}

func GetUserRepository() UserRepository {
	db := config.DB
	return &userRepository{db}
}

func (ur *userRepository) Create(insert dtos.CreateUserRequestBody) (*entity.User, *utils.AppError) {
	insertValue := entity.User{
		Username:  insert.Username,
		Password:  insert.Password,
		CompanyID: insert.CompanyID,
	}

	data := ur.db.Create(&insertValue)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &insertValue, nil
}
