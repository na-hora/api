package handlers

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/company-pet-size/dtos"
	companyPetSizeServices "na-hora/api/internal/models/company-pet-size/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"na-hora/api/internal/utils/conversor"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type CompanyPetSizeHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByCompanyID(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
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

func (cpt *CompanyPetSizeHandler) GetByCompanyID(w http.ResponseWriter, r *http.Request) {
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

	petSizes, appErr := cpt.companyPetSizeService.ListByCompanyID(companyIdConverted, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	var responsePetSizes = make([]dtos.ListPetSizesByCompanyResponse, 0)
	for _, petSize := range petSizes {
		responsePetSizes = append(responsePetSizes, dtos.ListPetSizesByCompanyResponse{
			ID:                 petSize.ID,
			Name:               petSize.Name,
			Description:        petSize.Description,
			CompanyPetTypeID:   petSize.CompanyPetTypeID,
			CompanyPetTypeName: petSize.CompanyPetType.Name,
		})
	}

	utils.ResponseJSON(w, http.StatusOK, responsePetSizes)
}

func (cpt *CompanyPetSizeHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	petSizeId := chi.URLParam(r, "ID")

	parsedPetSizeId, err := strconv.Atoi(petSizeId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := cpt.companyPetSizeService.DeleteByID(parsedPetSizeId, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, nil)
}

func (cpt *CompanyPetSizeHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	petSizeId := chi.URLParam(r, "ID")

	parsedPetSizeId, err := strconv.Atoi(petSizeId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	body := ctx.Value(utils.ValidatedBodyKey).(*dtos.UpdateCompanyPetSizeRequestBody)

	appErr := cpt.companyPetSizeService.UpdateByID(parsedPetSizeId, dtos.UpdateCompanyPetSizeParams{
		Name:        body.Name,
		Description: body.Description,
	}, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, nil)
}
