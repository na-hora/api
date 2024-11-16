package routes

import (
	"na-hora/api/internal/models/company-pet-hair/dtos"
	"na-hora/api/internal/models/company-pet-hair/handlers"
	"na-hora/api/internal/routes/middlewares"

	"github.com/go-chi/chi"
)

func PetHairRoutes(r chi.Router) {
	petHairHandler := handlers.GetCompanyPetHairHandler()

	authService := middlewares.NewAuthService()

	r.Route("/pet-hair", func(r chi.Router) {
		// Not authenticated routes
		r.Group(func(r chi.Router) {
			r.Get("/", petHairHandler.GetByCompanyID)
		})

		// Authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(authService.JwtAuthMiddleware)
			r.With(middlewares.ValidateStructBody(&dtos.CreateCompanyPetHairRequestBody{})).Post("/", petHairHandler.Create)
			r.With(middlewares.ValidateStructBody(&dtos.UpdateCompanyPetHairRequestBody{})).Put("/{ID}", petHairHandler.UpdateByID)
			r.Delete("/{ID}", petHairHandler.DeleteByID)
		})
	})
}
