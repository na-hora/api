package handlers

import (
	"encoding/json"
	config "na-hora/api/configs"
	"na-hora/api/internal/entity"
	"na-hora/api/internal/injector"
	companyPetHairServices "na-hora/api/internal/models/company-pet-hair/services"
	companyPetSizeServices "na-hora/api/internal/models/company-pet-size/services"
	"na-hora/api/internal/models/company-pet-type/dtos"
	companyPetTypeServices "na-hora/api/internal/models/company-pet-type/services"
	companyPetServiceServices "na-hora/api/internal/models/pet-service/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"na-hora/api/internal/utils/conversor"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type CompanyPetTypeInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	GetByCompanyID(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
	GetValuesCombinations(w http.ResponseWriter, r *http.Request)
}

type CompanyPetTypeHandler struct {
	companyPetTypeService    companyPetTypeServices.CompanyPetTypeServiceInterface
	companyPetHairService    companyPetHairServices.CompanyPetHairServiceInterface
	companyPetSizeServices   companyPetSizeServices.CompanyPetSizeServiceInterface
	companyPetServiceService companyPetServiceServices.PetServiceServiceInterface
}

func GetCompanyPetTypeHandler() CompanyPetTypeInterface {
	companyPetTypeService := injector.InitializeCompanyPetTypeService(config.DB)
	companyPetHairService := injector.InitializeCompanyPetHairService(config.DB)
	companyPetSizeServices := injector.InitializeCompanyPetSizeService(config.DB)
	companyPetServiceServices := injector.InitializePetServiceService(config.DB)

	return &CompanyPetTypeHandler{
		companyPetTypeService,
		companyPetHairService,
		companyPetSizeServices,
		companyPetServiceServices,
	}
}

func (cpt *CompanyPetTypeHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userLogged, userErr := authentication.JwtUserOrThrow(ctx)
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	petTypePayload := ctx.Value(utils.ValidatedBodyKey).(*dtos.CreatePetTypeRequestBody)

	appErr := cpt.companyPetTypeService.Create(userLogged.CompanyID, petTypePayload.Name, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
	}

	utils.ResponseJSON(w, http.StatusCreated, nil)
}

func (cpt *CompanyPetTypeHandler) GetByCompanyID(w http.ResponseWriter, r *http.Request) {
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

	petTypes, appErr := cpt.companyPetTypeService.GetByCompanyID(companyIdConverted)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	var responsePetTypes = make([]dtos.ListPetTypesByCompanyResponse, 0)
	for _, petType := range petTypes {
		responsePetTypes = append(responsePetTypes, dtos.ListPetTypesByCompanyResponse{
			ID:   petType.ID,
			Name: petType.Name,
		})
	}

	utils.ResponseJSON(w, http.StatusOK, responsePetTypes)
}

func (cpt *CompanyPetTypeHandler) GetValuesCombinations(w http.ResponseWriter, r *http.Request) {
	petTypeId := chi.URLParam(r, "ID")
	petServiceId := r.URL.Query().Get("petServiceId")

	strConv := conversor.GetStringConversor()
	parsedPetTypeId, err := strConv.ToInt(petTypeId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	parsedServiceId := 0
	if petServiceId != "" {
		parsedServiceId, err = strConv.ToInt(petServiceId)
		if err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	sizes, sizesErr := cpt.companyPetSizeServices.ListByPetTypeID(parsedPetTypeId, nil)
	if sizesErr != nil {
		utils.ResponseJSON(w, sizesErr.StatusCode, sizesErr.Message)
		return
	}

	hairs, hairsErr := cpt.companyPetHairService.ListByPetTypeID(parsedPetTypeId, nil)
	if hairsErr != nil {
		utils.ResponseJSON(w, hairsErr.StatusCode, hairsErr.Message)
		return
	}

	var responseCombinations = make([]dtos.ListPetTypeCombinationsResponse, 0)
	for _, size := range sizes {
		for _, hair := range hairs {
			var configFound *entity.CompanyPetServiceValue
			if parsedServiceId != 0 {
				found, findErr := cpt.companyPetServiceService.GetConfigurationBySizeAndHair(parsedServiceId, size.ID, hair.ID, nil)
				if findErr != nil {
					utils.ResponseJSON(w, findErr.StatusCode, findErr.Message)
					return
				}
				if found != nil {
					configFound = found
				}
			}

			toAppend := dtos.ListPetTypeCombinationsResponse{
				Hair: dtos.PetTypeCombinationsHair{ID: hair.ID, Name: hair.Name},
				Size: dtos.PetTypeCombinationsSize{ID: size.ID, Name: size.Name},
			}

			if configFound != nil {
				toAppend.Price = &configFound.Price
				toAppend.ExecutionTime = &configFound.ExecutionTime
			}
			responseCombinations = append(responseCombinations, toAppend)
		}
	}

	utils.ResponseJSON(w, http.StatusOK, responseCombinations)
}

func (cpt *CompanyPetTypeHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	petTypeId := chi.URLParam(r, "ID")

	parsedPetTypeId, err := strconv.Atoi(petTypeId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := cpt.companyPetTypeService.DeleteByID(parsedPetTypeId, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, nil)
}

func (cpt *CompanyPetTypeHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	petTypeId := chi.URLParam(r, "ID")

	parsedPetTypeId, err := strconv.Atoi(petTypeId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var body dtos.CreateCompanyPetTypeParams

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := cpt.companyPetTypeService.UpdateByID(parsedPetTypeId, body, nil)
	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, nil)
}
