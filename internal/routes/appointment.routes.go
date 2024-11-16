package routes

import (
	"na-hora/api/internal/models/appointment/dtos"
	"na-hora/api/internal/models/appointment/handlers"
	"na-hora/api/internal/routes/middlewares"

	"github.com/go-chi/chi"
)

func AppointmentRoutes(r chi.Router) {
	appointmentHandler := handlers.GetAppointmentHandler()

	authService := middlewares.NewAuthService()

	r.Route("/appointments", func(r chi.Router) {
		// Not authenticated routes
		r.Group(func(r chi.Router) {
			r.Get("/notifications", appointmentHandler.SseUpdates)
			r.With(middlewares.ValidateStructBody(&dtos.CreateAppointmentsRequestBody{})).Post("/", appointmentHandler.Create)
		})

		// Authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(authService.JwtAuthMiddleware)
			r.Get("/", appointmentHandler.List)
		})
	})
}
