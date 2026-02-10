package app

import (
	"time"

	"github.com/fed-605/weatherApi/internal/cache"
	"github.com/fed-605/weatherApi/internal/provider"
)

const (
	version = "v1"
)

type Service struct {
	cache    cache.Cache
	provider provider.WeatherProvider
	ttl      time.Duration
}

func NewService(cache cache.Cache, provider provider.WeatherProvider, ttl time.Duration) *Service {
	return &Service{
		cache:    cache,
		provider: provider,
		ttl:      ttl,
	}
}

func (s *Service) GetWeatherByCity(location string) (*cache.WeatherResponse, error) {

	key := cache.BuildKey(version, location)

	cached, err := s.cache.Get(key)
	if err != nil {
		return nil, err
	}
	if cached != nil {
		return cached, nil
	}

	weather, err := s.provider.GetCurrent(location)
	if err != nil {
		return nil, err
	}

	if err = s.cache.Set(key, weather, s.ttl); err != nil {
		return nil, err
	}
	return weather, nil

}
