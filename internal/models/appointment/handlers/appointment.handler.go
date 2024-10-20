package handlers

import (
	"fmt"
	config "na-hora/api/configs"
	"na-hora/api/internal/injector"
	"na-hora/api/internal/models/appointment/dtos"
	"na-hora/api/internal/models/appointment/services"
	"na-hora/api/internal/utils"
	"na-hora/api/internal/utils/authentication"
	"na-hora/api/internal/utils/conversor"
	"net/http"
)

type AppointmentHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type AppointmentHandler struct {
	appointmentService services.AppointmentServiceInterface
}

func GetAppointmentHandler() AppointmentHandlerInterface {
	appointmentService := injector.InitializeAppointmentService(config.DB)
	return &AppointmentHandler{
		appointmentService,
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
			ID:                appointment.ID,
			PetName:           appointment.PetName,
			StartTime:         appointment.StartTime.Format("2006-01-02 15:04:05"),
			TotalTime:         appointment.TotalTime,
			TotalPrice:        appointment.TotalPrice,
			PaymentMode:       appointment.PaymentMode,
			Canceled:          appointment.Canceled,
			CancelationReason: appointment.CancelationReason,
		})
	}

	utils.ResponseJSON(w, http.StatusOK, response)
}
