package handlers

import (
	"encoding/json"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	petServiceDTOs "na-hora/api/internal/models/pet-service/dtos"
	"na-hora/api/internal/utils"
	"net/http"
	"strconv"

	petServiceServices "na-hora/api/internal/models/pet-service/services"
	tokenServices "na-hora/api/internal/models/token/services"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PetServiceInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	ListAll(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
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

	tx := config.StartTransaction()
	petServiceCreated, appErr := ph.petServiceService.CreatePetService(petServicePayload, tx)
	if appErr != nil {
		tx.Rollback()
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	dbInfo := tx.Commit()
	if dbInfo.Error != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, dbInfo.Error.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, petServiceCreated)
}

func (ph *petServiceHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	companyId := chi.URLParam(r, "companyId")

	companyIdParsedToUUID, err := uuid.Parse(companyId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	petServices, appErr := ph.petServiceService.GetByCompanyID(companyIdParsedToUUID, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, petServices)
}

func (ph *petServiceHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	serviceId := r.URL.Query().Get("serviceId")

	serviceIdParsedToInt, err := strconv.Atoi(serviceId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := ph.petServiceService.DeleteByPetServiceID(serviceIdParsedToInt, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
