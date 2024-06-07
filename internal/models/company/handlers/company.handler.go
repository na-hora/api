package handlers

import (
	"encoding/json"
	companyDTOs "na-hora/api/internal/models/company/dtos"
	companyServices "na-hora/api/internal/models/company/services"
	tokenServices "na-hora/api/internal/models/token/services"
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
	tokenService   tokenServices.TokenService
}

func GetCompanyHandler() CompanyHandler {
	companyService := companyServices.GetCompanyService()
	userService := userServices.GetUserService()
	tokenService := tokenServices.GetTokenService()
	return &companyHandler{
		companyService,
		userService,
		tokenService,
	}
}

func (c *companyHandler) Register(w http.ResponseWriter, r *http.Request) {
	var companyPayload companyDTOs.CreateCompanyRequestBody

	err := json.NewDecoder(r.Body).Decode(&companyPayload)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(companyPayload)
	if err != nil {
		utils.ResponseValidationErrors(err, w)
		return
	}

	validatorFound, tokenErr := c.tokenService.GetValidToken(companyPayload.Validator)
	if tokenErr != nil {
		utils.ResponseJSON(w, tokenErr.StatusCode, tokenErr.Message)
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

	tokenErr = c.tokenService.UseToken(validatorFound.Key, company.ID)

	response := &companyDTOs.CreateCompanyResponse{
		ID: company.ID,
	}

	if tokenErr != nil {
		utils.ResponseJSON(w, tokenErr.StatusCode, tokenErr.Message)
		return
	}

	userAlreadyExists, usernameErr := c.userService.GetByUsername(companyPayload.Email)
	if usernameErr != nil {
		utils.ResponseJSON(w, usernameErr.StatusCode, usernameErr.Message)
		return
	}

	if userAlreadyExists != nil {
		response.Inconsistency = "The company and address were created successfully, but the username was already taken"
	} else {
		_, userErr := c.userService.Create(userDTOs.CreateUserRequestBody{
			Username:  companyPayload.Email,
			Password:  companyPayload.Password,
			CompanyID: company.ID,
		})
		if userErr != nil {
			utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
			return
		}
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}
