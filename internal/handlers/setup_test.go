package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"github.com/youssef-aly1996/bookings/internal/config"
	"github.com/youssef-aly1996/bookings/internal/render"
)

var appConfig = config.NewAppConfig()
var functions = template.FuncMap{}
var pathToTemplates = "./../../templates"

//handlers loading

var session *scs.SessionManager
var repo = NewRepository(appConfig)

func getRoutes() http.Handler {
	//what i am going to put into the session
	gob.Register(EmptyReservation)
	//creating application session
	appConfig.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction
	appConfig.Session = session

	//creating template cache
	tc, err := CreateTestTempalteCache()
	if err != nil {
		log.Fatal(err)
	}
	appConfig.TempalteCache = tc
	appConfig.UseCache = true
	render.NewTemplate(appConfig)

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)
	// mux.Use(scrfLoad)
	mux.Use(SessionLoad)

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

//noSurf adds csrf protection on all post requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})
	return csrfHandler
}

func ScrfLoad(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Form") == "post-it" {
			SetCsrf(r)
			fmt.Println("post form")
		}
		next.ServeHTTP(rw, r)
	})
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func CreateTestTempalteCache() (map[string]*template.Template, error) {
	tempalteCache := map[string]*template.Template{}
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tempalteSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return nil, err
		}
		if len(matches) > 0 {
			tempalteSet, err = tempalteSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		}
		if err != nil {
			return nil, err
		}
		tempalteCache[name] = tempalteSet
	}
	return tempalteCache, nil
}
