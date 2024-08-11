package routes

import (
	"github.com/go-chi/chi"

	petServiceHandlers "na-hora/api/internal/models/pet-service/handlers"

	authentication "na-hora/api/internal/routes/middlewares"
)

func PetServiceRoutes(r chi.Router) {
	petServiceHandler := petServiceHandlers.GetPetServiceHandler()

	authService := authentication.NewAuthService()

	r.Route("/services/pet", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(authService.JwtAuthMiddleware)
			r.Post("/", petServiceHandler.Register)
			r.Get("/{companyId}/list-all", petServiceHandler.ListAll)
			r.Delete("/{companyId}", petServiceHandler.DeleteByID)
		})
	})
}
