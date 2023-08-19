package router

import (
	"devstream-rest-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r *chi.Mux, authHandler *handler.AuthHandler) {
	r.Route("/auth", func(router chi.Router) {
		router.Post("/login", authHandler.Login)
		router.Post("/register", authHandler.Register)
	})
}
