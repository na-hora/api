package handlers

import (
	"encoding/json"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/providers"

	companyDTOs "na-hora/api/internal/models/company/dtos"
	userDTOs "na-hora/api/internal/models/user/dtos"

	cityServices "na-hora/api/internal/models/city/services"
	companyServices "na-hora/api/internal/models/company/services"
	tokenServices "na-hora/api/internal/models/token/services"
	userServices "na-hora/api/internal/models/user/services"

	"na-hora/api/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CompanyHandlerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type CompanyHandler struct {
	companyService companyServices.CompanyServiceInterface
	userService    userServices.UserServiceInterface
	tokenService   tokenServices.TokenServiceInterface
	cityService    cityServices.CityServiceInterface
}

func GetCompanyHandler() CompanyHandlerInterface {
	companyService := injector.InitializeCompanyService(config.DB)
	userService := injector.InitializeUserService(config.DB)
	tokenService := injector.InitializeTokenService(config.DB)
	cityService := injector.InitializeCityService(config.DB)

	return &CompanyHandler{
		companyService,
		userService,
		tokenService,
		cityService,
	}
}

func (c *CompanyHandler) Register(w http.ResponseWriter, r *http.Request) {
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

	if validatorFound == nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, "validator not found")
		return
	}

	appErr := utils.ValidateCNPJ(companyPayload.CNPJ)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	tx := config.StartTransaction()
	company, companyErr := c.companyService.CreateCompany(companyPayload, tx)
	if companyErr != nil {
		tx.Rollback()
		utils.ResponseJSON(w, companyErr.StatusCode, companyErr.Message)
		return
	}

	if companyPayload.Address != nil {
		cityFound, cityErr := c.cityService.GetByIBGE(companyPayload.Address.CityIBGE)

		if cityErr != nil {
			tx.Rollback()
			utils.ResponseJSON(w, cityErr.StatusCode, cityErr.Message)
			return
		}

		if cityFound == nil {
			tx.Rollback()
			utils.ResponseJSON(w, http.StatusNotFound, "city not found")
			return
		}

		servicePayload := &companyDTOs.CreateCompanyAddressParams{
			ZipCode:      companyPayload.Address.ZipCode,
			CityID:       cityFound.ID,
			Neighborhood: companyPayload.Address.Neighborhood,
			Street:       companyPayload.Address.Street,
			Number:       companyPayload.Address.Number,
			Complement:   companyPayload.Address.Complement,
		}

		_, addressErr := c.companyService.CreateAddress(company.ID, *servicePayload, tx)

		if addressErr != nil {
			tx.Rollback()
			utils.ResponseJSON(w, addressErr.StatusCode, addressErr.Message)
			return
		}
	}

	tokenErr = c.tokenService.UseToken(validatorFound.Key, company.ID, tx)

	response := &companyDTOs.CreateCompanyResponse{
		ID: company.ID,
	}

	if tokenErr != nil {
		tx.Rollback()
		utils.ResponseJSON(w, tokenErr.StatusCode, tokenErr.Message)
		return
	}

	userAlreadyExists, usernameErr := c.userService.GetByUsername(companyPayload.Email)
	if usernameErr != nil {
		tx.Rollback()
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
		}, tx)

		if userErr != nil {
			tx.Rollback()
			utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
			return
		}

		emailProvider := providers.NewEmailProvider()
		emailProvider.SendWelcomeEmail(companyPayload.Email)
	}

	dbInfo := tx.Commit()
	if dbInfo.Error != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, dbInfo.Error.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}
