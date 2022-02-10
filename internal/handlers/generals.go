package handlers

import (
	"net/http"

	"github.com/youssef-aly1996/bookings/internal/render"
)

//Generals renders the generals page template
func (repo *Repository) Generals(rw http.ResponseWriter, r *http.Request) {
	SetCsrf(r)
	render.Template(rw, "generals.page.tmpl", td)
}
