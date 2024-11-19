package handlers

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/company-pet-hair/dtos"
	companyPetHairServices "na-hora/api/internal/models/company-pet-hair/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"na-hora/api/internal/utils/conversor"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type CompanyPetHairHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByCompanyID(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
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

func (cpt *CompanyPetHairHandler) GetByCompanyID(w http.ResponseWriter, r *http.Request) {
	companyID := r.URL.Query().Get("companyId")

	if companyID == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, "companyId is required")
		return
	}

	strConv := conversor.GetStringConversor()
	companyIdConverted, err := strConv.ToUUID(companyID)

	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	petHairs, appErr := cpt.companyPetHairService.ListByCompanyID(companyIdConverted, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	var responsePetHairs = make([]dtos.ListPetHairsByCompanyResponse, 0)
	for _, petHair := range petHairs {
		responsePetHairs = append(responsePetHairs, dtos.ListPetHairsByCompanyResponse{
			ID:                 petHair.ID,
			Name:               petHair.Name,
			CompanyPetTypeID:   petHair.CompanyPetTypeID,
			CompanyPetTypeName: petHair.CompanyPetType.Name,
		})
	}

	utils.ResponseJSON(w, http.StatusOK, responsePetHairs)
}

func (cpt *CompanyPetHairHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	petHairId := chi.URLParam(r, "ID")

	parsedPetHairId, err := strconv.Atoi(petHairId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := cpt.companyPetHairService.DeleteByID(parsedPetHairId, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, nil)
}

func (cpt *CompanyPetHairHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	petHairId := chi.URLParam(r, "ID")

	parsedPetHairId, err := strconv.Atoi(petHairId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	body := ctx.Value(utils.ValidatedBodyKey).(*dtos.UpdateCompanyPetHairRequestBody)

	appErr := cpt.companyPetHairService.UpdateByID(parsedPetHairId, dtos.UpdateCompanyPetHairParams{
		Name:        body.Name,
		Description: body.Description,
	}, nil)

	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, nil)
}
