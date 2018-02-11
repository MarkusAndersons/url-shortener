// Package api contains all API logic to handle requests
package api

import (
	"encoding/json"
	"net/http"
)

// TODO replace this with a database
var dataStore = make(map[string]string)

func decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	if validatable, ok := v.(Ok); ok {
		return validatable.OK()
	}
	return nil
}
