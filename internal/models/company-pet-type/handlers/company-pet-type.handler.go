package handlers

import (
	"encoding/json"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/company-pet-type/dtos"
	companyPetTypeServices "na-hora/api/internal/models/company-pet-type/services"
	tokenServices "na-hora/api/internal/models/token/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type CompanyPetTypeInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	GetByCompanyID(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
}

type CompanyPetTypeHandler struct {
	companyPetTypeService companyPetTypeServices.CompanyPetTypeServiceInterface
	tokenService          tokenServices.TokenServiceInterface
}

func GetCompanyPetTypeHandler() CompanyPetTypeInterface {
	companyPetTypeService := injector.InitializeCompanyPetTypeService(config.DB)
	tokenService := injector.InitializeTokenService(config.DB)

	return &CompanyPetTypeHandler{
		companyPetTypeService,
		tokenService,
	}
}

func (cpt *CompanyPetTypeHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userLogged, userErr := authentication.JwtUserOrThrow(ctx)

	petTypePayload := ctx.Value(utils.ValidatedBodyKey).(*dtos.CreatePetTypeRequestBody)
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	appErr := cpt.companyPetTypeService.Create(userLogged.CompanyID, petTypePayload.Name, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
	}

	utils.ResponseJSON(w, http.StatusCreated, nil)
}

func (cpt *CompanyPetTypeHandler) GetByCompanyID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userLogged, userErr := authentication.JwtUserOrThrow(ctx)
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	petTypes, appErr := cpt.companyPetTypeService.GetByCompanyID(userLogged.CompanyID)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	var responsePetTypes = make([]dtos.ListPetTypesByCompanyResponse, 0)
	for _, petType := range petTypes {
		responsePetTypes = append(responsePetTypes, dtos.ListPetTypesByCompanyResponse{
			ID:   petType.ID,
			Name: petType.Name,
		})
	}

	utils.ResponseJSON(w, http.StatusOK, responsePetTypes)
}

func (cpt *CompanyPetTypeHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	petTypeId := chi.URLParam(r, "ID")

	parsedPetTypeId, err := strconv.Atoi(petTypeId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := cpt.companyPetTypeService.DeleteByID(parsedPetTypeId, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, nil)
}

func (cpt *CompanyPetTypeHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	petTypeId := chi.URLParam(r, "ID")

	parsedPetTypeId, err := strconv.Atoi(petTypeId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var body dtos.CreateCompanyPetTypeParams

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := cpt.companyPetTypeService.UpdateByID(parsedPetTypeId, body, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, nil)
}
