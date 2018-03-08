// Package api contains all API logic to handle requests
package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Driver for postgres
	_ "github.com/lib/pq"

	"github.com/MarkusAndersons/url-shortener/constants"
)

var db *sql.DB

func connect() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
}

func exec(query string) error {
	if _, err := db.Exec(query); err != nil {
		return DatabaseErr{Msg: err.Error()}
	}
	return nil
}

// DbInit sets up the database connection
func DbInit() error {
	connect()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS links (short VARCHAR(%d) NOT NULL, long text, PRIMARY KEY(short))", constants.ShortLen)
	return exec(query)
}

// DbStore stores a key-value pair
func DbStore(key string, value string) error {
	query := fmt.Sprintf("INSERT INTO links (short, long) VALUES (\"%s\", \"%s\")", key, value)
	return exec(query)
}

// DbGet returns the stored value for the given key
func DbGet(key string) DbResult {
	query := fmt.Sprintf("SELECT long FROM links WHERE short=\"%s\"", key)
	rows, err := db.Query(query)
	if err != nil {
		return DbResult{Value: "", Error: DatabaseErr{Msg: err.Error()}}
	}
	defer rows.Close()
	var longLink string
	if err := rows.Scan(&longLink); err != nil {
		return DbResult{Value: "", Error: DatabaseErr{Msg: err.Error()}}
	}
	return DbResult{Value: longLink, Error: nil}
}
