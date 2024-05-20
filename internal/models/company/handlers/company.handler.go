package handlers

import (
	"encoding/json"
	companyDTOs "na-hora/api/internal/models/company/dtos"
	companyServices "na-hora/api/internal/models/company/services"
	userDTOs "na-hora/api/internal/models/user/dtos"
	userServices "na-hora/api/internal/models/user/services"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CompanyHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type companyHandler struct {
	companyService companyServices.CompanyService
	userService    userServices.UserService
}

func GetCompanyHandler() CompanyHandler {
	companyService := companyServices.GetCompanyService()
	userService := userServices.GetUserService()
	return &companyHandler{
		companyService,
		userService,
	}
}

func (c *companyHandler) Register(w http.ResponseWriter, r *http.Request) {
	var companyPayload companyDTOs.CreateCompanyRequestBody

	err := json.NewDecoder(r.Body).Decode(&companyPayload)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, "Invalid body")
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
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

	if companyPayload.Address != nil {
		_, serviceErr = c.companyService.CreateAddress(company.ID, *companyPayload.Address)

		if serviceErr != nil {
			utils.ResponseJSON(w, serviceErr.StatusCode, serviceErr.Message)
			return
		}
	}

	_, userErr := c.userService.Create(userDTOs.CreateUserRequestBody{
		Username:  companyPayload.Email,
		Password:  companyPayload.Password,
		CompanyID: company.ID,
	})
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	response := &companyDTOs.CreateCompanyResponse{
		ID: company.ID,
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}
