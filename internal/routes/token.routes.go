package routes

import (
	"github.com/go-chi/chi"

	"na-hora/api/internal/models/token/handlers"
	authentication "na-hora/api/internal/routes/middlewares"
)

func TokenRoutes(r chi.Router) {
	tokenHandler := handlers.GetTokenHandler()
	authService := authentication.NewAuthService()

	r.Route("/tokens", func(r chi.Router) {
		r.Use(authService.JwtAuthMiddleware)

		r.Post("/generate", tokenHandler.GenerateRegisterLink)
	})
}
