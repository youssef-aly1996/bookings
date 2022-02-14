package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/youssef-aly1996/bookings/internal/models/reservation"
	"github.com/youssef-aly1996/bookings/internal/render"
)

//SearchAvailability renders the search availability page template
func (repo *Repository) SearchAvailability(rw http.ResponseWriter, r *http.Request) {
	SetCsrf(r)
	render.Template(rw, "search-availability.page.tmpl", td)
}

//PostAvailability renders the search availability page template
func (repo *Repository) PostAvailability(rw http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02"
	sd := r.FormValue("start")
	ed := r.FormValue("end")

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
	rooms, err := rr.Search(startDate, endDate)
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}
	if len(rooms) == 0 {
		http.Redirect(rw, r, "/make-reservation", http.StatusSeeOther)
		repo.App.Session.Put(r.Context(), "error", "not available")
		return
	}
	data := make(map[string]interface{})
	data["rooms"] = rooms
	td.Data = data
	res := reservation.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}
	repo.App.Session.Put(r.Context(), "reservation", res)
	render.Template(rw, "choose-room.page.tmpl", td)
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
		repo.ServerErrors(rw, err)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jres)
}

func (repo *Repository) ChooseRoom(rw http.ResponseWriter, r *http.Request) {
	roomId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}
	res, ok := repo.App.Session.Get(r.Context(), "reservation").(reservation.Reservation)
	if !ok {
		repo.Erroring.ServerErrors(rw, err)
		return
	}
	res.RoomId = roomId
	repo.App.Session.Put(r.Context(), "reservation", res)
	http.Redirect(rw, r, "/make-reservation", http.StatusSeeOther)

}
