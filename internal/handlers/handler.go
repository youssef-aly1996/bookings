package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/youssef-aly1996/bookings/internal/config"
	"github.com/youssef-aly1996/bookings/internal/forms"
	"github.com/youssef-aly1996/bookings/internal/models"
	"github.com/youssef-aly1996/bookings/internal/render"
)

type Repository struct {
	App *config.AppConfig
}

//intializing the tempalate data
var td = models.NewTemplateData()
var EmptyReservation = models.NewReservation()

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func SetCsrf(r *http.Request) {
	td.CSRF = nosurf.Token(r)
}

//Home renders the home page template
func (repo *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(rw, "home.page.tmpl", td)
}

//About renders the about page template
func (repo *Repository) About(rw http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	remoteIp := repo.App.Session.GetString(r.Context(), "remote_ip")
	m["remote_ip"] = remoteIp
	render.RenderTemplate(rw, "about.page.tmpl", td)
}

//Generals renders the generals page template
func (repo *Repository) Generals(rw http.ResponseWriter, r *http.Request) {
	SetCsrf(r)
	render.RenderTemplate(rw, "generals.page.tmpl", td)
}

//Majors renders the major suite page template
func (repo *Repository) Majors(rw http.ResponseWriter, r *http.Request) {
	SetCsrf(r)
	render.RenderTemplate(rw, "majors.page.tmpl", td)
}

//SearchAvailability renders the search availability page template
func (repo *Repository) SearchAvailability(rw http.ResponseWriter, r *http.Request) {
	SetCsrf(r)
	render.RenderTemplate(rw, "search-availability.page.tmpl", td)
}

//PostAvailability renders the search availability page template
func (repo *Repository) PostAvailability(rw http.ResponseWriter, r *http.Request) {
	// start := r.Form.Get("start")
	// end := r.Form.Get("end")
	start := r.FormValue("start")
	end := r.FormValue("end")
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
	data := make(map[string]interface{})
	data["reservation"] = EmptyReservation
	SetCsrf(r)
	td.Form = forms.NewForm(nil)
	td.Data = data
	render.RenderTemplate(rw, "make-reservation.page.tmpl", td)
}

//PostReservation allows clients to fill out a new reservation form
func (repo *Repository) PostReservation(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	EmptyReservation.FirstName = r.Form.Get("first_name")
	EmptyReservation.LastName = r.FormValue("last_name")
	EmptyReservation.Email = r.FormValue("email")
	EmptyReservation.Phone = r.FormValue("phone")

	form := forms.NewForm(r.PostForm)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 5, r)
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = EmptyReservation
		td.Data = data
		td.Form = form
		render.RenderTemplate(rw, "make-reservation.page.tmpl", td)
		return
	}
	repo.App.Session.Put(r.Context(), "reservation", EmptyReservation)
	http.Redirect(rw, r, "/reservation-summary", http.StatusSeeOther)
}

func (repo *Repository) ReservationSummary(rw http.ResponseWriter, r *http.Request) {
	reservation, ok := repo.App.Session.Get(r.Context(), "reservation").(*models.Reservation)
	if !ok {
		log.Println("cannot get the reservation model from the seesion")
		repo.App.Session.Put(r.Context(), "error", "cannot get reservation model from the seesion")
		td.Error = repo.App.Session.PopString(r.Context(), "error")
		http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
		return
	}
	repo.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	td.Data = data
	render.RenderTemplate(rw, "reservation-summary.page.tmpl", td)
}

//Contact renders the contact page template
func (repo *Repository) Contact(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "contact.page.tmpl", td)
}
