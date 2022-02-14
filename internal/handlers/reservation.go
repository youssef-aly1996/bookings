package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/youssef-aly1996/bookings/internal/forms"
	"github.com/youssef-aly1996/bookings/internal/models/reservation"
	"github.com/youssef-aly1996/bookings/internal/models/roomrestriction"
	"github.com/youssef-aly1996/bookings/internal/render"
)

var res = reservation.Reservation{}

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
	err := r.ParseForm()
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}
	layout := "2006-01-01"
	sd := r.FormValue("start_date")
	ed := r.FormValue("end_date")

	startDate, err := time.Parse(layout, sd)
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}
	rId, _ := strconv.Atoi(r.FormValue("room_id"))

	res.FirstName = r.Form.Get("first_name")
	res.LastName = r.FormValue("last_name")
	res.Email = r.FormValue("email")
	res.Phone = r.FormValue("phone")
	res.StartDate = startDate
	res.EndDate = endDate
	res.RoomId = rId

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
	id, err := rs.Insert(res)
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
	}
	rm := roomrestriction.RoomRestriction{
		StartDate:     startDate,
		EndDate:       endDate,
		RoomId:        rId,
		ReservationId: id,
		RestrictionId: 1,
	}
	err = rr.Insert(rm)
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}

	repo.App.Session.Put(r.Context(), "reservation", res)
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
	repo.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	td.Data = data
	render.Template(rw, "reservation-summary.page.tmpl", td)
}
