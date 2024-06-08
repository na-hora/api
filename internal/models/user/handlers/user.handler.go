package handlers

import (
	"encoding/json"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	userDTOs "na-hora/api/internal/models/user/dtos"
	"na-hora/api/internal/models/user/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
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

func (u *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userPayload userDTOs.LoginUserRequestBody

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

	user, serviceErr := u.userService.GetByUsername(userPayload.Username)
	if serviceErr != nil {
		utils.ResponseJSON(w, serviceErr.StatusCode, serviceErr.Message)
		return
	}

	if user == nil {
		utils.ResponseJSON(w, http.StatusNotFound, "User not found")
		return
	}

	passwordCheck, passwordErr := u.userService.CheckPassword(userPayload)
	if passwordErr != nil {
		utils.ResponseJSON(w, passwordErr.StatusCode, passwordErr.Message)
		return
	}

	if passwordCheck == nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, tokenErr := authentication.GenerateToken(userPayload.Username)
	if tokenErr != nil {
		utils.ResponseJSON(w, tokenErr.StatusCode, tokenErr.Message)
		return
	}

	response := &userDTOs.LoginUserResponse{
		ID:    user.ID,
		Token: token,
	}

	utils.ResponseJSON(w, http.StatusOK, response)
}
