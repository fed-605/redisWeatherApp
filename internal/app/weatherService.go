package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fed-605/weatherApi/env"
	"github.com/fed-605/weatherApi/internal/cache"
)

type Service struct {
	cache cache.Cache
}

func NewService(cache cache.Cache) *Service {
	return &Service{
		cache: cache,
	}
}

func (s *Service) GetWeatherByCity(location string) (*cache.WeatherResponse, error) {
	baseUrl := env.GetEnvString("PROVIDER_BASE_URL", "")
	if baseUrl == "" {
		return nil, fmt.Errorf("given an empty location")
	}
	key := env.GetEnvString("API_KEY", "")
	if key == "" {
		return nil, fmt.Errorf("api key is not valid")
	}
	url := fmt.Sprintf("%s/%s?unitGroup=metric&contentType=json&include=current&key=%s", baseUrl, location, key)
	log.Printf("request external url %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("invalid url for weather retrieving")
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch weather")
	}
	defer resp.Body.Close()

	var weatherData cache.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}
	return &weatherData, nil
}
