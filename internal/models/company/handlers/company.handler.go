package handlers

import (
	"encoding/json"
	companyDTOs "na-hora/api/internal/models/company/dtos"
	"na-hora/api/internal/models/company/services"
	"na-hora/api/internal/utils"
	"net/http"
)

type CompanyHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type companyHandler struct {
	companyService services.CompanyService
}

func GetCompanyHandler() CompanyHandler {
	companyService := services.GetCompanyService()
	return &companyHandler{
		companyService,
	}
}

func (c *companyHandler) Register(w http.ResponseWriter, r *http.Request) {
	var companyPayload companyDTOs.CreateCompanyRequestBody

	err := json.NewDecoder(r.Body).Decode(&companyPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	appErr := utils.ValidateCNPJ(companyPayload.CNPJ)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	appErr = c.companyService.CreateCompany(companyPayload)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}
}
