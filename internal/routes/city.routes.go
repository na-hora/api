package routes

import (
	"na-hora/api/internal/models/city/handlers"

	"github.com/go-chi/chi"
)

func CityRoutes(r chi.Router) {
	cityHandler := handlers.GetCityHandler()

	r.Route("/cities", func(r chi.Router) {
		r.Get("/state/{stateID}", cityHandler.ListAllByState)
		r.Get("/ibge/{ibge}", cityHandler.GetByIBGE)
	})
}
