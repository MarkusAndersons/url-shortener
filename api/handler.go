// Package api contains all API logic to handle requests
package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// TODO replace this with a database
var dataStore = make(map[string]string)

// Store the requested URL in the data store and returns the shortened URL
func Store(w http.ResponseWriter, r *http.Request) {
	var request Request
	if err := decode(r, &request); err != nil {
		logResponse(r.URL, 400)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	shortURL := storeValue(request)
	if len(shortURL) == 0 {
		logResponse(r.URL, 500)
		w.WriteHeader(500)
		return
	}
	response := Response{ShortURL: shortURL}
	logResponse(r.URL, 200)
	json.NewEncoder(w).Encode(response)
}

// Get a previously stored short URL
func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["shortUrl"]
	longURL := getValue(shortURL)
	logRedirect(shortURL, longURL)
	http.Redirect(w, r, longURL, 301)
}

func storeValue(request Request) string {
	// TODO
	return ""
}

func createShortURL(url string) string {
	// TODO
	return ""
}

func getValue(key string) string {
	return dataStore[key]
}

func decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	if validatable, ok := v.(Ok); ok {
		return validatable.OK()
	}
	return nil
}

func logResponse(url *url.URL, status int) {
	log.Printf("%v %v", status, url)
}

func logRedirect(shortURL string, mappedURL string) {
	log.Printf("%v mapped to %v", shortURL, mappedURL)
}
