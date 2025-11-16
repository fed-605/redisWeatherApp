package app

import "github.com/fed-605/weatherApi/internal/cache"

type Service struct {
	cache cache.Cache
}

func NewService(cache cache.Cache) *Service {
	return &Service{
		cache: cache,
	}
}

func (s *Service) GetWeatherByCity() (*cache.WeatherResponse, error) {
	return nil, nil
}
