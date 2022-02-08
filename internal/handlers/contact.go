package handlers

import (
	"net/http"

	"github.com/youssef-aly1996/bookings/internal/render"
)

//Contact renders the contact page template
func (repo *Repository) Contact(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "contact.page.tmpl", td)
}
