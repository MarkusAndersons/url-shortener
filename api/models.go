package api

import "database/sql"

// Ok interface is to be implemented by any object which can be validated
type Ok interface {
	OK() error
}

// DatabaseErr is the error returned if an error occurs with database access
type DatabaseErr struct {
	Msg string `json:"error"`
}

func (e DatabaseErr) Error() string {
	return e.Msg
}

// ErrRequired is the error returned if a required field in a request is missing
type ErrRequired struct {
	Msg string `json:"error"`
}

func (e ErrRequired) Error() string {
	return e.Msg
}

// Request is the in memory representation of a JSON request
type Request struct {
	URL string `json:"url"`
}

// OK validates a received request
func (r *Request) OK() error {
	if len(r.URL) == 0 {
		return ErrRequired{Msg: "url must be specified"}
	}
	return nil
}

// Response is the in memory representation of a JSON response
type Response struct {
	ShortURL string `json:"shortened"`
}

// DbResult is the value returned from a database
type DbResult struct {
	Value string
	Error error
}

// Db is an interface wrapper for a sql.DB to allow for testing
type Db interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (Rows, error)
}

// Rows is an interface wrapper for a sql.Rows to allow for testing
type Rows interface {
	Close() error
	Next() bool
	Scan(dest ...interface{}) error
}

// Database is a wrapper for a sql.DB object
type Database struct {
	Db *sql.DB
}

// RowsImpl is the implementation of Rows
type RowsImpl struct {
	R *sql.Rows
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.Db.Exec(query, args)
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (d *Database) Query(query string, args ...interface{}) (Rows, error) {
	r, err := d.Db.Query(query, args)
	return RowsImpl{R: r}, err
}

// Close closes the Rows, preventing further enumeration. If Next is called
// and returns false and there are no further result sets,
// the Rows are closed automatically and it will suffice to check the
// result of Err. Close is idempotent and does not affect the result of Err.
func (r RowsImpl) Close() error {
	return r.R.Close()
}

// Next prepares the next result row for reading with the Scan method. It
// returns true on success, or false if there is no next result row or an error
// happened while preparing it. Err should be consulted to distinguish between
// the two cases.
//
// Every call to Scan, even the first one, must be preceded by a call to Next.
func (r RowsImpl) Next() bool {
	return r.R.Next()
}

// Scan copies the columns in the current row into the values pointed
// at by dest. The number of values in dest must be the same as the
// number of columns in Rows.
func (r RowsImpl) Scan(dest ...interface{}) error {
	return r.R.Scan(dest)
}
