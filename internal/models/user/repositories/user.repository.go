package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/user/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(insert dtos.CreateUserRequestBody, tx *gorm.DB) (*entity.User, *utils.AppError)
	GetByID(ID uuid.UUID) (*entity.User, *utils.AppError)
	GetByUsername(username string) (*entity.User, *utils.AppError)
}

type UserRepository struct {
	db *gorm.DB
}

func GetUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (ur *UserRepository) Create(insert dtos.CreateUserRequestBody, tx *gorm.DB) (*entity.User, *utils.AppError) {
	if tx == nil {
		tx = ur.db
	}

	insertValue := entity.User{
		Username:  insert.Username,
		Password:  insert.Password,
		CompanyID: insert.CompanyID,
	}

	data := tx.Create(&insertValue)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &insertValue, nil
}

func (ur *UserRepository) GetByID(ID uuid.UUID) (*entity.User, *utils.AppError) {
	var user entity.User
	data := ur.db.Where("ID = ?", ID).First(&user)
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

func (ur *UserRepository) GetByUsername(username string) (*entity.User, *utils.AppError) {
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
