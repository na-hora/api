package handlers

import (
	"encoding/json"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/company-pet-type/dtos"
	companyPetTypeServices "na-hora/api/internal/models/company-pet-type/services"
	"na-hora/api/internal/utils"
	"net/http"
)

type CompanyPetTypeInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type CompanyPetTypeHandler struct {
	companyPetTypeService companyPetTypeServices.CompanyPetTypeServiceInterface
}

func GetCompanyPetTypeHandler() CompanyPetTypeInterface {
	companyPetTypeService := injector.InitializeCompanyPetTypeService(config.DB)

	return &CompanyPetTypeHandler{
		companyPetTypeService,
	}
}

func (cpt *CompanyPetTypeHandler) Register(w http.ResponseWriter, r *http.Request) {
	var petTypePayload dtos.CreatePetTypeRequestBody

	err := json.NewDecoder(r.Body).Decode(&petTypePayload)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// validatorFound, tokenErr := cpt.tokenService.GetValidToken(companyPayload.Validator)
	// if tokenErr != nil {
	// 	utils.ResponseJSON(w, tokenErr.StatusCode, tokenErr.Message)
	// 	return
	// }

	// if validatorFound == nil {
	// 	utils.ResponseJSON(w, http.StatusUnauthorized, "validator not found")
	// 	return
	// }

	// appErr := cpt.companyPetTypeService.CreatePetType(companyPetServicePayload.CompanyID, companyPetServicePayload.Name, nil)
	// if appErr != nil {
	// 	utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
	// }

	utils.ResponseJSON(w, http.StatusCreated, nil)
}
