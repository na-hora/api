package routes

import (
	"github.com/go-chi/chi"

	"na-hora/api/internal/models/service/handlers"
	authentication "na-hora/api/internal/routes/middlewares"
)

func ServiceRoutes(r chi.Router) {
	serviceHandler := handlers.GetServiceHandler()

	authService := authentication.NewAuthService()

	r.Route("/services", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(authService.JwtAuthMiddleware)
			r.Post("/", serviceHandler.Create)
			r.Get("/", serviceHandler.ListAll)
			r.Get("/{id}", serviceHandler.Get)
			r.Put("/{id}", serviceHandler.Update)
		})
	})
}
