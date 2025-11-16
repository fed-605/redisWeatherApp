package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type errDto struct {
	Message string    `json:"message"`
	ErrTime time.Time `json:errTime"`
}

func newErrDto(message string) errDto {
	return errDto{
		Message: message,
		ErrTime: time.Now(),
	}
}

func errResponse(w http.ResponseWriter, message string, code int) error {
	errDto := newErrDto(message)
	b, _ := json.MarshalIndent(errDto, "", "	")

	http.Error(w, string(b), code)
	return nil
}

// could be an error
