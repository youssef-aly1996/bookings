package models

import "github.com/youssef-aly1996/bookings/internal/forms"

type TemplateData struct {
	CSRF      string
	Form      *forms.From
	Data      map[string]interface{}
	StringMap map[string]string
	Flash     string
	Error     string
	Warning   string
	IsAuthenticated int
}

func NewTemplateData() TemplateData {
	return TemplateData{}
}
