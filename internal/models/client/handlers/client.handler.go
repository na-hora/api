package handlers

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/client/dtos"
	clientServices "na-hora/api/internal/models/client/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"na-hora/api/internal/utils/conversor"
	"net/http"
)

type ClientHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	GetByEmail(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type ClientHandler struct {
	clientService clientServices.ClientServiceInterface
}

func GetClientHandler() ClientHandlerInterface {
	clientService := injector.InitializeClientService(config.DB)
	return &ClientHandler{
		clientService,
	}
}

func (ch *ClientHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body := ctx.Value(utils.ValidatedBodyKey).(*dtos.CreateClientRequestBody)

	appointment, appErr := ch.clientService.Create(body.CompanyID, *body, nil)

	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	response := &dtos.CreateClientResponse{
		ID:        appointment.ID,
		Name:      appointment.Name,
		Phone:     appointment.Phone,
		Email:     appointment.Email,
		CompanyID: appointment.CompanyID,
	}

	utils.ResponseJSON(w, http.StatusOK, response)
}

func (ch *ClientHandler) List(w http.ResponseWriter, r *http.Request) {
	userLogged, userErr := authentication.JwtUserOrThrow(r.Context())
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	clients, appError := ch.clientService.List(
		userLogged.CompanyID,
	)

	if appError != nil {
		utils.ResponseJSON(w, appError.StatusCode, appError.Message)
		return
	}

	response := &dtos.ListClientsResponse{
		Clients: []dtos.Client{},
	}

	for _, client := range clients {
		response.Clients = append(response.Clients, dtos.Client{
			ID:        client.ID,
			Name:      client.Name,
			Phone:     client.Phone,
			Email:     client.Email,
			CompanyID: client.CompanyID,
		})
	}

	utils.ResponseJSON(w, http.StatusOK, response)
}

func (ch *ClientHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, "email is required")
		return
	}

	stringCompanyID := r.URL.Query().Get("companyId")
	if stringCompanyID == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, "companyId is required")
		return
	}

	strConv := conversor.GetStringConversor()
	companyID, err := strConv.ToUUID(stringCompanyID)

	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	client, appError := ch.clientService.GetByEmail(
		companyID,
		email,
	)

	if appError != nil {
		utils.ResponseJSON(w, appError.StatusCode, appError.Message)
		return
	}

	if client == nil {
		utils.ResponseJSON(w, http.StatusNotFound, "client not found")
		return
	}

	response := &dtos.GetUniqueClientResponse{
		ID:        client.ID,
		Name:      client.Name,
		Phone:     client.Phone,
		Email:     client.Email,
		CompanyID: client.CompanyID,
	}

	utils.ResponseJSON(w, http.StatusOK, response)
}

func (ch *ClientHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userLogged, userErr := authentication.JwtUserOrThrow(ctx)
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}
	body := ctx.Value(utils.ValidatedBodyKey).(*dtos.UpdateClientRequestBody)

	client, appErr := ch.clientService.Update(userLogged.CompanyID, *body, nil)

	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	response := &dtos.UpdateClientResponse{
		ID: client.ID,
	}

	utils.ResponseJSON(w, http.StatusOK, response)
}
