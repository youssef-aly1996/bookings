package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/youssef-aly1996/bookings/internal/handlers"
)

func Test(t *testing.T) {
	var repo *handlers.Repository
	h := routes(repo)
	switch v := h.(type) {
	case *chi.Mux:
		//do nothin
	default:
		t.Error(fmt.Sprintf("this type is not *chi.mux, %T", v))
	}
}
