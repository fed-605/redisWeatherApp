package provider

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fed-605/weatherApi/internal/cache"
)

type WeatherProvider interface {
	GetCurrent(location string) (*cache.WeatherResponse, error)
}

type visualCrossingProvider struct {
	baseURL string
	apiKey  string
}

func NewvisualCrossingProvider(baseUrl string, apiKey string) *visualCrossingProvider {
	return &visualCrossingProvider{
		baseURL: baseUrl,
		apiKey:  apiKey,
	}
}

func (p *visualCrossingProvider) GetCurrent(location string) (*cache.WeatherResponse, error) {
	if err := validate(p.baseURL, p.apiKey, location); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/%s?unitGroup=metric&contentType=json&include=current&key=%s", p.baseURL, location, p.apiKey)
	log.Printf("request external url %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errInvalidUrl
	}
	if resp.StatusCode != 200 {
		resp.Body.Close()
		return nil, errFetchWeather
	}
	defer resp.Body.Close()
	var weatherData cache.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}
	return &weatherData, nil
}
