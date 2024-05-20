package handlers

import (
	"encoding/json"
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
	userService services.UserService
}

func GetUserHandler() UserHandler {
	userService := services.GetUserService()
	return &userHandler{userService}
}

func (u *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userPayload userDTOs.CreateUserRequestBody

	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, "Invalid body")
		return
	}

	validate := validator.New()
	err = validate.Struct(userPayload)
	if err != nil {
		utils.ResponseValidationErrors(err, w)
		return
	}

	hash, passwordError := utils.HashPassword(userPayload.Password)
	if passwordError != nil {
		utils.ResponseJSON(w, passwordError.StatusCode, passwordError.Message)
		return
	}

	userPayload.Password = hash

	user, serviceErr := u.userService.Create(userPayload)
	if serviceErr != nil {
		utils.ResponseJSON(w, serviceErr.StatusCode, serviceErr.Message)
		return
	}

	response := &userDTOs.CreateUserResponse{
		ID: user.ID,
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}
