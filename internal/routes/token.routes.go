package routes

import (
	"github.com/go-chi/chi"

	"na-hora/api/internal/models/token/handlers"
)

func TokenRoutes(r chi.Router) {
	tokenHandler := handlers.GetTokenHandler()

	r.Route("/tokens", func(r chi.Router) {
		r.Post("/generate", tokenHandler.GenerateRegisterLink)
	})
}
