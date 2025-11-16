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

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.WriteHeader(code)
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		errResponse(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	if _, err := w.Write(b); err != nil {
		errResponse(w, err.Error(), http.StatusInternalServerError)
	}
	return nil
}

// could be an error
