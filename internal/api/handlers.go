package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fed-605/weatherApi/env"
	"github.com/fed-605/weatherApi/internal/cache"
	"github.com/go-chi/chi/v5"
)

/*
Contract

	pattern:/api/v1/weather?city=...
	method:GET
	info:query params

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
	// add getenv
	baseUrl := env.GetEnvString("PROVIDER_BASE_URL", "")
	if baseUrl == "" {
		errResponse(w, "given an empty location", http.StatusBadRequest)
		return
	}
	key := env.GetEnvString("API_KEY", "")
	if key == "" {
		errResponse(w, "api key is not valid", http.StatusInternalServerError)
		return
	}
	url := fmt.Sprintf("%s/%s?unitGroup=metric&contentType=json&include=current&key=%s", baseUrl, location, key)
	log.Printf("request external url %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		errResponse(w, "invalid url for weather retrieving", http.StatusInternalServerError)
		return
	}
	if resp.StatusCode != 200 {
		errResponse(w, "failed to fetch weather", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// // for debugging (very dirty)
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	errResponse(w, "failed to read response body", http.StatusInternalServerError)
	// 	return
	// }
	// log.Printf("external raw body: %s", string(body))
	// // end of debugging

	var weatherData cache.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		errResponse(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	b, err := json.MarshalIndent(weatherData, "", "    ")
	if err != nil {
		errResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		errResponse(w, err.Error(), http.StatusInternalServerError)
	}

}

func (app *Application) HealthzChecker(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
