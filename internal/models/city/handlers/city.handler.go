package handlers

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/city/dtos"
	"na-hora/api/internal/models/city/services"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/go-chi/chi"
)

type CityHandlerInterface interface {
	ListAllByState(w http.ResponseWriter, r *http.Request)
}

type CityHandler struct {
	cityService services.CityServiceInterface
}

func GetCityHandler() CityHandlerInterface {
	cityService := injector.InitializeCityService(config.DB)
	return &CityHandler{
		cityService,
	}
}

func (ch *CityHandler) ListAllByState(w http.ResponseWriter, r *http.Request) {
	stateID := chi.URLParam(r, "stateID")

	if stateID == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, "stateID is required")
		return
	}

	stateIDUint, _ := utils.StringToUint(stateID)

	allCities, sErr := ch.cityService.ListAllByState(uint(stateIDUint))
	if sErr != nil {
		utils.ResponseJSON(w, sErr.StatusCode, sErr.Message)
		return
	}

	response := &dtos.ListCitiesByStateResponse{
		Cities: []dtos.City{},
	}

	for _, city := range allCities {
		response.Cities = append(response.Cities, dtos.City{
			ID:   city.ID,
			Name: city.Name,
		})
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}
