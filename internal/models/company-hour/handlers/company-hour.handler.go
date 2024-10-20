package handlers

import (
	"encoding/json"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/company-hour/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"net/http"

	companyHourDTOs "na-hora/api/internal/models/company-hour/dtos"

	"github.com/go-playground/validator/v10"
)

type CompanyHourHandlerInterface interface {
	CreateMany(w http.ResponseWriter, r *http.Request)
}

type CompanyHourHandler struct {
	companyHourService services.CompanyHourServiceInterface
}

func GetCompanyHourHandler() CompanyHourHandlerInterface {
	companyHourService := injector.InitializeCompanyHourService(config.DB)

	return &CompanyHourHandler{
		companyHourService,
	}
}

func (chh *CompanyHourHandler) CreateMany(w http.ResponseWriter, r *http.Request) {
	var companyHourPayload companyHourDTOs.CreateCompanyHourRequestBody

	err := json.NewDecoder(r.Body).Decode(&companyHourPayload)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(companyHourPayload)
	if err != nil {
		utils.ResponseValidationErrors(err, w, "body")
		return
	}

	userLogged, userErr := authentication.JwtUserOrThrow(r.Context())

	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	hourErr := chh.companyHourService.CreateManyCompanyHour(companyHourPayload, userLogged.CompanyID, nil)
	if hourErr != nil {
		utils.ResponseJSON(w, hourErr.StatusCode, hourErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, "hours created successfully")
}
