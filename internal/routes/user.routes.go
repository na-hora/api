package routes

import (
	userControllers "na-hora/api/internal/models/users/controllers"
	"net/http"

	"github.com/go-chi/chi"
)

func UserRoutes(r chi.Router) {
	r.Post("/login", userControllers.Login)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("users"))
		})
	})
}
