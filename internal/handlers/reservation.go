package handlers

import (
	"log"
	"net/http"

	"github.com/youssef-aly1996/bookings/internal/forms"
	"github.com/youssef-aly1996/bookings/internal/models"
	"github.com/youssef-aly1996/bookings/internal/render"
)

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
		repo.ServerErrors(rw, err)
		return
	}
	EmptyReservation.FirstName = r.Form.Get("first_name")
	EmptyReservation.LastName = r.FormValue("last_name")
	EmptyReservation.Email = r.FormValue("email")
	EmptyReservation.Phone = r.FormValue("phone")

	form := forms.NewForm(r.PostForm)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 5)
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
		// repo.App.Session.Remove(r.Context(), "error")
		http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
		return
	}
	repo.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	td.Data = data
	render.RenderTemplate(rw, "reservation-summary.page.tmpl", td)
}
