package routes

import (
	"na-hora/api/internal/models/company/handlers"

	"github.com/go-chi/chi"
)

func CompanyRoutes(r chi.Router) {
	companyHandler := handlers.GetCompanyHandler()

	r.Route("/companies", func(r chi.Router) {
		r.Post("/register", companyHandler.Register)
	})
}
