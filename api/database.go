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

// DbInit sets up the database connection
func DbInit() error {
	connect()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS links (short VARCHAR(%d) NOT NULL, long text, PRIMARY KEY(short))", constants.ShortLen)
	if _, err := db.Exec(query); err != nil {
		return DatabaseErr{Msg: err.Error()}
	}
	return nil
}
