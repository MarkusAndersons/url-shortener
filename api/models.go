package api

// Ok interface is to be implemented by any object which can be validated
type Ok interface {
	OK() error
}

// ErrRequired is the error returned if a required field in a request is missing
type ErrRequired struct {
	Msg string
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
		return ErrRequired{Msg: "url"}
	}
	return nil
}

// Response is the in memory representation of a JSON response
type Response struct {
	ShortURL string `json:"shortened"`
}
