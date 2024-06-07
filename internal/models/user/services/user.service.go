package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/user/dtos"
	repositories "na-hora/api/internal/models/user/repositories"
	"na-hora/api/internal/utils"
)

type UserService interface {
	Create(userCreate dtos.CreateUserRequestBody) (*entity.User, *utils.AppError)
	GetByUsername(username string) (*entity.User, *utils.AppError)
}

type userService struct {
	userRepository repositories.UserRepository
}

func GetUserService() UserService {
	userRepository := repositories.GetUserRepository()
	return &userService{
		userRepository,
	}
}

func (us *userService) Create(userCreate dtos.CreateUserRequestBody) (*entity.User, *utils.AppError) {

	hash, passwordError := utils.HashPassword(userCreate.Password)
	if passwordError != nil {
		return nil, &utils.AppError{
			Message:    passwordError.Message,
			StatusCode: passwordError.StatusCode,
		}
	}

	userCreate.Password = hash

	userCreated, err := us.userRepository.Create(userCreate)
	if err != nil {
		return nil, err
	}
	return userCreated, nil
}

func (us *userService) GetByUsername(username string) (*entity.User, *utils.AppError) {
	user, err := us.userRepository.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
