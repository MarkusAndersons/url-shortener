package api_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MarkusAndersons/url-shortener/api"
)

// Model Tests
func TestRequestOk(t *testing.T) {
	database := testDatabase{}
	api.DbInit(&database)
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
	database := testDatabase{}
	api.DbInit(&database)
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
	database := testDatabase{}
	api.DbInit(&database)
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

type testDatabase struct {
}

type testRows struct {
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d *testDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (d *testDatabase) Query(query string, args ...interface{}) (api.Rows, error) {
	return testRows{}, nil
}

// Close closes the Rows, preventing further enumeration. If Next is called
// and returns false and there are no further result sets,
// the Rows are closed automatically and it will suffice to check the
// result of Err. Close is idempotent and does not affect the result of Err.
func (r testRows) Close() error {
	return nil
}

// Next prepares the next result row for reading with the Scan method. It
// returns true on success, or false if there is no next result row or an error
// happened while preparing it. Err should be consulted to distinguish between
// the two cases.
//
// Every call to Scan, even the first one, must be preceded by a call to Next.
func (r testRows) Next() bool {
	return true
}

// Scan copies the columns in the current row into the values pointed
// at by dest. The number of values in dest must be the same as the
// number of columns in Rows.
func (r testRows) Scan(dest ...interface{}) error {
	switch d := dest[0].(type) {
	case *string:
		*d = "http://google.com"
	}
	return nil
}
