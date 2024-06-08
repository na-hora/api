package handlers

import (
	"encoding/json"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	userDTOs "na-hora/api/internal/models/user/dtos"
	"na-hora/api/internal/models/user/services"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService services.UserServiceInterface
}

func GetUserHandler() *userHandler {
	userService := injector.InitializeUserService(config.DB)

	return &userHandler{userService}
}

func (u *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userPayload userDTOs.CreateUserRequestBody

	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(userPayload)
	if err != nil {
		utils.ResponseValidationErrors(err, w)
		return
	}

	userAlreadyExists, usernameErr := u.userService.GetByUsername(userPayload.Username)
	if usernameErr != nil {
		utils.ResponseJSON(w, usernameErr.StatusCode, usernameErr.Message)
		return
	}

	if userAlreadyExists != nil {
		utils.ResponseJSON(w, http.StatusConflict, "User already exists")
		return
	}

	user, serviceErr := u.userService.Create(userPayload, nil)
	if serviceErr != nil {
		utils.ResponseJSON(w, serviceErr.StatusCode, serviceErr.Message)
		return
	}

	response := &userDTOs.CreateUserResponse{
		ID: user.ID,
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}
