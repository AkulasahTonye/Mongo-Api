package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/mongo-Api/handlers"
	"net/http"
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
		router.Post("/signup", handlers.CreateUserHandler)
		router.Post("/login", handlers.LoginHandler)

		router.Get("/users", handlers.GetAllUsersHandler)
		router.Get("/user/{id}", handlers.GetUserHandler)
		router.Put("/user/{id}", handlers.UpdateUserHandler)
		router.Delete("/user/{id}", handlers.DeleteUserHandler)

		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, world!"))
		})

	})

	return router
}
