package routesConfig

import (
	"na-hora/api/internal/routes"

	"github.com/go-chi/chi"
)

func VersionedRoutes(r chi.Router, version string) {
	r.Route(version, func(r chi.Router) {
		routes.CompanyRoutes(r)
	})
}
