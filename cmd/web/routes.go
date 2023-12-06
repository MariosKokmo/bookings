package main

import (
	"net/http"

	"github.com/MariosKokmo/bookings/pkg/config"
	"github.com/MariosKokmo/bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomeTmpl)
	mux.Get("/about", handlers.Repo.AboutTmpl)
	mux.Get("/singleroom", handlers.Repo.SingleRoom)
	mux.Get("/doubleroom", handlers.Repo.DoubleRoom)
	mux.Get("/check-dates", handlers.Repo.CheckDates)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
