package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/youssef-aly1996/bookings/pkg/config"
	"github.com/youssef-aly1996/bookings/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(rw http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error
	if app.UseCache {
		tc = app.TempalteCache
	} else {
		tc, err = CreateTempalteCache()
		if err != nil {
			log.Fatal(err)
		}
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("coudn't get the template from the template cache")
	}
	buffer := new(bytes.Buffer)
	_ = t.Execute(buffer, td)
	_, err = buffer.WriteTo(rw)
	if err != nil {
		log.Fatal("error writing template to the browser", err.Error())
	}
}

func CreateTempalteCache() (map[string]*template.Template, error) {
	tempalteCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tempalteSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return nil, err
		}
		if len(matches) > 0 {
			tempalteSet, err = tempalteSet.ParseGlob("./templates/*.layout.html")
		}
		if err != nil {
			return nil, err
		}
		tempalteCache[name] = tempalteSet
	}
	return tempalteCache, nil
}
