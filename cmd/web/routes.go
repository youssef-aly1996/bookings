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
	// mux.Use(noSurf)
	// mux.Use(scrfLoad)
	mux.Use(sessionLoad)

	mux.Get("/", repo.Home)
	mux.Get("/about", repo.About)
	mux.Get("/generals-quarters", repo.Generals)
	mux.Get("/majors-suite", repo.Majors)

	mux.Get("/search-availability", repo.SearchAvailability)
	mux.Post("/search-availability", repo.PostAvailability)
	mux.Post("/search-availability-json", repo.CkeckAvailabilityJson)
	mux.Get("/choose-room/{id}", repo.ChooseRoom)
	mux.Get("/book-room", repo.BookRoom)

	mux.Get("/make-reservation", repo.Reservation)
	mux.Post("/make-reservation", repo.PostReservation)
	mux.Get("/reservation-summary", repo.ReservationSummary)

	mux.Get("/contact", repo.Contact)

	mux.Get("/login", repo.Login)
	mux.Post("/login", repo.PostLogin)
	mux.Get("/signup", repo.Signup)
	mux.Post("/signup", repo.PostSignup)
	mux.Get("/logout", repo.Logut)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)
		mux.Get("/dashboard", repo.AdminDashboard)
		mux.Get("/reservations-new", repo.AdminNewReservations)
		mux.Get("/reservations-all", repo.AdminAllReservations)
		mux.Get("/reservations-calender", repo.AdminReservationsCalender)
		mux.Get("/reservations/{src}/{id}", repo.AdminShowReservation)
		mux.Post("/reservations/{src}/{id}", repo.AdminPostShowReservation)
		mux.Delete("/reservations/{src}/{id}", repo.AdminDeleteReservation)
	})


	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
