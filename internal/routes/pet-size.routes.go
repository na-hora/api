package routes

import (
	"na-hora/api/internal/models/company-pet-size/dtos"
	"na-hora/api/internal/models/company-pet-size/handlers"
	"na-hora/api/internal/routes/middlewares"
	authentication "na-hora/api/internal/routes/middlewares"

	"github.com/go-chi/chi"
)

func PetSizeRoutes(r chi.Router) {
	petSizeHandler := handlers.GetCompanyPetSizeHandler()

	authService := authentication.NewAuthService()

	r.Route("/pet-size", func(r chi.Router) {
		// Not authenticated routes
		r.Group(func(r chi.Router) {})

		// Authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(authService.JwtAuthMiddleware)
			r.With(middlewares.ValidateStructBody(&dtos.CreateCompanyPetSizeRequestBody{})).Post("/", petSizeHandler.Create)
		})
	})
}
