package handlers

import (
	"encoding/json"
	"fmt"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"strings"

	userDTOs "na-hora/api/internal/models/user/dtos"
	userServices "na-hora/api/internal/models/user/services"

	tokenDTOs "na-hora/api/internal/models/token/dtos"
	tokenServices "na-hora/api/internal/models/token/services"

	"na-hora/api/internal/providers"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserHandlerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	ForgotPassword(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
	// UpdatePassword(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService  userServices.UserServiceInterface
	tokenService tokenServices.TokenServiceInterface
}

func GetUserHandler() UserHandlerInterface {
	userService := injector.InitializeUserService(config.DB)
	tokenService := injector.InitializeTokenService(config.DB)

	return &userHandler{
		userService,
		tokenService,
	}
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

	userChecked, passwordErr := u.userService.CheckPassword(userPayload)
	if passwordErr != nil {
		utils.ResponseJSON(w, passwordErr.StatusCode, passwordErr.Message)
		return
	}

	if userChecked == nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, tokenErr := authentication.GenerateToken(userChecked.ID, userChecked.Username)
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

func (u *userHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var userPayload userDTOs.ForgotUserPasswordRequestBody

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

	user, serviceErr := u.userService.GetByUsername(userPayload.Email)
	if serviceErr != nil {
		utils.ResponseJSON(w, serviceErr.StatusCode, serviceErr.Message)
		return
	}

	if user == nil {
		utils.ResponseJSON(w, http.StatusNotFound, "user not found")
		return
	}

	resetPassToken, tokenErr := u.tokenService.Generate(
		tokenDTOs.GenerateTokenRequestBody{
			Note: fmt.Sprintf("forgot-password:%s", user.ID),
		},
	)

	if tokenErr != nil {
		utils.ResponseJSON(w, tokenErr.StatusCode, tokenErr.Message)
		return
	}

	emailProvider := providers.NewEmailProvider()
	emailProvider.SendForgotPasswordEmail(
		userPayload.Email,
		resetPassToken.Key,
	)
}

func (u *userHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var userPayload userDTOs.ResetUserPasswordRequestBody

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

	validatorFound, tokenErr := u.tokenService.GetValidToken(userPayload.Validator)
	if tokenErr != nil {
		utils.ResponseJSON(w, tokenErr.StatusCode, tokenErr.Message)
		return
	}

	if validatorFound == nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, "validator not found")
		return
	}

	user, serviceErr := u.userService.GetByUsername(userPayload.Email)

	if serviceErr != nil {
		utils.ResponseJSON(w, serviceErr.StatusCode, serviceErr.Message)
		return
	}

	if user == nil {
		utils.ResponseJSON(w, http.StatusNotFound, "user not found")
		return
	}

	if strings.Split(validatorFound.Note, ":")[1] != user.ID.String() {
		utils.ResponseJSON(w, http.StatusUnauthorized, "invalid validator for this user")
		return
	}

	updatePassErr := u.userService.UpdatePassword(user.ID, userPayload.Password, nil)
	if updatePassErr != nil {
		utils.ResponseJSON(w, updatePassErr.StatusCode, updatePassErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "password updated successfully")
}
