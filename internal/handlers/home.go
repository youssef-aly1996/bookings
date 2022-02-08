package handlers

import (
	"net/http"

	"github.com/youssef-aly1996/bookings/internal/render"
)

//Home renders the home page template
func (repo *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.tmpl", td)
}
