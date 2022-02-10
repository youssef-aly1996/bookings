package handlers

import (
	"net/http"

	"github.com/youssef-aly1996/bookings/internal/render"
)

//About renders the about page template
func (repo *Repository) About(rw http.ResponseWriter, r *http.Request) {
	render.Template(rw, "about.page.tmpl", td)
}
