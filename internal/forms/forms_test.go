package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	postedData := url.Values{}
	form := NewForm(postedData)
	if !form.Valid() {
		t.Error("got invalid form values")
	}
	// form.Errors.Add("form", "this is not a valid form")
}

func TestForm_Required(t *testing.T) {
	fields := []string{"aa", "bb", "cc"}
	form := NewForm(url.Values{})
	form.Required(fields...)
	if form.Valid() {
		t.Error("form is valid but post data are missing")
	}
	form.Errors.Add("sayed", "error")
	isError := form.Errors.Get("sayed")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}
	postedData := url.Values{}
	postedData.Add("aa", "aa")
	postedData.Add("bb", "bb")
	postedData.Add("cc", "cc")
	form = NewForm(postedData)
	form.Required(fields...)
	if !form.Valid() {
		t.Error("form required fields are missing")
	}
	isError = form.Errors.Get("aa")
	if isError != "" {
		t.Error("should not have an error but got one")
	}
}

func TestForm_Has(t *testing.T) {
	form := NewForm(url.Values{})
	has := form.Has("shit")
	if has {
		t.Error("form has a field but actually it doesn't")
	}
	postedData := url.Values{}
	postedData.Add("name", "name")
	form = NewForm(postedData)
	form.Has("name")
	if !form.Valid() {
		t.Error("the form does not have the field value")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("name", "youssef")
	form := NewForm(postedData)
	form.MinLength("name", 5)
	if !form.Valid() {
		t.Error("field length is not correct")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "you611@gmail.com")
	form := NewForm(postedData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got invalid form email")
	}
}
