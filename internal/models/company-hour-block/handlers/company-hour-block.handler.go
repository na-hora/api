package handlers

import (
	"encoding/json"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/company-hour-block/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"net/http"

	companyHourBlockDTOs "na-hora/api/internal/models/company-hour-block/dtos"

	"github.com/go-playground/validator/v10"
)

type CompanyHourBlockHandlerInterface interface {
	CreateMany(w http.ResponseWriter, r *http.Request)
}

type CompanyHourBlockHandler struct {
	companyHourBlockService services.CompanyHourBlockServiceInterface
}

func GetCompanyHourBlockHandler() CompanyHourBlockHandlerInterface {
	companyHourBlockService := injector.InitializeCompanyHourBlockService(config.DB)

	return &CompanyHourBlockHandler{
		companyHourBlockService,
	}
}

func (chh *CompanyHourBlockHandler) CreateMany(w http.ResponseWriter, r *http.Request) {
	var companyHourBlockPayload companyHourBlockDTOs.CreateCompanyHourBlockRequestBody

	err := json.NewDecoder(r.Body).Decode(&companyHourBlockPayload)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(companyHourBlockPayload)
	if err != nil {
		utils.ResponseValidationErrors(err, w, "body")
		return
	}

	userLogged, userErr := authentication.JwtUserOrThrow(r.Context())

	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	hourBlockErr := chh.companyHourBlockService.CreateManyCompanyHourBlock(companyHourBlockPayload, userLogged.CompanyID, nil)
	if hourBlockErr != nil {
		utils.ResponseJSON(w, hourBlockErr.StatusCode, hourBlockErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, "hours block created successfully")
}
