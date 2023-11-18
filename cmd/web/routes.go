package main

import (
	"net/http"

	"github.com/MariosKokmo/bookings/pkg/config"
	"github.com/MariosKokmo/bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()

	//mux.Get("/", http.HandlerFunc(handlers.Repo.HomeTmpl))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutTmpl))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomeTmpl)
	mux.Get("/about", handlers.Repo.AboutTmpl)
	mux.Get("/personal", handlers.Repo.PersonalDetailsTmpl)
	mux.Get("/personaldetailsDisplay", handlers.Repo.PersonalDetailsDisplayTmpl)

	return mux
}
