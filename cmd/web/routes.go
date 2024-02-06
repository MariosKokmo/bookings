package main

import (
	"net/http"

	"github.com/MariosKokmo/bookings/internal/config"
	"github.com/MariosKokmo/bookings/internal/handlers"

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
	mux.Post("/check-dates", handlers.Repo.PostCheckDates)
	mux.Post("/check-dates-json", handlers.Repo.CheckDatesJSON)

	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)
	
	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
