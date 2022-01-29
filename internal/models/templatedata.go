package models

import "github.com/youssef-aly1996/bookings/internal/forms"

type TemplateData struct {
	CSRF    string
	Form    *forms.From
	Data    map[string]interface{}
	Flash   string
	Error   string
	Warning string
}

func NewTemplateData() *TemplateData {
	return &TemplateData{}
}
