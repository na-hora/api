package routes

import (
	"na-hora/api/internal/models/appointment/handlers"

	"github.com/go-chi/chi"
)

func AppointmentRoutes(r chi.Router) {
	appointmentHandler := handlers.GetAppointmentHandler()

	r.Route("/appointment", func(r chi.Router) {
		r.Get("/ssr", appointmentHandler.SseUpdates)
	})
}
