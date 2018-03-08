package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/MarkusAndersons/url-shortener/api"
	"github.com/MarkusAndersons/url-shortener/constants"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = constants.Port
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	database := api.Database{Db: db}
	if err := api.DbInit(&database); err != nil {
		log.Fatalf("Error creating database")
	}

	router := mux.NewRouter()
	router.HandleFunc("/store", api.Store).Methods("POST")
	router.HandleFunc("/{shortUrl}", api.Get).Methods("GET")
	log.Println("Starting server at :" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
