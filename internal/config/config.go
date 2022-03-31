package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/youssef-aly1996/bookings/internal/models"
)

type AppConfig struct {
	UseCache      bool
	TempalteCache map[string]*template.Template
	Logger        *log.Logger
	ErrorLog      *log.Logger
	PortNumber    string
	Session       *scs.SessionManager
	InProduction  bool
	MailChan chan models.MailModel
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}
