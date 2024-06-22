package routes

import (
	"github.com/go-chi/chi"

	"na-hora/api/internal/models/user/handlers"
)

func UserRoutes(r chi.Router) {
	userHandler := handlers.GetUserHandler()

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)
		r.Post("/forgot-password", userHandler.ForgotPassword)
		r.Post("/reset-password", userHandler.ResetPassword)
	})
}
