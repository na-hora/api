package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/appointment/dtos"
	appointmentServices "na-hora/api/internal/models/appointment/services"
	petServiceServices "na-hora/api/internal/models/pet-service/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"na-hora/api/internal/utils/conversor"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type AppointmentHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	SseUpdates(w http.ResponseWriter, r *http.Request)
}

type AppointmentHandler struct {
	appointmentService appointmentServices.AppointmentServiceInterface
	petServiceService  petServiceServices.PetServiceServiceInterface
}

func GetAppointmentHandler() AppointmentHandlerInterface {
	appointmentService := injector.InitializeAppointmentService(config.DB)
	petServiceService := injector.InitializePetServiceService(config.DB)
	return &AppointmentHandler{
		appointmentService,
		petServiceService,
	}
}

type ConnectedClient struct {
	CompanyID uuid.UUID
	UserID    uuid.UUID
	Username  string
}

var (
	connectedClients = make(map[http.ResponseWriter]ConnectedClient)
	mu               sync.Mutex
)

func (ah *AppointmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body := ctx.Value(utils.ValidatedBodyKey).(*dtos.CreateAppointmentsRequestBody)

	appointment, appErr := ah.appointmentService.Create(body.CompanyID, *body, nil)

	if appErr != nil {
		utils.ResponseJSON(w, appErr.StatusCode, appErr.Message)
		return
	}

	companyPetService, petServiceErr := ah.petServiceService.GetByID(body.CompanyPetServiceID, nil)

	if petServiceErr != nil {
		utils.ResponseJSON(w, petServiceErr.StatusCode, petServiceErr.Message)
		return
	}

	response := &dtos.CreateAppointmentResponse{
		ID:          appointment.ID,
		StartTime:   appointment.StartTime,
		TotalTime:   appointment.TotalTime,
		ServiceName: companyPetService.Name,
	}

	notificationErr := notifyClients(*response, body.CompanyID)

	if notificationErr != nil {
		utils.ResponseJSON(w, notificationErr.StatusCode, notificationErr.Message)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, response)
}

func (ah *AppointmentHandler) List(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	if startDate == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, "startDate is required")
		return
	}

	if endDate == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, "endDate is required")
		return
	}

	strConv := conversor.GetStringConversor()
	startDateConverted, err := strConv.ToDateTimeTZ(startDate)

	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	endDateConverted, err := strConv.ToDateTimeTZ(endDate)

	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	userLogged, userErr := authentication.JwtUserOrThrow(r.Context())
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	appointments, appError := ah.appointmentService.List(
		userLogged.CompanyID,
		startDateConverted,
		endDateConverted,
	)

	if appError != nil {
		utils.ResponseJSON(w, appError.StatusCode, appError.Message)
		return
	}

	response := &dtos.ListAppointmentsResponse{
		Appointments: []dtos.Appointment{},
	}

	for _, appointment := range appointments {
		response.Appointments = append(response.Appointments, dtos.Appointment{
			ID:          appointment.ID,
			ServiceName: appointment.CompanyPetServiceValue.CompanyPetService.Name,
			StartTime:   appointment.StartTime.Format("2006-01-02 15:04:05"),
			TotalTime:   appointment.TotalTime,
			TotalPrice:  appointment.TotalPrice,
			Canceled:    appointment.Canceled,
		})
	}

	utils.ResponseJSON(w, http.StatusOK, response)
}

func (ah *AppointmentHandler) SseUpdates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ctx := r.Context()

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	stringToken := r.URL.Query().Get("token")

	if stringToken == "" {
		utils.ResponseJSON(w, http.StatusUnauthorized, "token is required")
	}

	userLogged, userErr := authentication.UserFromStringToken(stringToken)
	if userErr != nil {
		utils.ResponseJSON(w, userErr.StatusCode, userErr.Message)
		return
	}

	mu.Lock()
	connectedClients[w] = ConnectedClient{
		CompanyID: userLogged.CompanyID,
		UserID:    userLogged.ID,
		Username:  userLogged.Username,
	}
	mu.Unlock()

	defer func() {
		mu.Lock()
		if _, exists := connectedClients[w]; exists {
			delete(connectedClients, w)
			fmt.Println("Client disconnected and removed:", userLogged.Username)
		}
		mu.Unlock()
	}()

	go func() {
		<-ctx.Done()
		mu.Lock()
		if _, exists := connectedClients[w]; exists {
			delete(connectedClients, w)
			fmt.Println("Removed client on disconnect:", userLogged.Username)
		}
		mu.Unlock()
	}()

	select {}
}

func notifyClients(
	appointment dtos.CreateAppointmentResponse,
	companyID uuid.UUID,
) *utils.AppError {
	mu.Lock()
	for clientWriter, clientInfo := range connectedClients {
		if clientInfo.CompanyID == companyID {
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)

			if err := enc.Encode(appointment); err != nil {

				mu.Unlock()
				return &utils.AppError{
					Message:    "Failed to encode appointment",
					StatusCode: http.StatusInternalServerError,
				}
			}

			fmt.Fprintf(clientWriter, "data: %s\n\n", buf.String())
			if f, ok := clientWriter.(http.Flusher); ok {
				f.Flush()
			}
		}
	}
	mu.Unlock()
	return nil
}
