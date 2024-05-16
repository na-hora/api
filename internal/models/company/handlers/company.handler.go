package handlers

import (
	"encoding/json"
	"na-hora/api/internal/dto"
	"na-hora/api/internal/models/company/services"

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
	var companyPayload dto.CompanyCreate

	err := json.NewDecoder(r.Body).Decode(&companyPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	appErr := companyPayload.ValidateCNPJ()
	if appErr != nil {
		w.WriteHeader(appErr.StatusCode)
		json.NewEncoder(w).Encode(appErr.Message)
		return
	}

	appErr = c.companyService.CreateCompany(companyPayload)

	if appErr != nil {
		w.WriteHeader(appErr.StatusCode)
		json.NewEncoder(w).Encode(appErr.Message)
		return
	}
}
