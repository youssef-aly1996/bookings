package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/youssef-aly1996/bookings/internal/handlers"
)

//noSurf adds csrf protection on all post requests
func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})
	return csrfHandler
}

func scrfLoad(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Form") == "post-it" {
			handlers.SetCsrf(r)
			fmt.Println("post form")
		}
		next.ServeHTTP(rw, r)
	})
}

func sessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if ! repo.IsAutuhroized(r) {
			session.Put(r.Context(), "error", "you need to log in")
			http.Redirect(rw, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(rw, r)
	})
}