package cache

import (
	"context"
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
	Get(ctx context.Context,key string) (*WeatherResponse,error)
	Set(ctx context.Context,key string,value *WeatherResponse,ttl time.Duration) error
}