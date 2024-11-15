package handlers

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/company-pet-size/dtos"
	companyPetSizeServices "na-hora/api/internal/models/company-pet-size/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"net/http"
)

type CompanyPetSizeHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type CompanyPetSizeHandler struct {
	companyPetSizeService companyPetSizeServices.CompanyPetSizeServiceInterface
}

func GetCompanyPetSizeHandler() CompanyPetSizeHandlerInterface {
	companyPetSizeService := injector.InitializeCompanyPetSizeService(config.DB)
	return &CompanyPetSizeHandler{
		companyPetSizeService,
	}
}

func (cphh *CompanyPetSizeHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body := ctx.Value(utils.ValidatedBodyKey).(*dtos.CreateCompanyPetSizeRequestBody)

	userLogged, userErr := authentication.JwtUserOrThrow(ctx)

	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	appErr := cphh.companyPetSizeService.Create(userLogged.CompanyID, *body, nil)

	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, nil)
}
