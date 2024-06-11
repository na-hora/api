package routes

import (
	"github.com/go-chi/chi"

	"na-hora/api/internal/models/token/handlers"
	authentication "na-hora/api/internal/routes/middlewares"
)

func TokenRoutes(r chi.Router) {
	tokenHandler := handlers.GetTokenHandler()

	r.Route("/tokens", func(r chi.Router) {
		r.Use(authentication.JwtAuthentication)

		r.Post("/generate", tokenHandler.GenerateRegisterLink)
	})
}
