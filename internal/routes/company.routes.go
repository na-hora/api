package routes

import (
	"github.com/go-chi/chi"

	companyControllers "na-hora/api/internal/models/company/controllers"
)

func CompanyRoutes(r chi.Router) {
	companyController := companyControllers.GetCompanyController()

	r.Post("/register", companyController.Register)
}
