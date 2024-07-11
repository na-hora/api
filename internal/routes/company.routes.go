package routes

import (
	companyHourBlockHandlers "na-hora/api/internal/models/company-hour-block/handlers"
	companyHourHandlers "na-hora/api/internal/models/company-hour/handlers"
	companyHandlers "na-hora/api/internal/models/company/handlers"
	authentication "na-hora/api/internal/routes/middlewares"

	"github.com/go-chi/chi"
)

func CompanyRoutes(r chi.Router) {
	companyHandler := companyHandlers.GetCompanyHandler()
	companyHourHandler := companyHourHandlers.GetCompanyHourHandler()
	companyHourBlockHandler := companyHourBlockHandlers.GetCompanyHourBlockHandler()

	authService := authentication.NewAuthService()

	r.Route("/companies", func(r chi.Router) {
		// Not authenticated routes
		r.Group(func(r chi.Router) {
			r.Post("/register", companyHandler.Register)
		})

		// Authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(authService.JwtAuthMiddleware)
			r.Post("/hour", companyHourHandler.CreateMany)
			r.Post("/hour/block", companyHourBlockHandler.CreateMany)
		})
	})
}
