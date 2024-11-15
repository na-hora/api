package handlers

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/company-pet-hair/dtos"
	companyPetHairServices "na-hora/api/internal/models/company-pet-hair/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"net/http"
)

type CompanyPetHairHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type CompanyPetHairHandler struct {
	companyPetHairService companyPetHairServices.CompanyPetHairServiceInterface
}

func GetCompanyPetHairHandler() CompanyPetHairHandlerInterface {
	companyPetHairService := injector.InitializeCompanyPetHairService(config.DB)
	return &CompanyPetHairHandler{
		companyPetHairService,
	}
}

func (cphh *CompanyPetHairHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body := ctx.Value(utils.ValidatedBodyKey).(*dtos.CreateCompanyPetHairRequestBody)

	userLogged, userErr := authentication.JwtUserOrThrow(ctx)

	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	appErr := cphh.companyPetHairService.Create(userLogged.CompanyID, *body, nil)

	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, nil)
}
