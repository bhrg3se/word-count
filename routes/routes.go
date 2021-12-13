package routes

import (
	"assignment/api"
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api/v1/", func(r chi.Router) {
		r.Post("/count", api.GetMostUsedWords)
	})
	return r
}
