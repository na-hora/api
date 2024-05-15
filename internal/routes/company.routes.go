package routes

import (
	companyControllers "na-hora/api/internal/models/company/controllers"

	"github.com/go-chi/chi"
)

func CompanyRoutes(r chi.Router) {
	r.Post("/register", companyControllers.Register)
}
