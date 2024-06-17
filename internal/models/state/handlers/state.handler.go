package handlers

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/state/dtos"
	"na-hora/api/internal/models/state/services"
	"na-hora/api/internal/utils"
	"net/http"
)

type StateHandlerInterface interface {
	ListAll(w http.ResponseWriter, r *http.Request)
}

type StateHandler struct {
	stateService services.StateServiceInterface
}

func GetStateHandler() StateHandlerInterface {
	stateService := injector.InitializeStateService(config.DB)
	return &StateHandler{
		stateService,
	}
}

func (sh *StateHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	allStates, sErr := sh.stateService.ListAll()
	if sErr != nil {
		utils.ResponseJSON(w, sErr.StatusCode, sErr.Message)
		return
	}

	response := &dtos.ListAllStatesResponse{
		States: []dtos.State{},
	}

	for _, state := range allStates {
		response.States = append(response.States, dtos.State{
			ID:   state.ID,
			Name: state.Name,
			UF:   state.UF,
		})
	}

	utils.ResponseJSON(w, http.StatusCreated, response)
}
