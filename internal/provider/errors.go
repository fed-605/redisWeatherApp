package provider

import "fmt"

var (
	errEmptyUrl      error = fmt.Errorf("url is not specified")
	errEmptyLocation error = fmt.Errorf("given an empty location")
	errInvalidApiKey error = fmt.Errorf("api key is not valid")
	errInvalidUrl    error = fmt.Errorf("invalid url for weather retrieving")
	errFetchWeather  error = fmt.Errorf("failed to fetch weather")
)
