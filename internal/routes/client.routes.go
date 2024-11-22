package routes

import (
	"na-hora/api/internal/models/client/dtos"
	"na-hora/api/internal/models/client/handlers"
	"na-hora/api/internal/routes/middlewares"

	"github.com/go-chi/chi"
)

func ClientRoutes(r chi.Router) {
	clientHandler := handlers.GetClientHandler()

	authService := middlewares.NewAuthService()

	r.Route("/clients", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.With(middlewares.ValidateStructBody(&dtos.CreateClientRequestBody{})).Post("/", clientHandler.Create)
			r.Get("/by-email", clientHandler.GetByEmail)
		})

		r.Group(func(r chi.Router) {
			r.Use(authService.JwtAuthMiddleware)
			r.Get("/", clientHandler.List)
			r.With(middlewares.ValidateStructBody(&dtos.UpdateClientRequestBody{})).Put("/", clientHandler.Update)
		})
	})
}
