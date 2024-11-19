package routes

import (
	"github.com/go-chi/chi"

	"na-hora/api/internal/models/company-pet-type/dtos"
	petTypeHandlers "na-hora/api/internal/models/company-pet-type/handlers"
	"na-hora/api/internal/routes/middlewares"
)

func PetTypeRoutes(r chi.Router) {
	petTypeHandler := petTypeHandlers.GetCompanyPetTypeHandler()

	authService := middlewares.NewAuthService()

	r.Route("/pet-type", func(r chi.Router) {
		// Not authenticated routes
		r.Group(func(r chi.Router) {
			r.Get("/", petTypeHandler.GetByCompanyID)
		})

		// Authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(authService.JwtAuthMiddleware)
			r.With(middlewares.ValidateStructBody(&dtos.CreatePetTypeRequestBody{})).Post("/", petTypeHandler.Register)
			r.Get("/{ID}/combinations", petTypeHandler.GetValuesCombinations)
			r.Put("/{ID}", petTypeHandler.UpdateByID)
			r.Delete("/{ID}", petTypeHandler.DeleteByID)
		})
	})
}
