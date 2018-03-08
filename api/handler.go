// Package api contains all API logic to handle requests
package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/MarkusAndersons/url-shortener/constants"

	"github.com/gorilla/mux"
)

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
	if shortURL.Error != nil {
		logResponse(r.URL, 500)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(shortURL.Error)
		return
	}
	if len(shortURL.Value) == 0 {
		logResponse(r.URL, 500)
		w.WriteHeader(500)
		return
	}
	response := Response{ShortURL: shortURL.Value}
	logResponse(r.URL, 200)
	json.NewEncoder(w).Encode(response)
}

// Get a previously stored short URL
func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["shortUrl"]
	longURL := getValue(shortURL)
	if longURL.Error != nil {
		logResponse(r.URL, 500)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(longURL.Error)
		return
	}
	logRedirect(shortURL, longURL.Value)
	http.Redirect(w, r, longURL.Value, 301)
}

func storeValue(request Request) DbResult {
	short := createShortURL(request.URL)
	if err := DbStore(short, request.URL); err != nil {
		return DbResult{Value: "", Error: err}
	}
	return DbResult{Value: short, Error: nil}
}

func createShortURL(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	short := hex.EncodeToString(hasher.Sum(nil))
	return short[len(short)-constants.ShortLen:]
}

func getValue(key string) DbResult {
	return DbGet(key)
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
