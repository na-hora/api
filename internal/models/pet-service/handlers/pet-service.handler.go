package handlers

import (
	"encoding/json"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	petServiceDTOs "na-hora/api/internal/models/pet-service/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	petServiceServices "na-hora/api/internal/models/pet-service/services"
	tokenServices "na-hora/api/internal/models/token/services"

	"github.com/go-playground/validator/v10"
)

type PetServiceInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type petServiceHandler struct {
	petServiceService petServiceServices.PetServiceInterface
	tokenService      tokenServices.TokenServiceInterface
}

func GetPetServiceHandler() PetServiceInterface {
	petServiceService := injector.InitializePetServiceService(config.DB)
	tokenService := injector.InitializeTokenService(config.DB)

	return &petServiceHandler{
		petServiceService,
		tokenService,
	}
}

func (ph *petServiceHandler) Register(w http.ResponseWriter, r *http.Request) {
	var petServicePayload petServiceDTOs.CreatePetServiceRequestBody

	err := json.NewDecoder(r.Body).Decode(&petServicePayload)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(petServicePayload)
	if err != nil {
		utils.ResponseValidationErrors(err, w)
		return
	}

	petServiceCreated, appErr := ph.petServiceService.CreatePetService(petServicePayload, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, petServiceCreated)
}
