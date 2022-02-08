package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	TempalteCache map[string]*template.Template
	Logger        *log.Logger
	ErrorLog      *log.Logger
	PortNumber    string
	Session       *scs.SessionManager
	InProduction  bool
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}
