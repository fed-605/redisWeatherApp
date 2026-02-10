package api

import "time"

type errDto struct {
	Message string    `json:"message"`
	ErrTime time.Time `json:"errTime"`
}

func newErrDto(message string) errDto {
	return errDto{
		Message: message,
		ErrTime: time.Now(),
	}
}
