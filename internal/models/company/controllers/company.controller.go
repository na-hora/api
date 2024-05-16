package controllers

import (
	"encoding/json"
	"na-hora/api/internal/dto"
	usecases "na-hora/api/internal/models/company/usecases"

	"net/http"
)

type CompanyController interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type companyController struct {
	companyUsecase usecases.CompanyUsecase
}

func GetCompanyController() CompanyController {
	companyUsecase := usecases.GetCompanyUsecase()
	return &companyController{
		companyUsecase,
	}
}

func (c *companyController) Register(w http.ResponseWriter, r *http.Request) {
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

	appErr = c.companyUsecase.CreateCompany(companyPayload)

	if appErr != nil {
		w.WriteHeader(appErr.StatusCode)
		json.NewEncoder(w).Encode(appErr.Message)
		return
	}
}
