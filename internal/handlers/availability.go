package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/youssef-aly1996/bookings/internal/render"
)

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
		repo.ServerErrors(rw, err)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jres)
}
