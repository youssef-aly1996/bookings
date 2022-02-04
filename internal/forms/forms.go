package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

//From creates a custom form struct, embeds a url value object
type From struct {
	url.Values
	Errors errors
}

//intializes a new pointer to form struct
func NewForm(data url.Values) *From {
	return &From{
		data,
		errors(map[string][]string{}),
	}
}

//Required checks for required fields
func (f *From) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "this field cannot be empty")
		}
	}
}

//Has checks a form field is in post and not empty
func (f *From) Has(field string) bool {
	formField := f.Get(field)
	return formField != ""
}

//Valid checks whether form inputs are valid or not
func (f *From) Valid() bool {
	return len(f.Errors) == 0
}

//MinLength checks for string minimum lenght
func (f *From) MinLength(field string, length int) bool {
	formField := f.Get(field)
	if len(formField) < length {
		f.Errors.Add(field, fmt.Sprintf("this field must be %d characters long", length))
		return false
	}
	return true
}

//IsEmail chkecks emails to be valid
func (f *From) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "this is invaild email address")
	}
}
