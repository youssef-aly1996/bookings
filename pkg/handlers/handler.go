package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/youssef-aly1996/bookings/pkg/config"
	"github.com/youssef-aly1996/bookings/pkg/models"
	"github.com/youssef-aly1996/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

//intializing the tempalate data
var td = models.NewTemplateData()

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

//Home renders the home page template
func (repo *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(rw, "home.page.html", nil)
}

//About renders the about page template
func (repo *Repository) About(rw http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	remoteIp := repo.App.Session.GetString(r.Context(), "remote_ip")
	m["remote_ip"] = remoteIp
	render.RenderTemplate(rw, "about.page.html", nil)
}

//Generals renders the generals page template
func (repo *Repository) Generals(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "generals.page.html", nil)
}

//Majors renders the major suite page template
func (repo *Repository) Majors(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "majors.page.html", nil)
}

//SearchAvailability renders the search availability page template
func (repo *Repository) SearchAvailability(rw http.ResponseWriter, r *http.Request) {
	td.CSRF = nosurf.Token(r)
	render.RenderTemplate(rw, "search-availability.page.html", td)
}

//PostAvailability renders the search availability page template
func (repo *Repository) PostAvailability(rw http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	rw.Write([]byte(fmt.Sprintf("start date is %s and end data is %s", start, end)))
}

//CkeckAvailabilityJson handels requests for availability and sends json response
func (repo *Repository) CkeckAvailabilityJson(rw http.ResponseWriter, r *http.Request) {
	type jsonRes struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
	}
	res := jsonRes{Ok: true, Message: "available"}
	jres, err := json.MarshalIndent(res, "", "     ")
	if err != nil {
		log.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jres)
}

//Reservation renders the make reservation page template
func (repo *Repository) Reservation(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "make-reservation.page.html", nil)
}

//Contact renders the contact page template
func (repo *Repository) Contact(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "contact.page.html", nil)
}
