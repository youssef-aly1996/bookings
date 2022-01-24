package handlers

import (
	"fmt"
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

func (repo *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(rw, "home.page.html", nil)
}

func (repo *Repository) About(rw http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	remoteIp := repo.App.Session.GetString(r.Context(), "remote_ip")
	m["remote_ip"] = remoteIp
	render.RenderTemplate(rw, "about.page.html", nil)
}

func (repo *Repository) Generals(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "generals.page.html", nil)
}

func (repo *Repository) Majors(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "majors.page.html", nil)
}
func (repo *Repository) SearchAvailability(rw http.ResponseWriter, r *http.Request) {
	td.CSRF = nosurf.Token(r)
	render.RenderTemplate(rw, "search-availability.page.html", td)
}

func (repo *Repository) PostAvailability(rw http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	rw.Write([]byte(fmt.Sprintf("start date is %s and end data is %s", start, end)))
}

func (repo *Repository) Reservation(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "make-reservation.page.html", nil)
}
func (repo *Repository) Contact(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "contact.page.html", nil)
}
