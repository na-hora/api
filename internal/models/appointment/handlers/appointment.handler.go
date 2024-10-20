package handlers

import (
	"fmt"
	"na-hora/api/internal/models/appointment/dtos"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/conversor"
	"net/http"
)

type AppointmentHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type AppointmentHandler struct {
	// appointmentService services.AppointmentServiceInterface
}

func GetAppointmentHandler() AppointmentHandlerInterface {
	// appointmentService := injector.InitializeAppointmentService(config.DB)
	return &AppointmentHandler{
		// appointmentService,
	}
}

func (ah *AppointmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body := ctx.Value(utils.ValidatedBodyKey).(*dtos.CreateAppointmentsRequestBody)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s", body.StartDate)))
}

func (ah *AppointmentHandler) List(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startDate")

	if startDate == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, "startDate is required")
		return
	}

	strConv := conversor.GetStringConversor()
	startDateConverted, err := strConv.ToDateTimeTZ(startDate)

	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s", startDateConverted)))
}
