package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
Contract

	pattern:/weather/{location}
	method:GET
	info:pattern

suceed:

	-status code:200 OK
	- response: json with city weather

failed:

	-status code:...400...
	-response body:json with error + time
*/
func (app *Application) GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	location := chi.URLParam(r, "location")
	if location == "" {
		errResponse(w, "You give a non-existent location", http.StatusBadRequest)
		return
	}

	weather, err := app.Service.GetWeatherByCity(location)
	if err != nil {
		errResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, http.StatusOK, weather); err != nil {
		log.Println(err)
		errResponse(w, "critical server error", http.StatusInternalServerError)
		return
	}

}

func (app *Application) HealthzChecker(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
