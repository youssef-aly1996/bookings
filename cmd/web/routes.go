package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/youssef-aly1996/bookings/internal/handlers"
)

func routes(repo *handlers.Repository) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(noSurf)
	// mux.Use(scrfLoad)
	mux.Use(sessionLoad)

	mux.Get("/", repo.Home)
	mux.Get("/about", repo.About)
	mux.Get("/generals-quarters", repo.Generals)
	mux.Get("/majors-suite", repo.Majors)

	mux.Get("/search-availability", repo.SearchAvailability)
	mux.Post("/search-availability", repo.PostAvailability)
	mux.Post("/search-availability-json", repo.CkeckAvailabilityJson)

	mux.Get("/make-reservation", repo.Reservation)
	mux.Post("/make-reservation", repo.PostReservation)
	mux.Get("/reservation-summary", repo.ReservationSummary)

	mux.Get("/contact", repo.Contact)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
