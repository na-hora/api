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
	GetByUsername(username string) (*entity.User, *utils.AppError)
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

func (ur *userRepository) GetByUsername(username string) (*entity.User, *utils.AppError) {
	var user entity.User
	data := ur.db.Where("username = ?", username).First(&user)
	if data.Error != nil {
		if data.Error != gorm.ErrRecordNotFound {
			return nil, &utils.AppError{
				Message:    data.Error.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		return nil, nil
	}
	return &user, nil
}
