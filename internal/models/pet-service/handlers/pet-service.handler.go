package handlers

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	petServiceDTOs "na-hora/api/internal/models/pet-service/dtos"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"na-hora/api/internal/utils/conversor"
	"net/http"
	"strconv"

	petServiceServices "na-hora/api/internal/models/pet-service/services"
	tokenServices "na-hora/api/internal/models/token/services"

	"github.com/go-chi/chi"
)

type PetServiceHandlerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	RelateValues(w http.ResponseWriter, r *http.Request)
	ListAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
}

type petServiceHandler struct {
	petServiceService petServiceServices.PetServiceServiceInterface
	tokenService      tokenServices.TokenServiceInterface
}

func GetPetServiceHandler() PetServiceHandlerInterface {
	petServiceService := injector.InitializePetServiceService(config.DB)
	tokenService := injector.InitializeTokenService(config.DB)

	return &petServiceHandler{
		petServiceService,
		tokenService,
	}
}

func (ph *petServiceHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	petServicePayload := ctx.Value(utils.ValidatedBodyKey).(*petServiceDTOs.CreatePetServiceRequestBody)

	userLogged, userErr := authentication.JwtUserOrThrow(r.Context())
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	tx := config.StartTransaction()
	petServiceCreated, appErr := ph.petServiceService.CreatePetService(
		userLogged.CompanyID,
		*petServicePayload,
		tx,
	)

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

	response := petServiceDTOs.CreatePetServiceResponse{
		ID: petServiceCreated.ID,
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}

func (ph *petServiceHandler) RelateValues(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")

	strConv := conversor.GetStringConversor()
	IDParsedToInt, err := strConv.ToInt(ID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	payload := ctx.Value(utils.ValidatedBodyKey).(*petServiceDTOs.PetServiceValuesRelateRequestBody)

	userLogged, userErr := authentication.JwtUserOrThrow(r.Context())
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	tx := config.StartTransaction()
	appErr := ph.petServiceService.RelateValues(userLogged.CompanyID, IDParsedToInt, *payload, tx)

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

	utils.ResponseJSON(w, http.StatusOK, nil)
}

func (ph *petServiceHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	companyID := r.URL.Query().Get("companyId")

	if companyID == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, "companyId is required")
		return
	}

	strConv := conversor.GetStringConversor()
	companyIdConverted, err := strConv.ToUUID(companyID)

	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	petServices, appErr := ph.petServiceService.GetByCompanyID(companyIdConverted, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	var responsePetService = make([]petServiceDTOs.ListPetServicesByCompanyResponse, 0)
	for _, petService := range petServices {
		petTypes := make([]petServiceDTOs.PetTypeResponse, 0)
		for _, serviceType := range petService.ServiceTypes {
			if serviceType.CompanyPetType.ID != 0 {
				petTypes = append(petTypes, petServiceDTOs.PetTypeResponse{
					ID:   serviceType.CompanyPetType.ID,
					Name: serviceType.CompanyPetType.Name,
				})
			}
		}

		responsePetService = append(responsePetService, petServiceDTOs.ListPetServicesByCompanyResponse{
			ID:          petService.ID,
			Name:        petService.Name,
			Paralellism: petService.Paralellism,
			PetTypes:    petTypes,
		})
	}

	utils.ResponseJSON(w, http.StatusOK, responsePetService)
}

func (ph *petServiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")

	strConv := conversor.GetStringConversor()
	IDParsedToInt, err := strConv.ToInt(ID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	petService, appErr := ph.petServiceService.GetByID(IDParsedToInt, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	if petService == nil {
		utils.ResponseJSON(w, http.StatusNotFound, "pet service not found")
		return
	}

	var responsePetService petServiceDTOs.GetPetServiceByIDResponse

	responsePetService.ID = petService.ID
	responsePetService.Name = petService.Name
	responsePetService.Paralellism = petService.Paralellism
	responsePetService.Configurations = make([]petServiceDTOs.PetServiceConfiguration, 0)
	responsePetService.PetTypes = make([]int, 0)

	for _, configuration := range petService.Configurations {
		responsePetService.Configurations = append(responsePetService.Configurations, petServiceDTOs.PetServiceConfiguration{
			ID:               configuration.ID,
			Price:            configuration.Price,
			ExecutionTime:    configuration.ExecutionTime,
			CompanyPetHairID: configuration.CompanyPetHairID,
			CompanyPetSizeID: configuration.CompanyPetSizeID,
		})
	}

	for _, serviceType := range petService.ServiceTypes {
		responsePetService.PetTypes = append(responsePetService.PetTypes, serviceType.CompanyPetTypeID)
	}

	utils.ResponseJSON(w, http.StatusOK, responsePetService)
}

func (ph *petServiceHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	petServiceId := chi.URLParam(r, "ID")

	petServiceIdParsedToInt, err := strconv.Atoi(petServiceId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := ph.petServiceService.DeleteByID(petServiceIdParsedToInt, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ph *petServiceHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	petServiceId := chi.URLParam(r, "ID")

	strConv := conversor.GetStringConversor()
	petServiceIdParsedToInt, err := strConv.ToInt(petServiceId)

	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	petServicePayload := ctx.Value(utils.ValidatedBodyKey).(*petServiceDTOs.UpdatePetServiceRequestBody)

	userLogged, userErr := authentication.JwtUserOrThrow(r.Context())
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	tx := config.StartTransaction()
	petServiceUpdated, appErr := ph.petServiceService.UpdatePetService(
		userLogged.CompanyID,
		petServiceIdParsedToInt,
		*petServicePayload,
		tx,
	)

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

	response := petServiceDTOs.UpdatePetServiceResponse{
		ID: petServiceUpdated.ID,
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}
