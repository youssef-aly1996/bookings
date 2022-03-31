package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/youssef-aly1996/bookings/internal/forms"
	"github.com/youssef-aly1996/bookings/internal/models"
	"github.com/youssef-aly1996/bookings/internal/models/reservation"
	"github.com/youssef-aly1996/bookings/internal/render"
)

//Reservation renders the make reservation page template
func (repo *Repository) Reservation(rw http.ResponseWriter, r *http.Request) {
	res, ok := repo.App.Session.Get(r.Context(), "reservation").(reservation.Reservation)
	if !ok {
		repo.Erroring.ServerErrors(rw, errors.New("cannot get reservation from the session"))
		return
	}
	roomName, err := rh.GetById(res.RoomId)
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}
	layout := "2006-01-02"
	sd := res.StartDate.Format(layout)
	ed := res.EndDate.Format(layout)
	strMap := make(map[string]string)
	strMap["start_date"] = sd
	strMap["end_date"] = ed
	strMap["room_name"] = roomName
	data := make(map[string]interface{})
	data["reservation"] = res
	SetCsrf(r)
	td.Form = forms.NewForm(nil)
	td.Data = data
	td.StringMap = strMap
	render.Template(rw, "make-reservation.page.tmpl", td)
}

//PostReservation allows clients to fill out a new reservation form
func (repo *Repository) PostReservation(rw http.ResponseWriter, r *http.Request) {
	res, ok := repo.App.Session.Get(r.Context(), "reservation").(reservation.Reservation)
	if !ok {
		repo.App.Session.Put(r.Context(), "error", "cannot get reservation model from the seesion")
		td.Error = repo.App.Session.PopString(r.Context(), "error")
		http.Redirect(rw, r, "/search-availability", http.StatusTemporaryRedirect)
		return
	}
	err := r.ParseForm()
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}

	res.FirstName = r.Form.Get("first_name")
	res.LastName = r.FormValue("last_name")
	res.Email = r.FormValue("email")
	res.Phone = r.FormValue("phone")

	form := forms.NewForm(r.PostForm)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 5)
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = res
		td.Data = data
		td.Form = form
		render.Template(rw, "make-reservation.page.tmpl", td)
		return
	}
	_, err = rs.Insert(res)
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
	}

	//sending email for room booking
	htmlMsg := fmt.Sprintf(
		`<strong>reservation confirmation</strong><br>
		Dear %s:, <br>
		this completes your reservation from %s to %s.	
		`, res.FirstName, res.StartDate.Format("2006-02-01"), res.EndDate.Format("2006-02-01"))
	mailData := models.MailModel {
		To: res.Email,
		From: "joe@here.com",
		Subject: "reservation confirmation",
		Content: htmlMsg,
	}
	repo.App.MailChan <- mailData
	
	repo.App.Session.Put(r.Context(), "reservation", res)
	repo.App.Session.Put(r.Context(), "success", "your reservation has been completed")
	http.Redirect(rw, r, "/reservation-summary", http.StatusSeeOther)
}

func (repo *Repository) ReservationSummary(rw http.ResponseWriter, r *http.Request) {
	reservation, ok := repo.App.Session.Get(r.Context(), "reservation").(reservation.Reservation)
	if !ok {
		log.Println("cannot get the reservation model from the seesion")
		repo.App.Session.Put(r.Context(), "error", "cannot get reservation model from the seesion")
		td.Error = repo.App.Session.PopString(r.Context(), "error")
		// repo.App.Session.Remove(r.Context(), "error")
		http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
		return
	}
	td.Flash = repo.App.Session.PopString(r.Context(), "success")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	td.Data = data
	render.Template(rw, "reservation-summary.page.tmpl", td)
}


