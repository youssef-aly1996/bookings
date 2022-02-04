package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type mockHandler struct {
}

func (mh mockHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
