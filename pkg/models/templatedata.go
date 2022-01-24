package models

type TemplateData struct {
	CSRF string
}

func NewTemplateData() *TemplateData {
	return &TemplateData{}
}
