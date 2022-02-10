package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/youssef-aly1996/bookings/internal/config"
	"github.com/youssef-aly1996/bookings/internal/models"
)

var functions = template.FuncMap{}
var pathToTemplates = "./templates"
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func Template(rw http.ResponseWriter, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template
	var err error
	if app.UseCache {
		tc = app.TempalteCache
	} else {
		tc, err = CreateTempalteCache()
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	t, ok := tc[tmpl]
	if !ok {
		return errors.New("this template is does not exsit")
	}
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, td)
	if err != nil {
		return err
	}
	_, err = buffer.WriteTo(rw)
	if err != nil {
		log.Println("error writing template to the browser", err.Error())
		return err
	}
	return nil
}

func CreateTempalteCache() (map[string]*template.Template, error) {
	tempalteCache := map[string]*template.Template{}
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tempalteSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return nil, err
		}
		if len(matches) > 0 {
			tempalteSet, err = tempalteSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		}
		if err != nil {
			return nil, err
		}
		tempalteCache[name] = tempalteSet
	}
	return tempalteCache, nil
}
