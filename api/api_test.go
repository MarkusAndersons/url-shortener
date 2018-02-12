package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/markusandersons/url-shortener/api"
)

// Model Tests
func TestRequestOk(t *testing.T) {
	request := api.Request{URL: ""}
	err := request.OK()
	if err == nil {
		t.Fail()
	}
	if _, ok := err.(api.ErrRequired); !ok {
		t.Fail()
	}
}

// Controller Tests
func TestGetHandlerRedirects(t *testing.T) {
	r, err := http.NewRequest("GET", "/abcd", nil)
	if err != nil {
		t.Fail()
	}
	w := httptest.NewRecorder()

	api.Get(w, r)

	if w.Code != http.StatusMovedPermanently {
		t.Fail()
	}
}

func TestStoreHandler(t *testing.T) {
	body := strings.NewReader(
		"{\"url\":\"http://google.com\"}",
	)
	r, err := http.NewRequest("POST", "/store", body)
	if err != nil {
		t.Fail()
	}
	w := httptest.NewRecorder()

	api.Store(w, r)

	if w.Code != http.StatusOK {
		t.Fail()
	}
	s := w.Body.String()
	if s != "{\"shortened\":\"4e4da5\"}\n" {
		t.Fail()
	}
}
