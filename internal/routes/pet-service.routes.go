package routes

import (
	"github.com/go-chi/chi"

	"na-hora/api/internal/models/pet-service/dtos"
	petServiceHandlers "na-hora/api/internal/models/pet-service/handlers"

	"na-hora/api/internal/routes/middlewares"
)

func PetServiceRoutes(r chi.Router) {
	petServiceHandler := petServiceHandlers.GetPetServiceHandler()

	authService := middlewares.NewAuthService()

	r.Route("/services/pet", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(authService.JwtAuthMiddleware)
			r.With(middlewares.ValidateStructBody(&dtos.CreatePetServiceRequestBody{})).Post("/", petServiceHandler.Register)
			r.With(middlewares.ValidateStructBody(&dtos.PetServiceValuesRelateRequestBody{})).Post("/{ID}/values", petServiceHandler.RelateValues)
			r.Get("/", petServiceHandler.ListAll)
			r.Get("/{ID}", petServiceHandler.GetByID)
			r.Delete("/{ID}", petServiceHandler.DeleteByID)
			r.With(middlewares.ValidateStructBody(&dtos.UpdatePetServiceRequestBody{})).Put("/{ID}", petServiceHandler.UpdateByID)
		})
	})
}
