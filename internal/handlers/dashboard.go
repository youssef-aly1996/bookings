package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/youssef-aly1996/bookings/internal/forms"
	"github.com/youssef-aly1996/bookings/internal/render"
)

//AdminDashboard renders the generals page template
func (repo *Repository) AdminDashboard(rw http.ResponseWriter, r *http.Request) {
	SetCsrf(r)
	render.Template(rw, "admin-dashboard.page.tmpl", td)
}

func (repo *Repository) AdminNewReservations(rw http.ResponseWriter, r *http.Request) {
	ress, err := rs.New()
	if err != nil {
		repo.ServerErrors(rw, err)
	}
	data := make(map[string]interface{})
	data["reservations"] = ress
	td.Data = data
	render.Template(rw, "admin-new-reservations.page.tmpl", td)
}

func (repo *Repository) AdminAllReservations(rw http.ResponseWriter, r *http.Request) {
	ress, err := rs.All()
	if err != nil {
		repo.ServerErrors(rw, err)
	}
	data := make(map[string]interface{})
	data["reservations"] = ress
	td.Data = data
	render.Template(rw, "admin-all-reservations.page.tmpl", td)
}


func (repo *Repository) AdminReservationsCalender(rw http.ResponseWriter, r *http.Request) {
	render.Template(rw, "admin-reservations-calender.page.tmpl", td)
}

func (repo *Repository) AdminShowReservation(rw http.ResponseWriter, r *http.Request) {
	exploded := strings.Split(r.RequestURI, "/")

	id, err := strconv.Atoi(exploded[4])
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}

	src := exploded[3]

	stringMap := make(map[string]string)
	stringMap["src"] = src

	
	// get reservation from the database
	res, err := rs.GetByID(id)
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}
	SetCsrf(r)

	data := make(map[string]interface{})
	data["reservation"] = res
	td.Data = data
	td.StringMap = stringMap
	td.Form = forms.NewForm(nil)
	render.Template(rw, "admin-reservations-show.page.tmpl", td)
}

// AdminPostShowReservation posts a reservation
func (repo *Repository) AdminPostShowReservation(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}

	exploded := strings.Split(r.RequestURI, "/")

	id, err := strconv.Atoi(exploded[4])
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}

	src := exploded[3]

	stringMap := make(map[string]string)
	stringMap["src"] = src

	res, err := rs.GetByID(id)
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}

	res.FirstName = r.Form.Get("first_name")
	res.LastName = r.Form.Get("last_name")
	res.Email = r.Form.Get("email")
	res.Phone = r.Form.Get("phone")

	err = rs.Update(res)
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}
	td.Flash = "Reservation is Updated and Changes Saved"
	// url := fmt.Sprintf("/admin/reservations-%s", src)
	http.Redirect(rw, r, "/admin/reservations-all", http.StatusSeeOther)
	td.Flash = ""
}

func (repo *Repository) AdminDeleteReservation(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}

	exploded := strings.Split(r.RequestURI, "/")

	id, err := strconv.Atoi(exploded[4])
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}

	src := exploded[3]

	stringMap := make(map[string]string)
	stringMap["src"] = src

	err = rs.Delete(id)
	if err != nil {
		repo.Erroring.ServerErrors(rw, err)
		return
	}

	td.Flash = "Reservation is Deleted and Changes Saved"
	http.Redirect(rw, r, fmt.Sprintf("/admin/reservations-%s", src), http.StatusSeeOther)
	td.Flash = ""
}