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

	"github.com/go-playground/validator/v10"
)

type CompanyPetTypeInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
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
	var petTypePayload dtos.CreatePetTypeRequestBody

	err := json.NewDecoder(r.Body).Decode(&petTypePayload)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(petTypePayload)
	if err != nil {
		utils.ResponseValidationErrors(err, w, "body")
		return
	}

	ctx := r.Context()
	userLogged, userErr := authentication.JwtUserOrThrow(ctx)
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	appErr := cpt.companyPetTypeService.CreatePetType(userLogged.CompanyID, petTypePayload.Name, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
	}

	utils.ResponseJSON(w, http.StatusCreated, nil)
}
