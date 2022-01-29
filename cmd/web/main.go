package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/youssef-aly1996/bookings/internal/config"
	"github.com/youssef-aly1996/bookings/internal/handlers"
	"github.com/youssef-aly1996/bookings/internal/render"
)

var appConfig = config.NewAppConfig()
var session *scs.SessionManager

func main() {
	//what i am going to put into the session
	gob.Register(handlers.EmptyReservation)
	//creating application session
	appConfig.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction
	appConfig.Session = session

	//creating template cache
	tc, err := render.CreateTempalteCache()
	if err != nil {
		log.Fatal(err)
	}
	appConfig.TempalteCache = tc
	appConfig.UseCache = false
	render.NewTemplate(appConfig)

	//handlers loading
	repo := handlers.NewRepository(appConfig)
	appConfig.PortNumber = ":3000"
	server := &http.Server{
		Addr:    appConfig.PortNumber,
		Handler: routes(repo),
	}
	log.Println("server is up and running on port 3000")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
