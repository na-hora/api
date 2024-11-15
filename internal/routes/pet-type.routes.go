package routes

import (
	"github.com/go-chi/chi"

	petTypeHandlers "na-hora/api/internal/models/company-pet-type/handlers"
	authentication "na-hora/api/internal/routes/middlewares"
)

func PetTypeRoutes(r chi.Router) {
	petTypeHandler := petTypeHandlers.GetCompanyPetTypeHandler()

	authService := authentication.NewAuthService()

	r.Route("/pet-type", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(authService.JwtAuthMiddleware)
			r.Post("/", petTypeHandler.Register)
			r.Get("/", petTypeHandler.GetByCompanyID)
			r.Delete("/{ID}", petTypeHandler.DeleteByID)
		})
	})
}
