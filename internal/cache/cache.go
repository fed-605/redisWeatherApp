package cache

import (
	"fmt"
	"time"
)

type WeatherResponse struct {
	ResolvedAddress   string            `json:"resolvedAddress"`
	Timezone          string            `json:"timezone"`
	CurrentConditions CurrentConditions `json:"currentConditions"`
}

type CurrentConditions struct {
	Datetime      string  `json:"datetime"`
	Temperature   float64 `json:"temp"`
	Precipitation float64 `json:"precip"`
}

type Cache interface {
	Get(key string) (*WeatherResponse, error)
	Set(key string, value *WeatherResponse, ttl time.Duration) error
}

func BuildKey(version string, locaton string) string {
	return fmt.Sprintf("weather:%s:%s", version, locaton)
}
