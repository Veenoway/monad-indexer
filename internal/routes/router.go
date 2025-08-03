package routes

import (
	"encoding/json"
	"monad-indexer/internal/handlers"
	"monad-indexer/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Monad Dev Portfolio API"))
	})

	r.Route("/projects", func(r chi.Router) {
		r.Get("/", handlers.GetAllProjects)
		r.With(middleware.IsAdmin).Post("/", handlers.CreateProject)
	})

	// r.Route("/devs", func(r chi.Router) {
	// 	r.With(middleware.IsAdmin).Post("/", handlers.CreateDev)
		
	// })
	r.Get("/debug/routes", func(w http.ResponseWriter, r *http.Request) {
		// Chi n'a pas de méthode Routes() directe
		// Utilisez cette approche simple pour lister vos routes
		routes := map[string]string{
			"GET /debug/routes": "debug routes",
			"POST /create-dev": "create developer", 
			"GET /devs": "get all developers",
			// Ajoutez vos autres routes ici
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(routes)
	})

	r.Get("/devs", handlers.GetAllDevs)
	r.Post("/create-dev", handlers.CreateDev)

	r.Get("/dev", handlers.GetDev)
	
	return r
}