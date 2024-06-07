package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/user/dtos"
	repositories "na-hora/api/internal/models/user/repositories"
	"na-hora/api/internal/utils"
)

type UserServiceInterface interface {
	Create(userCreate dtos.CreateUserRequestBody) (*entity.User, *utils.AppError)
	GetByUsername(username string) (*entity.User, *utils.AppError)
}

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func GetUserService(repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repo,
	}
}

func (us *UserService) Create(userCreate dtos.CreateUserRequestBody) (*entity.User, *utils.AppError) {

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

func (us *UserService) GetByUsername(username string) (*entity.User, *utils.AppError) {
	user, err := us.userRepository.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
