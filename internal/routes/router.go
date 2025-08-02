package routes

import (
	"monad-indexer/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Monad Dev Portfolio API"))
	})

	r.Route("/devs", func(r chi.Router) {
		r.Post("/", handlers.CreateDev)
		r.Get("/", handlers.GetAllDevs)
	})

	return r
}