package handlers

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/youssef-aly1996/bookings/internal/config"
	"github.com/youssef-aly1996/bookings/internal/erroring"
	"github.com/youssef-aly1996/bookings/internal/models"
)

var (
	//intializing the tempalate data
	td = models.NewTemplateData()
)

type Repository struct {
	App *config.AppConfig
	erroring.Erroring
}

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{App: a, Erroring: erroring.NewErroring(a)}
}

func SetCsrf(r *http.Request) {
	td.CSRF = nosurf.Token(r)
}
