package erroring

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/youssef-aly1996/bookings/internal/config"
)

type Erroring struct {
	appConfig *config.AppConfig
}

func NewErroring(a *config.AppConfig) Erroring {
	return Erroring{appConfig: a}
}

func (e Erroring) ClientErrors(rw http.ResponseWriter, codeStatus int) {
	e.appConfig.Logger.Println("Client Error with status", codeStatus)
	http.Error(rw, http.StatusText(codeStatus), codeStatus)
}

func (e Erroring) ServerErrors(rw http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	e.appConfig.ErrorLog.Println(trace)
	http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
}
