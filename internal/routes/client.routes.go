package routes

import (
	"na-hora/api/internal/models/client/dtos"
	"na-hora/api/internal/models/client/handlers"
	"na-hora/api/internal/routes/middlewares"
	authentication "na-hora/api/internal/routes/middlewares"

	"github.com/go-chi/chi"
)

func ClientRoutes(r chi.Router) {
	clientHandler := handlers.GetClientHandler()

	authService := authentication.NewAuthService()

	r.Route("/clients", func(r chi.Router) {
		// Not authenticated routes
		r.Group(func(r chi.Router) {
			r.With(middlewares.ValidateStructBody(&dtos.CreateClientRequestBody{})).Post("/", clientHandler.Create)
		})

		// Authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(authService.JwtAuthMiddleware)
			r.Get("/", clientHandler.List)
		})
	})
}
