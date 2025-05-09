package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func CreateRouter() *chi.Mux {

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
	}))

	router.Route("/api", func(r chi.Router) {
		router.Post("/event", CreateUser)
		router.Get("/event", GetUser)

	})

	return router
}
