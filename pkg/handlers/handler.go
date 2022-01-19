package handlers

import (
	"net/http"

	"github.com/youssef-aly1996/bookings/pkg/config"
	"github.com/youssef-aly1996/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func (repo *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(rw, "home.page.html", nil)
}

func (repo *Repository) About(rw http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	remoteIp := repo.App.Session.GetString(r.Context(), "remote_ip")
	m["remote_ip"] = remoteIp
	render.RenderTemplate(rw, "about.page.html", m)
}
