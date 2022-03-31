package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/youssef-aly1996/bookings/internal/forms"
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

	form := forms.NewForm(r.PostForm)
	form.Required("start", "end")
	if !form.Valid() {
		repo.App.Session.Put(r.Context(), "error", "input dates are empty")
		td.Error = repo.App.Session.PopString(r.Context(), "error")
		http.Redirect(rw, r, "/search-availability", http.StatusSeeOther)
		return
	}

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

	rooms, err := rs.Check(startDate, endDate)
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}
	if len(rooms) == 0 {
		repo.App.Session.Put(r.Context(), "error", "no reservation available choose another data")
		td.Error = repo.App.Session.PopString(r.Context(), "error")
		http.Redirect(rw, r, "/search-availability", http.StatusSeeOther)
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
		StartDate string `json:"start_date"`
		EndtDate string `json:"end_date"`
		RoomId string `json:"room_id"`
	}
	sd := r.FormValue("start")
	ed := r.Form.Get("end")

	layotu := "2006-01-02"

	startDate, err := time.Parse(layotu, sd)
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}
	endDate, err := time.Parse(layotu, ed)
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}

	roomId := r.Form.Get("room_id")

	rooms, err := rs.CheckByRoomId(startDate, endDate, roomId)
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}
	if len(rooms) == 0 {
		res := jsonRes{Ok: false, Message: "available"}
	jres, err := json.MarshalIndent(res, "", "     ")
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jres)
	return
	}

	res := jsonRes{Ok: true, Message: "available", 
	StartDate: sd, EndtDate: ed, RoomId: roomId}
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

func (repo *Repository) BookRoom(rw http.ResponseWriter, r *http.Request) {
	var res reservation.Reservation
	sd := r.URL.Query().Get("sd")
	ed := r.URL.Query().Get("ed")

	layotu := "2006-01-02"

	startDate, err := time.Parse(layotu, sd)
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}
	endDate, err := time.Parse(layotu, ed)
	if err != nil {
		repo.ServerErrors(rw, err)
		return
	}

	roomId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}

	// roomName, err := rh.GetById(roomId)
	// if err != nil {
	// 	repo.Erroring.ServerErrors(rw, err)
	// 	return
	// }
	res.StartDate = startDate
	res.EndDate = endDate
	res.RoomId = roomId
	repo.App.Session.Put(r.Context(), "reservation", res)
	// repo.App.Session.Put(r.Context(), "room_name", roomName)
	http.Redirect(rw, r, "/make-reservation", http.StatusSeeOther)

}
