package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func UserRoutes(r chi.Router) {

	r.Route("/users", func(r chi.Router) {
		// r.Post("/login", userControllers.Login)
		fmt.Println("login user")
	})
}
