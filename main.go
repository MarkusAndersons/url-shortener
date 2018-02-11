package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/markusandersons/url-shortener/api"
)

const (
	// PORT is the default port number to use if none specified
	PORT string = "8080"
	// HOSTNAME is the host to run the server on
	HOSTNAME string = "127.0.0.1"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = PORT
	}
	uri := HOSTNAME + ":" + port

	router := mux.NewRouter()
	router.HandleFunc("/store", api.Store).Methods("POST")
	router.HandleFunc("/{shortUrl}", api.Get).Methods("GET")
	log.Println("Starting server at", uri)
	log.Fatal(http.ListenAndServe(uri, router))
}
