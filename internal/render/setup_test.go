package render

import (
	"net/http"
	"os"
	"testing"

	"github.com/youssef-aly1996/bookings/internal/config"
	"github.com/youssef-aly1996/bookings/internal/models"
)

var td = models.NewTemplateData()
var testApp = config.NewAppConfig()

func TestMain(m *testing.M) {
	NewTemplate(testApp)
	testApp.UseCache = false
	os.Exit(m.Run())
}

type myWriter struct{}

func (mw myWriter) Header() http.Header {
	return http.Header{}
}

func (mw myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil

}
func (mw myWriter) WriteHeader(statusCode int) {

}
