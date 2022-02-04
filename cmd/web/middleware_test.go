package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNosurf(t *testing.T) {
	var mh mockHandler
	h := noSurf(mh)
	switch v := h.(type) {
	case http.Handler:
		//do nothin
	default:
		t.Error(fmt.Sprintf("this type is not http handler, %T", v))
	}
}
func TestScrffLoad(t *testing.T) {
	var mh mockHandler
	h := scrfLoad(mh)
	switch v := h.(type) {
	case http.Handler:
		//do nothin
	default:
		t.Error(fmt.Sprintf("this type is not http handler, %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var mh mockHandler
	h := sessionLoad(mh)
	switch v := h.(type) {
	case http.Handler:
		//do nothin
	default:
		t.Error(fmt.Sprintf("this type is not http handler, %T", v))
	}
}
