package routes

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

type Response struct {
	Msg  string
	Code int
}

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
	router.Get("/taskevent", func(w http.ResponseWriter, r *http.Request) {
		res := Response{
			Msg:  "TaskEvent",
			Code: 200,
		}
		jsonResponse, err := json.Marshal(res)
		if err != nil {
			log.Println("JSON marshal error:", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	})
	return router
}
