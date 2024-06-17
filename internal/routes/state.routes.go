package routes

import (
	"na-hora/api/internal/models/state/handlers"

	"github.com/go-chi/chi"
)

func StateRoutes(r chi.Router) {
	stateHandler := handlers.GetStateHandler()

	r.Route("/states", func(r chi.Router) {
		r.Get("/", stateHandler.ListAll)
	})
}
