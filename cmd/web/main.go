package main

import (
	"encoding/gob"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/youssef-aly1996/bookings/internal/config"
	"github.com/youssef-aly1996/bookings/internal/handlers"
	"github.com/youssef-aly1996/bookings/internal/models"
	"github.com/youssef-aly1996/bookings/internal/models/reservation"
	"github.com/youssef-aly1996/bookings/internal/models/user"
	"github.com/youssef-aly1996/bookings/internal/render"
)

var (
	appConfig = config.NewAppConfig()
	//handlers loading
	repo = handlers.NewRepository(appConfig)

	session *scs.SessionManager
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer close(appConfig.MailChan)
	listenForMail()
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
func run() error {
	//Logging info and error info part
	appConfig.Logger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	appConfig.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//what i am going to put into the session
	gob.Register(reservation.Reservation{})
	gob.Register(user.User{})

	//listening for email signals
	mailChan := make(chan models.MailModel)
	appConfig.MailChan = mailChan

	inProduction := flag.Bool("production", true, "Application is in production")
	useCache := flag.Bool("cache", true, "Use template cache")
	appConfig.InProduction = *inProduction
	appConfig.UseCache = *useCache
	flag.Parse()

	//creating application session
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
		return err
	}
	appConfig.TempalteCache = tc
	render.NewTemplate(appConfig)

	return nil
}
