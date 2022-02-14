package handlers

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/youssef-aly1996/bookings/internal/config"
	"github.com/youssef-aly1996/bookings/internal/db/pgrepo"
	"github.com/youssef-aly1996/bookings/internal/erroring"
	"github.com/youssef-aly1996/bookings/internal/models"
	"github.com/youssef-aly1996/bookings/internal/models/reservation"
	"github.com/youssef-aly1996/bookings/internal/models/room"
	"github.com/youssef-aly1996/bookings/internal/models/roomrestriction"
)

var (
	//intializing the tempalate data
	td          = models.NewTemplateData()
	dbrepo, err = pgrepo.NewPgRepo()
	rs          = reservation.New(dbrepo)
	rr          = roomrestriction.New(dbrepo)
	rh          = room.New(dbrepo)
)

type Repository struct {
	App *config.AppConfig
	erroring.Erroring
}

func NewRepository(a *config.AppConfig) *Repository {
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to database")
	return &Repository{App: a, Erroring: erroring.NewErroring(a)}
}

func SetCsrf(r *http.Request) {
	td.CSRF = nosurf.Token(r)
}
