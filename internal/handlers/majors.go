package handlers

import (
	"net/http"

	"github.com/youssef-aly1996/bookings/internal/render"
)

//Majors renders the major suite page template
func (repo *Repository) Majors(rw http.ResponseWriter, r *http.Request) {
	SetCsrf(r)
	render.RenderTemplate(rw, "majors.page.tmpl", td)
}
