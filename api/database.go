// Package api contains all API logic to handle requests
package api

import (
	"fmt"

	// Driver for postgres
	_ "github.com/lib/pq"

	"github.com/MarkusAndersons/url-shortener/constants"
)

var db Db

func exec(query string) error {
	if _, err := db.Exec(query); err != nil {
		return DatabaseErr{Msg: err.Error()}
	}
	return nil
}

// DbInit sets up the database connection
func DbInit(database Db) error {
	db = database
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS links (short VARCHAR(%d) NOT NULL, long text, PRIMARY KEY(short))", constants.ShortLen)
	return exec(query)
}

// DbStore stores a key-value pair
func DbStore(key string, value string) error {
	query := fmt.Sprintf("INSERT INTO links (short, long) VALUES ('%s', '%s')", key, value)
	return exec(query)
}

// DbGet returns the stored value for the given key
func DbGet(key string) DbResult {
	query := fmt.Sprintf("SELECT long FROM links WHERE short='%s'", key)
	rows, err := db.Query(query)
	if err != nil {
		return DbResult{Value: "", Error: DatabaseErr{Msg: err.Error()}}
	}
	defer rows.Close()
	var longLink string
	rows.Next()
	if err := rows.Scan(&longLink); err != nil {
		return DbResult{Value: "", Error: DatabaseErr{Msg: err.Error()}}
	}
	return DbResult{Value: longLink, Error: nil}
}
