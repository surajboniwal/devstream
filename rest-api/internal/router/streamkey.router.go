package router

import (
	"devstream-rest-api/internal/handler"
	"devstream-rest-api/internal/util/appauth"

	"github.com/go-chi/chi/v5"
)

func StreamKeyRoutes(r *chi.Mux, streamKeyHandler *handler.StreamKeyHandler) {
	r.Route("/streamkey", func(router chi.Router) {
		router.Use(appauth.AuthMiddleware)
		router.Get("/", streamKeyHandler.GetStreamKeysForUser)
		router.Post("/", streamKeyHandler.CreateStreamKey)
	})
}
