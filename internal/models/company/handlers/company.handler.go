package handlers

import (
	"encoding/json"
	companyDTOs "na-hora/api/internal/models/company/dtos"
	"na-hora/api/internal/models/company/services"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
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
		utils.ResponseJSON(w, http.StatusBadRequest, "Invalid body")
		return
	}

	validate := validator.New()
	err = validate.Struct(companyPayload)
	if err != nil {
		utils.ResponseValidationErrors(err, w)
		return
	}

	appErr := utils.ValidateCNPJ(companyPayload.CNPJ)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	company, serviceErr := c.companyService.CreateCompany(companyPayload)
	if serviceErr != nil {
		utils.ResponseJSON(w, serviceErr.StatusCode, serviceErr.Message)
		return
	}

	response := &companyDTOs.CreateCompanyResponse{
		ID: company.ID,
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}
