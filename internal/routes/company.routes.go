package routes

import (
	"github.com/go-chi/chi"

	"na-hora/api/internal/models/company/handlers"
)

func CompanyRoutes(r chi.Router) {
	companyHandler := handlers.GetCompanyHandler()

	r.Route("/companies", func(r chi.Router) {
		r.Post("/register", companyHandler.Register)
	})
}
