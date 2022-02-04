package handlers

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, 200},
	{"about", "/about", "GET", []postData{}, 200},
	{"gq", "/generals-quarters", "GET", []postData{}, 200},
	{"ms", "/majors-suite", "GET", []postData{}, 200},
	{"sa", "/search-availability", "GET", []postData{}, 200},
	{"contact", "/contact", "GET", []postData{}, 200},
	{"mr", "/make-reservation", "GET", []postData{}, 200},

	{"post-search-avial", "/search-availability", "POST", []postData{
		{key: "start", value: "2022-07-01"},
		{key: "end", value: "2022-07-02"},
	}, 200},
	{"post-search-avial-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2022-07-01"},
		{key: "end", value: "2022-07-02"},
	}, 200},
	{"post-make-reserve", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "youssef"},
		{key: "Last_name", value: "aly"},
		{key: "email", value: "ss@ss.com"},
		{key: "phone", value: "424737634"},
	}, 200},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()
	for _, test := range theTests {
		if test.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s expected code is %d but got %d", test.name,
					test.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range test.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+test.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s expected code is %d but got %d", test.name,
					test.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
