package render

import (
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"

	tc, err := CreateTempalteCache()
	if err != nil {
		t.Error(err)
	}
	testApp.TempalteCache = tc

	mw := myWriter{}
	err = RenderTemplate(mw, "home.page.tmpl", td)
	if err != nil {
		t.Error("failed writing template to the browser")
	}
	err = RenderTemplate(mw, "sayed.page.tmpl", td)
	if err == nil {
		t.Error("we rendered a template that not exist")
	}
}

func TestNewTemplate(t *testing.T) {
	NewTemplate(testApp)
}

func TestCreateTemplateCache(t *testing.T) {
	tc, err := CreateTempalteCache()
	if err != nil {
		t.Error("failed to run create template cache")
	}
	testApp.TempalteCache = tc
}

// func getRequest() (*http.Request, error) {
// 	r, err := http.NewRequest("GET", "some url", nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return r, nil
// }
