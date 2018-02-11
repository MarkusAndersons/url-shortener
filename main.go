package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/markusandersons/url-shortener/api"
	"github.com/markusandersons/url-shortener/constants"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = constants.Port
	}

	router := mux.NewRouter()
	router.HandleFunc("/store", api.Store).Methods("POST")
	router.HandleFunc("/{shortUrl}", api.Get).Methods("GET")
	log.Println("Starting server at :" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
